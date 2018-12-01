using System;
using System.Collections.Generic;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces.MessageBus;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnSetSeedEvent(int clientHash, int seed);
    public delegate void OnJoinGameEvent(int clientHash, eHero heroType, string playerName);
    public delegate void OnLocatePlayersEvent(int clientHash, IEnumerable<PlayerLocationDetails> playerLocationDetails);

    public interface ISessionEventProvider
    {
        OnSetSeedEvent OnSetSeed { get; set; }
        OnJoinGameEvent OnJoinGame { get; set; }
        OnLocatePlayersEvent OnLocatePlayers { get; set; }
    }
}
