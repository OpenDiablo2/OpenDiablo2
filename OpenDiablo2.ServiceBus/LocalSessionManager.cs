using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus
{
    public sealed class LocalSessionManager : ISessionManager
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly ISessionServer sessionServer;
        volatile bool running = false;

        public LocalSessionManager(ISessionServer sessionServer)
        {
            this.sessionServer = sessionServer;
        }

        public void Initialize()
        {
            log.Info("Initializing a local multiplayer session.");
            running = true;
            Task.Run(() => Listen());
        }

        private void Listen()
        {
            log.Info("Local session manager is starting.");
            while (running)
            {

            }

            log.Info("Local session manager has stopped.");
        }

        public void Dispose()
        {
            
        }

        public void Stop() => running = false;
    }
}
