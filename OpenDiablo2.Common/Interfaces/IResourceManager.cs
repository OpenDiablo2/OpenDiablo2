using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IResourceManager
    {
        ImageSet GetImageSet(string resourcePath);
        MPQFont GetMPQFont(string resourcePath);
        MPQDS1 GetMPQDS1(string resourcePath, LevelPreset level, LevelDetail levelDetail, LevelType levelType);
        MPQDT1 GetMPQDT1(string resourcePath);
        Palette GetPalette(string paletteName);
    }
}
