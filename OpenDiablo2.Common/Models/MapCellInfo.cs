using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Common.Models
{
    // Represents a single cell on a map
    public sealed class MapCellInfo
    {
        public Guid TileId;
        public int OffX;
        public int OffY;
        public int FrameWidth;
        public int FrameHeight;
        public Rectangle Rect;
        public ITexture Texture;
    }
}
