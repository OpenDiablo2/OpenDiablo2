using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public interface IMapInfo
    {
        Dictionary<eRenderCellType, MapCellInfo[]> CellInfo { get; set; }
        MPQDS1 FileData { get; set; }
        Rectangle TileLocation { get; set; }
    }

    public sealed class MapInfo : IMapInfo
    {
        public int LevelId { get; set; } = (int)eLevelId.None;
        public MPQDS1 FileData { get; set; }
        public Dictionary<eRenderCellType, MapCellInfo[]> CellInfo { get; set; }
        public Rectangle TileLocation { get; set; } = new Rectangle();
    }

    public sealed class SubMapInfo : IMapInfo
    {
        public IMapInfo PrimaryMap { get; set; }
        public Dictionary<eRenderCellType, MapCellInfo[]> CellInfo { get; set; }
        public MPQDS1 FileData { get; set; }
        public Rectangle TileLocation { get; set; } = new Rectangle();
    }
}
