using System;
using System.IO;
using System.Linq;

namespace OpenDiablo2.Common.Models
{

    public sealed class ImageFrame : IDisposable
    {
        public UInt32 Flip { get; internal set; }
        public UInt32 Width { get; internal set; }
        public UInt32 Height { get; internal set; }
        public Int32 OffsetX { get; internal set; }
        public Int32 OffsetY { get; internal set; }        // from bottom border, not up 
        public UInt32 Unknown { get; internal set; }
        public UInt32 NextBlock { get; internal set; }
        public UInt32 Length { get; internal set; }
        public Int16[] ImageData { get; internal set; }

        public void Dispose()
        {
            ImageData = null;
        }

        public UInt32 GetColor(int x, int y, Palette palette)
        {
            var i = x + (y * Width);
            if (i >= ImageData.Length)
                return 0;

            var index = ImageData[i];
            if (index == -1)
                return 0;

            return palette.Colors[index];
        }
    }

    public sealed class ImageSet : IDisposable
    {
        private UInt32[] framePointers;
        public ImageFrame[] Frames { get; private set; }

        public UInt32 Directions { get; private set; }
        public UInt32 FramesPerDirection { get; private set; }

        public static ImageSet LoadFromStream(Stream stream)
        {
            var br = new BinaryReader(stream);
            br.ReadUInt32(); // Version
            br.ReadUInt32(); // Unknown 1
            br.ReadUInt32(); // Unknown 2
            br.ReadUInt32(); // Termination

            var result = new ImageSet
            {
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

                result.Frames[i].ImageData = new Int16[result.Frames[i].Width * result.Frames[i].Height];
                for (int ty = 0; ty < result.Frames[i].Height; ty++)
                    for (int tx = 0; tx < result.Frames[i].Width; tx++)
                        result.Frames[i].ImageData[tx + (ty * result.Frames[i].Width)] = -1;


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

                        Enumerable.Repeat<Int16>(-1, transparentPixelsToWrite).ToArray()
                            .CopyTo(result.Frames[i].ImageData, x + (y * result.Frames[i].Width));
                        x += transparentPixelsToWrite;

                        continue;
                    }


                    br.ReadBytes(b).CopyTo(result.Frames[i].ImageData, x + (y * result.Frames[i].Width));
                    x += b;



                }
            }
            return result;
        }

        public void Dispose()
        {
        }
    }
}
