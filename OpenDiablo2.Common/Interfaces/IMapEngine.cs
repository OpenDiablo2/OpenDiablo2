using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMapEngine
    {
        int FocusedMobId { get; set; }
        PointF CameraLocation { get; set; } 
        void Update(long ms);
        void Render();
    }
}
