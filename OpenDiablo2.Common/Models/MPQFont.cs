using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQFont
    {
        public ImageSet FontImageSet;
        public Dictionary<byte, Size> CharacterMetric = new Dictionary<byte, Size>();

        public static MPQFont LoadFromStream(Stream fontStream, Stream tableStream)
        {
            var result = new MPQFont
            {
                FontImageSet = ImageSet.LoadFromStream(fontStream)
            };

            var br = new BinaryReader(tableStream);
            var wooCheck = Encoding.UTF8.GetString(br.ReadBytes(4));
            if (wooCheck != "Woo!")
                throw new ApplicationException("Error loading font. Missing the Woo!");
            br.ReadBytes(8);

            while (tableStream.Position < tableStream.Length)
            {
                br.ReadBytes(3);
                var size = new Size(br.ReadByte(), br.ReadByte());
                br.ReadBytes(3);
                var charCode = br.ReadByte();
                result.CharacterMetric[charCode] = size;
                br.ReadBytes(5);

            }

            return result;
        }
    }
}
