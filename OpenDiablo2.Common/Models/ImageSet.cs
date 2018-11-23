using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{

    public class ImageFrame : IDisposable
    {
        public UInt32 Flip;
        public UInt32 Width;
        public UInt32 Height;
        public Int32 OffsetX;
        public Int32 OffsetY;        // from bottom border, not up 
        public UInt32 Unknown;
        public UInt32 NextBlock;
        public UInt32 Length;
        public Int16[,] ImageData;

        public void Dispose()
        {
            ImageData = new Int16[0, 0];
        }

        public Color GetColor(int x, int y, Palette palette)
        {
            var index = ImageData[x, y];
            if (index == -1)
                return Color.Transparent;

            var color = palette.Colors[(int)index];

            return Color.FromArgb(255, color.R, color.G, color.B);
        }
    }

    public sealed class ImageSet : IDisposable
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private UInt32 version;
        private UInt32 unknown1;           // 01 00 00 00  ???
        private UInt32 unknown2;           // 00 00 00 00  ???
        private UInt32 termination;        // EE EE EE EE or CD CD CD CD  ???
        private UInt32[] framePointers;
        public ImageFrame[] Frames { get; private set; }

        public UInt32 Directions { get; private set; }
        public UInt32 FramesPerDirection { get; private set; }

        public static ImageSet LoadFromStream(Stream stream)
        {
            var br = new BinaryReader(stream);
            var result = new ImageSet
            {
                version = br.ReadUInt32(),
                unknown1 = br.ReadUInt32(),
                unknown2 = br.ReadUInt32(),
                termination = br.ReadUInt32(),
                Directions = br.ReadUInt32(),
                FramesPerDirection = br.ReadUInt32()
            };

            result.framePointers = new uint[result.Directions * result.FramesPerDirection];
            for (var i = 0; i < result.Directions * result.FramesPerDirection; i++)
                result.framePointers[i] = br.ReadUInt32();

            result.Frames = new ImageFrame[result.Directions * result.FramesPerDirection];
            for (var i = 0; i < result.Directions * result.FramesPerDirection; i++)
            {
                stream.Seek(result.framePointers[i], SeekOrigin.Begin);

                result.Frames[i] = new ImageFrame
                {
                    Flip = br.ReadUInt32(),
                    Width = br.ReadUInt32(),
                    Height = br.ReadUInt32(),
                    OffsetX = br.ReadInt32(),
                    OffsetY = br.ReadInt32(),
                    Unknown = br.ReadUInt32(),
                    NextBlock = br.ReadUInt32(),
                    Length = br.ReadUInt32()
                };

                result.Frames[i].ImageData = new Int16[result.Frames[i].Width, result.Frames[i].Height];
                for (int ty = 0; ty < result.Frames[i].Height; ty++)
                    for (int tx = 0; tx < result.Frames[i].Width; tx++)
                        result.Frames[i].ImageData[tx, ty] = -1;


                int x = 0;
                int y = (int)result.Frames[i].Height - 1;
                while (true)
                {
                    var b = br.ReadByte();
                    if (b == 0x80)
                    {
                        if (y == 0)
                            break;

                        y--;
                        x = 0;
                        continue;
                    }

                    if ((b & 0x80) > 0)
                    {
                        var transparentPixelsToWrite = b & 0x7F;
                        for (int p = 0; p < transparentPixelsToWrite; p++)
                        {
                            result.Frames[i].ImageData[x++, y] = -1;
                        }
                        continue;
                    }

                    for (int p = 0; p < b; p++)
                    {
                        result.Frames[i].ImageData[x++, y] = br.ReadByte();
                    }

                }
            }
            return result;
        }

        public void Dispose()
        {
        }
    }
}
