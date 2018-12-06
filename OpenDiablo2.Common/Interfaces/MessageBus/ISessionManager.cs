using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Interfaces
{

    public interface ISessionManager : ISessionEventProvider, IDisposable
    {
        void Initialize();
        void Stop();

        void JoinGame(string playerName, eHero heroType);
        void MoveRequest(byte direction, eMovementType movementType);
    }
}
