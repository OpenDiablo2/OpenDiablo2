using System;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IGameState : IDisposable
    {
        object ThreadLocker { get; }

        int Act { get; }
        int Seed { get; }
        string MapName { get; }
        Palette CurrentPalette { get; }
        IEnumerable<PlayerInfo> PlayerInfos { get; }

        Item SelectedItem { get; }
        void SelectItem(Item item);

        int CameraOffset { get; set; }

        void Initialize(string text, eHero value, eSessionType sessionType);
        void Update(long ms);
        IEnumerable<MapCellInfo> GetMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType);
        void UpdateMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType, IEnumerable<MapCellInfo> mapCellInfo);
        MapInfo LoadMap(eLevelId levelId, Point origin);
        MapInfo LoadSubMap(int levelDefId, Point origin);
    }
}
