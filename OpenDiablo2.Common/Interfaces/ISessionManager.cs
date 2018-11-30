using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ISessionManager : IDisposable
    {
        void Initialize();
        void Stop();
    }
}
