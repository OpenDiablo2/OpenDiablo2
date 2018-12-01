﻿using System;
using System.Collections.Generic;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IGameServer : IDisposable
    {
        IEnumerable<PlayerState> Players { get; }
        int Seed { get; }

        void InitializeNewGame();
        int SpawnNewPlayer(int clientHash, string playerName, eHero heroType);
    }
}
