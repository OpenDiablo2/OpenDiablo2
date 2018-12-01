using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnSetSeedEvent(object sender, int seed);
    public delegate void OnJoinGameEvent(object sender, Guid playerId, string playerName); // TODO: Not the final version..

    public interface ISessionEventProvider
    {

        OnSetSeedEvent OnSetSeed { get; set; }
        OnJoinGameEvent OnJoinGame { get; set; }
    }
}
