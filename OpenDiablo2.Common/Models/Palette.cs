using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public struct Palette
    {
        public string Name { get; set; }
        public UInt32[] Colors { get; set; }

        public static Palette LoadFromStream(Stream stream, string paletteName)
        {
            var result = new Palette
            {
                Name = paletteName,
                Colors = new UInt32[256]
            };

            var br = new BinaryReader(stream);
            for (var i = 0; i <= 255; i++)
            {
                var b = br.ReadByte();
                var g = br.ReadByte();
                var r = br.ReadByte();
                result.Colors[i] = ((UInt32)255 << 24) + ((UInt32)r << 16) + ((UInt32)g << 8) + b;
            }

            return result;
        }
    }
}
