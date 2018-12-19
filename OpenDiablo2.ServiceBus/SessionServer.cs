/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using NetMQ;
using NetMQ.Sockets;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using OpenDiablo2.ServiceBus.Message_Frames.Server;

namespace OpenDiablo2.ServiceBus
{
    public sealed class SessionServer : ISessionServer, ISessionEventProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly eSessionType sessionType;
        private readonly IGameServer gameServer;
        private readonly Func<eMessageFrameType, IMessageFrame> getMessageFrame;

        private readonly AutoResetEvent resetEvent = new AutoResetEvent(false);
        public AutoResetEvent WaitServerStartEvent { get; set; } = new AutoResetEvent(false);
        private bool running = false;
        private ResponseSocket responseSocket;

        public OnJoinGameEvent OnJoinGame { get; set; }
        public OnMoveRequest OnMoveRequest { get; set; }
        public OnUpdateEquipmentEvent OnUpdateEquipment { get; set; }

        // TODO: Fix interface so we don't need this in the session server
        public OnSetSeedEvent OnSetSeed { get; set; }
        public OnLocatePlayersEvent OnLocatePlayers { get; set; }
        public OnPlayerInfoEvent OnPlayerInfo { get; set; }
        public OnFocusOnPlayer OnFocusOnPlayer { get; set; }
        public OnChangeEquipment OnChangeEquipment { get; set; }

        const int serverUpdateRate = 30;

        public SessionServer(
            eSessionType sessionType,
            IGameServer gameServer,
            Func<eMessageFrameType, IMessageFrame> getMessageFrame
            )
        {
            this.sessionType = sessionType;
            this.getMessageFrame = getMessageFrame;
            this.gameServer = gameServer;
        }

        public void Start()
        {
            // TODO: Loading existing games...
            gameServer.InitializeNewGame();

            Task.Run(() => Serve());
        }

        private void Serve()
        {
            log.Info("Session server is starting.");
            responseSocket = new ResponseSocket();

            switch (sessionType)
            {
                case eSessionType.Local:
                    responseSocket.Bind("inproc://opendiablo2-session");
                    break;
                case eSessionType.Server:
                case eSessionType.Remote:
                default:
                    throw new OpenDiablo2Exception("This session type is currently unsupported.");
            }

            OnJoinGame += OnJoinGameHandler;
            OnMoveRequest += OnMovementRequestHandler;
            OnUpdateEquipment += OnUpdateEquipmentHandler;

            var proactor = new NetMQProactor(responseSocket, (socket, message) =>
            {
                foreach (var msg in message)
                {
                    using (var ms = new MemoryStream(msg.ToByteArray()))
                    using (var br = new BinaryReader(ms))
                    {
                        var messageFrame = getMessageFrame((eMessageFrameType)br.ReadByte());
                        messageFrame.LoadFrom(br);
                        messageFrame.Process(socket.GetHashCode(), this);
                    }
                }
            });
            running = true;
            WaitServerStartEvent.Set();
            Task.Run(() =>
            {
                var lastRun = DateTime.Now;
                while (running)
                {
                    var newTime = DateTime.Now;
                    var timeDiff = (newTime - lastRun).TotalMilliseconds;
                    lastRun = newTime;

                    gameServer.Update((int)timeDiff);
                    if (timeDiff < serverUpdateRate)
                        Thread.Sleep((int)Math.Min(serverUpdateRate, Math.Max(0, serverUpdateRate - timeDiff)));
                }
            });
            resetEvent.WaitOne();
            proactor.Dispose();
            running = false;
            responseSocket.Dispose();
            log.Info("Session server has stopped.");
        }


        private void OnMovementRequestHandler(int clientHash, PointF targetCell, eMovementType movementType)
        {
            var player = gameServer.Players.FirstOrDefault(x => x.ClientHash == clientHash);
            if (player == null)
                return;

            player.MovementType = movementType;
            player.MovementSpeed = (player.MovementType == eMovementType.Running ? player.GetRunVelocity() : player.GetWalkVeloicty()) / 4f;
            player.Waypoints = CalculateWaypoints(player, targetCell);

            Send(new MFLocatePlayers(gameServer.Players.Select(x => x.ToPlayerLocationDetails())));
        }

        private List<PointF> CalculateWaypoints(PlayerState player, PointF targetCell)
        {
            // TODO: Move this somewhere else...
            var result = new List<PointF>();
            result.Add(targetCell);
            /*
            // Ensure they aren't sending crazy coordinates..
            var targetX = Math.Round(targetCell.X, 1);
            var targetY = Math.Round(targetCell.Y, 1);

            // TODO: Legit Pathfind here...
            result.Add(new PointF(player.X, player.Y));
            int maxTries = 50;
            var curX = player.X;
            var curY = player.Y;
            var nextX = curX;
            var nextY = curY;
            while (--maxTries > 0)
            {
                if (curX < targetX)
                    nextX += .1f;
                else if (curX > targetX)
                    nextX -= .1f;

                if (curY < targetY)
                    nextY += .1f;
                else if (curY > targetY)
                    nextY -= .1f;

                result.Add(new PointF((float)Math.Round(nextX, 1), (float)Math.Round(nextY, 1)));

                curX = nextX;
                curY = nextY;

                // If we reached our target, stop here
                if (Math.Abs(curX - targetX) < 0.1f && Math.Abs(curY - targetY) < 0.1f)
                    break;
            }

            */
            return result;
        }

        public void Stop()
        {
            if (!running)
                return;

            resetEvent.Set();
        }

        public void Dispose()
        {
            Stop();
        }

        private void NoOp()
        {
            responseSocket.SendFrame(new byte[] { (byte)eMessageFrameType.None });
        }

        private void Send(IMessageFrame messageFrame, bool more = false)
        {
            var attr = messageFrame.GetType().GetCustomAttributes(true).First(x => (x is MessageFrameAttribute)) as MessageFrameAttribute;
            using (var ms = new MemoryStream())
            using (var br = new BinaryWriter(ms))
            {
                br.Write((byte)attr.FrameType);
                messageFrame.WriteTo(br);
                responseSocket.SendFrame(ms.ToArray(), more);
            }
        }

        private void OnJoinGameHandler(int clientHash, eHero heroType, string playerName)
        {
            gameServer.SpawnNewPlayer(clientHash, playerName, heroType);
            Send(new MFSetSeed(gameServer.Seed), true);
            Send(new MFPlayerInfo(gameServer.Players.Select(x => x.ToPlayerInfo())), true);
            Send(new MFLocatePlayers(gameServer.Players.Select(x => x.ToPlayerLocationDetails())), true);
            Send(new MFFocusOnPlayer(gameServer.Players.First(x => x.ClientHash == clientHash).UID));
        }

        private void OnUpdateEquipmentHandler(int clientHash, string Slot, ItemInstance itemInstance)
        {
            var player = gameServer.Players.FirstOrDefault(x => x.ClientHash == clientHash);
            if (player == null)
                return;

            var equipment = gameServer.UpdateEquipment(clientHash, Slot, itemInstance);
            Send(new MFChangeEquipment(player.UID, equipment));
        }
    }
}
