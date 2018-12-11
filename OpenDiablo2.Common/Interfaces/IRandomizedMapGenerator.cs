using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IRandomizedMapGenerator
    {
        void Generate(IMapInfo parentMap, Point location);
    }
}
