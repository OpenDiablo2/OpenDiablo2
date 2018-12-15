using System;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnSetSeedEvent(int clientHash, int seed);
    public delegate void OnJoinGameEvent(int clientHash, eHero heroType, string playerName);
    public delegate void OnLocatePlayersEvent(int clientHash, IEnumerable<LocationDetails> playerLocationDetails);
    public delegate void OnPlayerInfoEvent(int clientHash, IEnumerable<PlayerInfo> playerInfo);
    public delegate void OnFocusOnPlayer(int clientHash, Guid playerId);
    public delegate void OnMoveRequest(int clientHash, PointF targetCell, eMovementType movementType);

    public interface ISessionEventProvider
    {
        OnSetSeedEvent OnSetSeed { get; set; }
        OnJoinGameEvent OnJoinGame { get; set; }
        OnLocatePlayersEvent OnLocatePlayers { get; set; }
        OnPlayerInfoEvent OnPlayerInfo { get; set; }
        OnFocusOnPlayer OnFocusOnPlayer { get; set; }
        OnMoveRequest OnMoveRequest { get; set; }
    }
}
