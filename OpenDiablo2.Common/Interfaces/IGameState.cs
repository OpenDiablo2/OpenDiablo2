using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IGameState
    {
        int Act { get; }
        int Seed { get; }
        string MapName { get; }
        Palette CurrentPalette { get; }

        bool ToggleShowInventoryPanel();
        bool ShowInventoryPanel { get; set; }

        bool ToggleShowCharacterPanel();
        bool ShowCharacterPanel { get; set; }

        void Initialize(string text, eHero value);
        void Update(long ms);
        IEnumerable<MapCellInfo> GetMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType);
        void UpdateMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType, IEnumerable<MapCellInfo> mapCellInfo);
        MapInfo LoadMap(eLevelId levelId, Point origin);
        MapInfo LoadSubMap(int levelDefId, Point origin);
    }
}
