using System;
using System.Drawing;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Common.Models
{
    // Represents a single cell on a map
    public sealed class MapCellInfo
    {
        public Guid TileId { get; set; }
        public int AnimationId { get; set; }
        public int OffX { get; set; }
        public int OffY { get; set; }
        public int FrameWidth { get; set; }
        public int FrameHeight { get; set; }
        public MPQDT1Tile Tile { get; set; }
        public Rectangle Rect { get; set; }
        public ITexture Texture { get; set; }
    }
}
