using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Models;

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
