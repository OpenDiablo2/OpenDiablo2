﻿using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using NetMQ;
using NetMQ.Sockets;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.MessageBus;
using OpenDiablo2.ServiceBus.Message_Frames.Server;

namespace OpenDiablo2.ServiceBus
{
    public sealed class SessionServer : ISessionServer, ISessionEventProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly eSessionType sessionType;
        private readonly IGameServer gameServer;
        private readonly Func<eMessageFrameType, IMessageFrame> getMessageFrame;

        private AutoResetEvent resetEvent = new AutoResetEvent(false);
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
                    throw new ApplicationException("This session type is currently unsupported.");
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
            resetEvent.WaitOne();
            proactor.Dispose();
            running = false;
            responseSocket.Dispose();
            log.Info("Session server has stopped.");
        }

        private void OnMovementRequestHandler(int clientHash, int direction, eMovementType movementType)
        {
            // TODO: Actually move the player ....
            var player = gameServer.Players.FirstOrDefault(x => x.ClientHash == clientHash);
            if (player == null)
                return;
            log.Info($"Player {player.Id} requested to move in {direction} direction: {movementType.ToString()}");

            NoOp();
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
            var attr = messageFrame.GetType().GetCustomAttributes(true).First(x => typeof(MessageFrameAttribute).IsAssignableFrom(x.GetType())) as MessageFrameAttribute;
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
