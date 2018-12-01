using System;
using System.Threading;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ISessionServer : IDisposable
    {
        AutoResetEvent WaitServerStartEvent { get; set; }

        void Start();
        void Stop();
    }
}
