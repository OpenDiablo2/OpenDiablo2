using System.Collections.Generic;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IPaletteProvider
    {
        Dictionary<string, Palette> PaletteTable { get; }
    }
}
