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

        // TODO: Fix interface so we don't need this in the session server
        public OnSetSeedEvent OnSetSeed { get; set; }
        public OnLocatePlayersEvent OnLocatePlayers { get; set; }
        public OnPlayerInfoEvent OnPlayerInfo { get; set; }
        public OnFocusOnPlayer OnFocusOnPlayer { get; set; }

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

            var proactor = new NetMQProactor(responseSocket, (socket, message) =>
            {
                var bytes = message.First().ToByteArray();
                var frameType = (eMessageFrameType)bytes[0];
                var frameData = bytes.Skip(1).ToArray(); // TODO: Can we maybe use pointers? This seems wasteful
                    var messageFrame = getMessageFrame(frameType);
                messageFrame.Data = frameData;
                messageFrame.Process(socket.GetHashCode(), this);
            });
            running = true;
            WaitServerStartEvent.Set();
            Task.Run(() =>
            {
                var lastRun = DateTime.Now;
                while(running)
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
        

        private void OnMovementRequestHandler(int clientHash, byte direction, eMovementType movementType)
        {
            var player = gameServer.Players.FirstOrDefault(x => x.ClientHash == clientHash);
            if (player == null)
                return;

            player.MovementDirection = direction;
            player.MovementType = movementType;
            player.MovementDirection = direction;


            Send(new MFLocatePlayers(gameServer.Players.Select(x => x.ToPlayerLocationDetails())));
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
            responseSocket.SendFrame(new byte[] { (byte)attr.FrameType }.Concat(messageFrame.Data).ToArray(), more);
        }

        private void OnJoinGameHandler(int clientHash, eHero heroType, string playerName)
        {
            gameServer.SpawnNewPlayer(clientHash, playerName, heroType);
            Send(new MFSetSeed(gameServer.Seed), true);
            Send(new MFPlayerInfo(gameServer.Players.Select(x => x.ToPlayerInfo())), true);
            Send(new MFLocatePlayers(gameServer.Players.Select(x => x.ToPlayerLocationDetails())), true);
            Send(new MFFocusOnPlayer(gameServer.Players.First(x => x.ClientHash == clientHash).Id));
        }
    }
}
