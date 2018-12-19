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
using OpenDiablo2.ServiceBus.Message_Frames.Client;
using OpenDiablo2.ServiceBus.Message_Frames.Server;

namespace OpenDiablo2.ServiceBus
{
    public sealed class SessionManager : ISessionManager
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly Func<eSessionType, ISessionServer> getSessionServer;
        private readonly eSessionType sessionType;
        private readonly Func<eMessageFrameType, IMessageFrame> getMessageFrame;
        private readonly Func<IGameState> getGameState;

        private RequestSocket requestSocket;
        private readonly AutoResetEvent resetEvent = new AutoResetEvent(false);
        private ISessionServer sessionServer;
        private bool running = false;

        public OnSetSeedEvent OnSetSeed { get; set; }
        public OnJoinGameEvent OnJoinGame { get; set; }
        public OnLocatePlayersEvent OnLocatePlayers { get; set; }
        public OnPlayerInfoEvent OnPlayerInfo { get; set; }
        public OnFocusOnPlayer OnFocusOnPlayer { get; set; }
        public OnMoveRequest OnMoveRequest { get; set; }
        public OnUpdateEquipmentEvent OnUpdateEquipment { get; set; }
        public OnChangeEquipment OnChangeEquipment { get; set; }

        public SessionManager(
            eSessionType sessionType,
            Func<eSessionType, ISessionServer> getSessionServer,
            Func<eMessageFrameType, IMessageFrame> getMessageFrame,
            Func<IGameState> getGameState
            )
        {
            this.getSessionServer = getSessionServer;
            this.sessionType = sessionType;
            this.getMessageFrame = getMessageFrame;
            this.getGameState = getGameState;
        }

        public void Initialize()
        {
            if (sessionType == eSessionType.Local || sessionType == eSessionType.Server)
            {
                sessionServer = getSessionServer(sessionType);
                sessionServer.Start();
                sessionServer.WaitServerStartEvent.WaitOne(); // Wait until the server starts...
            }
            else sessionServer = null;

            log.Info("Initializing a local multiplayer session.");
            Task.Run(() => Listen());
        }

        private void Listen()
        {
            log.Info("Session manager is starting.");
            requestSocket = new RequestSocket();

            switch (sessionType)
            {
                case eSessionType.Local:
                    requestSocket.Connect("inproc://opendiablo2-session");
                    break;
                case eSessionType.Server:
                case eSessionType.Remote:
                default:
                    throw new OpenDiablo2Exception("This session type is currently unsupported.");
            }

            running = true;
            resetEvent.WaitOne();
            running = false;
            requestSocket.Dispose();
            log.Info("Session manager has stopped.");

        }
        public void Stop()
        {
            if (!running)
                return;

            resetEvent.Set();

            if (sessionType == eSessionType.Local || sessionType == eSessionType.Server)
                sessionServer?.Stop();

        }

        public void Dispose()
        {
            Stop();
        }

        public void Send(IMessageFrame messageFrame, bool more = false)
        {
            var attr = messageFrame.GetType().GetCustomAttributes(true).First(x => (x is MessageFrameAttribute)) as MessageFrameAttribute;
            using (var ms = new MemoryStream())
            using (var bw = new BinaryWriter(ms))
            {
                bw.Write((byte)attr.FrameType);
                messageFrame.WriteTo(bw);
                requestSocket.SendFrame(ms.ToArray(), more);
            }
        }

        private void ProcessMessageFrame<T>() where T : IMessageFrame, new()
        {
            if (!running)
                throw new OpenDiablo2Exception("You have made a terrible mistake. Cannot get a message frame if you are not connected.");

            using (var ms = new MemoryStream(requestSocket.ReceiveFrameBytes()))
            using (var br = new BinaryReader(ms))
            {
                var messageFrame = getMessageFrame((eMessageFrameType)br.ReadByte());

                if (messageFrame.GetType() != typeof(T))
                    throw new OpenDiablo2Exception("Recieved unexpected message frame!");

                messageFrame.LoadFrom(br);

                lock (getGameState().ThreadLocker)
                    messageFrame.Process(requestSocket.GetHashCode(), this);
            }
        }

        private void NoOp()
        {
            var bytes = requestSocket.ReceiveFrameBytes();
            if ((eMessageFrameType)bytes[0] != eMessageFrameType.None)
                throw new OpenDiablo2Exception("Excepted a NoOp but got a command instead!");
        }

        public void JoinGame(string playerName, eHero heroType)
        {
            Task.Run(() =>
            {
                Send(new MFJoinGame(playerName, heroType));
                ProcessMessageFrame<MFSetSeed>();
                ProcessMessageFrame<MFPlayerInfo>();
                ProcessMessageFrame<MFLocatePlayers>();
                ProcessMessageFrame<MFFocusOnPlayer>();
            });
        }

        public void MoveRequest(PointF targetCell, eMovementType movementType)
            => Task.Run(() =>
            {
                Send(new MFMoveRequest(targetCell, movementType));
                ProcessMessageFrame<MFLocatePlayers>();
            });

        public void UpdateEquipment(string slot, ItemInstance itemInstance)
        {
            Task.Run(() =>
            {
                Send(new MFUpdateEquipment(slot, itemInstance));
                ProcessMessageFrame<MFChangeEquipment>();
            });
        }
    }
}
