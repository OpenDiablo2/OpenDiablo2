using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMapRenderer
    {
        Guid FocusedPlayerId { get; set; }
        int CameraOffset { get; set; }
        PointF CameraLocation { get; set; } 
        void Update(long ms);
        void Render();
        PointF GetCellPositionAt(int x, int y);
    }
}
