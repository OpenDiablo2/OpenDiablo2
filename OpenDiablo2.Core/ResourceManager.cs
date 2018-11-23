using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class ResourceManager : IResourceManager
    {
        private readonly IMPQProvider mpqProvider;
        private Dictionary<string, ImageSet> ImageSets = new Dictionary<string, ImageSet>();
        private Dictionary<string, MPQFont> MPQFonts = new Dictionary<string, MPQFont>();
        private Dictionary<string, Palette> Palettes = new Dictionary<string, Palette>();

        public ResourceManager(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;
        }

        public ImageSet GetImageSet(string resourcePath)
        {
            if (!ImageSets.ContainsKey(resourcePath))
                ImageSets[resourcePath] = ImageSet.LoadFromStream(mpqProvider.GetStream(resourcePath));

            return ImageSets[resourcePath];
        }

        public MPQFont GetMPQFont(string resourcePath)
        {
            if (!MPQFonts.ContainsKey(resourcePath))
                MPQFonts[resourcePath] = MPQFont.LoadFromStream(mpqProvider.GetStream($"{resourcePath}.DC6"), mpqProvider.GetStream($"{resourcePath}.tbl"));

            return MPQFonts[resourcePath];
        }

        public Palette GetPalette(string paletteFile)
        {
            if (!Palettes.ContainsKey(paletteFile))
            {
                var paletteNameParts = paletteFile.Split('\\');
                var paletteName = paletteNameParts[paletteNameParts.Count() - 2];
                Palettes[paletteFile] = Palette.LoadFromStream(mpqProvider.GetStream(paletteFile), paletteName);
            }

            return Palettes[paletteFile];

            
        }
    }
}
