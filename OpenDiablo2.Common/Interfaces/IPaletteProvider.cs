using OpenDiablo2.Common.Models;
using System.Collections.Generic;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IPaletteProvider
    {
        Dictionary<string, Palette> PaletteTable { get; }
    }
}
