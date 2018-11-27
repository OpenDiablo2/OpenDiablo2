using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMapEngine
    {
        PointF CameraLocation { get; set; } 
        void Update(long ms);
        void Render();
        void NotifyMapChanged();
    }
}
