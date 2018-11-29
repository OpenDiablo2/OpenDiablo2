using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public sealed class MapInfo
    {
        public eLevelId LevelId { get; set; } = eLevelId.None;
        public MapInfo PrimaryMap { get; set; } = null;
        public MPQDS1 FileData { get; set; }
        public LevelPreset LevelPreset { get; set; }
        public LevelDetail LevelDetail { get; set; }
        public LevelType LevelType { get; set; }
        public Dictionary<eRenderCellType, MapCellInfo[]> CellInfo { get; set; }
        public Rectangle TileLocation { get; set; } = new Rectangle();
    }
}
