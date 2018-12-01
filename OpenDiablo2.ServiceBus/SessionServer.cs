using System;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using NetMQ;
using NetMQ.Sockets;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.ServiceBus.Message_Frames;

namespace OpenDiablo2.ServiceBus
{
    public sealed class SessionServer : ISessionServer, ISessionEventProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly eSessionType sessionType;
        private readonly Func<eMessageFrameType, IMessageFrame> getMessageFrame;
        private AutoResetEvent resetEvent = new AutoResetEvent(false);
        public AutoResetEvent WaitServerStartEvent { get; set; } = new AutoResetEvent(false);

        private int gameSeed;
        private bool running = false;
        private ResponseSocket responseSocket;

        public OnSetSeedEvent OnSetSeed { get; set; }
        public OnJoinGameEvent OnJoinGame { get; set; }

        public SessionServer(eSessionType sessionType, Func<eMessageFrameType, IMessageFrame> getMessageFrame)
        {
            this.sessionType = sessionType;
            this.getMessageFrame = getMessageFrame;
        }

        public void Start()
        {
            gameSeed = (new Random()).Next();
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

            var proactor = new NetMQProactor(responseSocket, (socket, message) =>
            {
                var bytes = message.First().ToByteArray();
                var frameType = (eMessageFrameType)bytes[0];
                var frameData = bytes.Skip(1).ToArray(); // TODO: Can we maybe use pointers? This seems wasteful
                    var messageFrame = getMessageFrame(frameType);
                messageFrame.Data = frameData;
                messageFrame.Process(socket, this);
            });
            running = true;
            WaitServerStartEvent.Set();
            resetEvent.WaitOne();
            proactor.Dispose();
            running = false;
            responseSocket.Dispose();
            log.Info("Session server has stopped.");
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

        private void Send(NetMQSocket target, IMessageFrame messageFrame)
        {
            var attr = messageFrame.GetType().GetCustomAttributes(true).First(x => typeof(MessageFrameAttribute).IsAssignableFrom(x.GetType())) as MessageFrameAttribute;
            responseSocket.SendFrame(new byte[] { (byte)attr.FrameType }.Concat(messageFrame.Data).ToArray());
        }

        private void OnJoinGameHandler(object sender, Guid playerId, string playerName)
        {
            // TODO: Try to make this less stupid
            Send(sender as NetMQSocket, new MFSetSeed(gameSeed));
        }
    }
}
