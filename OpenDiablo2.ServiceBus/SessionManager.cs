using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using NetMQ;
using NetMQ.Sockets;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
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
        private AutoResetEvent resetEvent = new AutoResetEvent(false);
        private ISessionServer sessionServer;
        public Guid PlayerId { get; private set; }
        private bool running = false;

        public OnSetSeedEvent OnSetSeed { get; set; }
        public OnJoinGameEvent OnJoinGame { get; set; }
        public OnLocatePlayersEvent OnLocatePlayers { get; set; }
        public OnPlayerInfoEvent OnPlayerInfo { get; set; }
        public OnFocusOnPlayer OnFocusOnPlayer { get; set; }

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
                    throw new ApplicationException("This session type is currently unsupported.");
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

        public void Send(IMessageFrame messageFrame)
        {
            var attr = messageFrame.GetType().GetCustomAttributes(true).First(x => typeof(MessageFrameAttribute).IsAssignableFrom(x.GetType())) as MessageFrameAttribute;
            requestSocket.SendFrame(new byte[] { (byte)attr.FrameType }.Concat(messageFrame.Data).ToArray());
        }

        private void ProcessMessageFrame<T>() where T : IMessageFrame, new()
        {
            if (!running)
                throw new ApplicationException("You have made a terrible mistake. Cannot get a message frame if you are not connected.");

            var bytes = requestSocket.ReceiveFrameBytes();
            var frameType = (eMessageFrameType)bytes[0];
            var frameData = bytes.Skip(1).ToArray(); // TODO: Can we maybe use pointers? This seems wasteful
            var messageFrame = getMessageFrame(frameType);
            if (messageFrame.GetType() != typeof(T))
                throw new ApplicationException("Recieved unexpected message frame!");
            messageFrame.Data = frameData;
            lock (getGameState().ThreadLocker)
            {
                messageFrame.Process(requestSocket.GetHashCode(), this);
            }
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
    }
}
