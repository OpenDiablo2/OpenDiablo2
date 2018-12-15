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
        eDifficulty Difficulty { get; }

        ItemInstance SelectedItem { get; }
        void SelectItem(ItemInstance item);

        int CameraOffset { get; set; }

        void Initialize(string characterName, eHero hero, eSessionType sessionType, eDifficulty difficulty);
        void Update(long ms);
        IEnumerable<MapCellInfo> GetMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType);
        void UpdateMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType, IEnumerable<MapCellInfo> mapCellInfo);
        IMapInfo InsertMap(eLevelId levelId, IMapInfo parentMap = null);
        IMapInfo InsertMap(int levelId, Point origin, IMapInfo parentMap = null);
        IMapInfo InsertSubMap(int levelPresetId, int levelTypeId, Point origin, IMapInfo primaryMap, int subTile = -1);
        IMapInfo GetSubMapInfo(int levelPresetId, int levelTypeId, IMapInfo primaryMap, Point origin, int subTile = -1);
        void AddMap(IMapInfo map);
        int HasMap(int cellX, int cellY);
        IEnumerable<Size> GetMapSizes(int cellX, int cellY);
        void RemoveEverythingAt(int cellX, int cellY);
    }
}
