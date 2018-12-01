using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{

    public interface ISessionManager : ISessionEventProvider, IDisposable
    {
        void Initialize();
        void Stop();

        void JoinGame(string playerName, Action<Guid> callback);
    }
}
