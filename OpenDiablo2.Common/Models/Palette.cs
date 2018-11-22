using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public struct PaletteEntry
    {
        public int R;
        public int G;
        public int B;
    }

    public struct Palette
    {
        public string Name { get; set; }
        public PaletteEntry[] Colors;

        public static Palette LoadFromStream(Stream stream)
        {
            var result = new Palette
            {
                Colors = new PaletteEntry[256]
            };

            var br = new BinaryReader(stream);
            for (var i = 0; i <= 255; i++)
                result.Colors[i] = new PaletteEntry
                {
                    B = br.ReadByte(),
                    G = br.ReadByte(),
                    R = br.ReadByte()
                };

            return result;
        }
    }
}
