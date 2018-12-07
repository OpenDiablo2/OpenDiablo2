using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using System.Drawing;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQDCC
    {
        public class PixelBufferEntry
        {
            public byte[] Value;
            public int Frame;
            public int FrameCellIndex;
        }

        public struct Cell
        {
            public int Width;
            public int Height;
            public int XOffset;
            public int YOffset;
        }

        public sealed class MPQDCCDirectionFrame
        {
            public int Width { get; private set; }
            public int Height { get; private set; }
            public int XOffset { get; private set; }
            public int YOffset { get; private set; }
            public int NumberOfOptionalBytes { get; private set; }
            public int NumberOfCodedBytes { get; private set; }
            public bool FrameIsBottomUp { get; private set; }
            public Rectangle Box { get; private set; }
            public Cell[] Cells { get; private set; }
            public int HorizontalCellCount { get; private set; }
            public int VerticalCellCount { get; private set; }



            public MPQDCCDirectionFrame(BitMuncher bits, MPQDCCDirection direction)
            {
                var variable0 = bits.GetBits(direction.Variable0Bits);
                Width = bits.GetBits(direction.WidthBits);
                Height = bits.GetBits(direction.HeightBits);
                XOffset = bits.GetSignedBits(direction.XOffsetBits);
                YOffset = bits.GetSignedBits(direction.YOffsetBits);
                NumberOfOptionalBytes = bits.GetBits(direction.OptionalDataBits);
                NumberOfCodedBytes = bits.GetBits(direction.CodedBytesBits);
                FrameIsBottomUp = bits.GetBit() == 1;

                Box = new Rectangle(
                    XOffset,
                    YOffset - Height + 1,
                    Width,
                    Height
                );

            }

            public void MakeCells(MPQDCCDirection direction)
            {
                var w = 4 - ((Box.Left - direction.Box.Left) % 4); // Width of the first column (in pixels)
                if ((Width - w) <= 1)
                    HorizontalCellCount = 1;
                else
                {
                    var tmp = Width - w - 1;
                    HorizontalCellCount = 2 + (tmp / 4);
                    if ((tmp % 4) == 0)
                        HorizontalCellCount--;
                }


                var h = 4 - ((Box.Top - direction.Box.Top) % 4); // Height of the first column (in pixels)
                if ((Height - h) <= 1)
                    VerticalCellCount = 1;
                else
                {
                    var tmp = Height - h - 1;
                    VerticalCellCount = 2 + (tmp / 4);
                    if ((tmp % 4) == 0)
                        VerticalCellCount--;
                }

                Cells = new Cell[HorizontalCellCount * VerticalCellCount];

                // Calculate the cell widths and heights
                var cellWidths = new int[HorizontalCellCount];
                if (HorizontalCellCount == 1)
                    cellWidths[0] = Width;
                else
                {
                    cellWidths[0] = w;
                    for (var i = 1; i < (HorizontalCellCount - 1); i++)
                        cellWidths[i] = 4;
                    cellWidths[HorizontalCellCount - 1] = Width - w - (4 * (HorizontalCellCount - 2));
                }

                var cellHeights = new int[VerticalCellCount];
                if (VerticalCellCount == 1)
                    cellHeights[0] = Height;
                else
                {
                    cellHeights[0] = h;
                    for (var i = 1; i < (VerticalCellCount - 1); i++)
                        cellHeights[i] = 4;
                    cellHeights[VerticalCellCount - 1] = Height - h - (4 * (VerticalCellCount - 2));
                }

                Cells = new Cell[HorizontalCellCount * VerticalCellCount];
                var offsetY = Box.Top - direction.Box.Top;
                for (var y = 0; y < VerticalCellCount; y++)
                {
                    var offsetX = Box.Left - direction.Box.Left;
                    for (var x = 0; x < HorizontalCellCount; x++)
                    {
                        Cells[x + (y * HorizontalCellCount)] = new Cell
                        {
                            XOffset = offsetX,
                            YOffset = offsetY,
                            Width = cellWidths[x],
                            Height = cellHeights[y]
                        };
                        offsetX += cellWidths[x];
                    }
                    offsetY += cellHeights[y];
                }
            }
        }
        public sealed class MPQDCCDirection
        {
            const int DCC_MAX_PB_ENTRY = 85000; // But why is this the magic number?
            public int OutSizeCoded { get; private set; }
            public int CompressionFlags { get; private set; }
            public int Variable0Bits { get; private set; }
            public int WidthBits { get; private set; }
            public int HeightBits { get; private set; }
            public int XOffsetBits { get; private set; }
            public int YOffsetBits { get; private set; }
            public int OptionalDataBits { get; private set; }
            public int CodedBytesBits { get; private set; }
            public int EqualCellsBitstreamSize { get; private set; }
            public int PixelMaskBitstreamSize { get; private set; }
            public int EncodingTypeBitsreamSize { get; private set; }
            public int RawPixelCodesBitstreamSize { get; private set; }
            public MPQDCCDirectionFrame[] Frames { get; private set; }
            public int[] PaletteEntries { get; private set; }
            public Rectangle Box { get; private set; }
            public Cell[] Cells { get; private set; }
            public int HorizontalCellCount { get; private set; }
            public int VerticalCellCount { get; private set; }
            public PixelBufferEntry[] PixelBuffer { get; private set; }

            private static readonly byte[] crazyBitTable = { 0, 1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 24, 26, 28, 30, 32 };
            public MPQDCCDirection(BitMuncher bm, MPQDCC file)
            {
                OutSizeCoded = bm.GetInt32();
                CompressionFlags = bm.GetBits(2);
                Variable0Bits = crazyBitTable[bm.GetBits(4)];
                WidthBits = crazyBitTable[bm.GetBits(4)];
                HeightBits = crazyBitTable[bm.GetBits(4)];
                XOffsetBits = crazyBitTable[bm.GetBits(4)];
                YOffsetBits = crazyBitTable[bm.GetBits(4)];
                OptionalDataBits = crazyBitTable[bm.GetBits(4)];
                CodedBytesBits = crazyBitTable[bm.GetBits(4)];

                Frames = new MPQDCCDirectionFrame[file.NumberOfFrames];

                // Load the frame headers
                for (var frameIdx = 0; frameIdx < file.NumberOfFrames; frameIdx++)
                    Frames[frameIdx] = new MPQDCCDirectionFrame(bm, this);

                Box = new Rectangle
                {
                    X = Frames.Min(z => z.Box.X),
                    Y = Frames.Min(z => z.Box.Y),
                    Width = Frames.Max(z => z.Box.Right - z.Box.Left),
                    Height = Frames.Max(z => z.Box.Bottom - z.Box.Top)
                };

                if (OptionalDataBits > 0)
                    throw new ApplicationException("Optional bits in DCC data is not currently supported.");

                if ((CompressionFlags & 0x2) > 0)
                    EqualCellsBitstreamSize = bm.GetBits(20);

                PixelMaskBitstreamSize = bm.GetBits(20);

                if ((CompressionFlags & 0x1) > 0)
                {
                    EncodingTypeBitsreamSize = bm.GetBits(20);
                    RawPixelCodesBitstreamSize = bm.GetBits(20);
                }


                // PixelValuesKey 
                var paletteEntries = new List<bool>();
                for (var i = 0; i < 256; i++)
                    paletteEntries.Add(bm.GetBit() != 0);

                PaletteEntries = new int[paletteEntries.Count(x => x == true)];
                var paletteOffset = 0;
                for (var i = 0; i < 256; i++)
                {
                    if (!paletteEntries[i])
                        continue;

                    PaletteEntries[paletteOffset++] = i;
                }



                var equalCellsBitstream = new BitMuncher(bm);
                bm.SkipBits(EqualCellsBitstreamSize);

                var pixelMaskBitstream = new BitMuncher(bm);
                bm.SkipBits(PixelMaskBitstreamSize);

                var encodingTypeBitsream = new BitMuncher(bm);
                bm.SkipBits(EncodingTypeBitsreamSize);

                var rawPixelCodesBitstream = new BitMuncher(bm);
                bm.SkipBits(RawPixelCodesBitstreamSize);

                var pixelCodeandDisplacement = new BitMuncher(bm);

                CalculateCellOffsets();

                foreach (var frame in Frames)
                    frame.MakeCells(this);

                FillPixelBuffer(pixelCodeandDisplacement, equalCellsBitstream, pixelMaskBitstream, encodingTypeBitsream, rawPixelCodesBitstream);

                if (equalCellsBitstream.BitsRead != EqualCellsBitstreamSize)
                    throw new ApplicationException("Did not read the correct number of bits!");

                if (pixelMaskBitstream.BitsRead != PixelMaskBitstreamSize)
                    throw new ApplicationException("Did not read the correct number of bits!");

                if (encodingTypeBitsream.BitsRead != EncodingTypeBitsreamSize)
                    throw new ApplicationException("Did not read the correct number of bits!");

                if (rawPixelCodesBitstream.BitsRead != RawPixelCodesBitstreamSize)
                    throw new ApplicationException("Did not read the correct number of bits!");


                bm.SkipBits(pixelCodeandDisplacement.BitsRead);
            }

            private static readonly int[] pixelMaskLookup = new int[] { 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4 };
            private void FillPixelBuffer(BitMuncher pcd, BitMuncher ec, BitMuncher pm, BitMuncher et, BitMuncher rp)
            {
                UInt32 lastPixel = 0;
                
                PixelBuffer = new PixelBufferEntry[DCC_MAX_PB_ENTRY];
                for (var i = 0; i < DCC_MAX_PB_ENTRY; i++)
                    PixelBuffer[i] = new PixelBufferEntry { Frame = -1, FrameCellIndex = -1, Value = new byte[4] };

                var cellBuffer = new PixelBufferEntry[Cells.Length];

                var frameIndex = -1;
                var pbIndex = -1;
                UInt32 pixelMask = 0x00;
                foreach (var frame in Frames)
                {
                    frameIndex++;
                    var originCellX = (frame.Box.Left - Box.Left) / 4;
                    var originCellY = (frame.Box.Top - Box.Top) / 4;
                    var frameCellIndex = 0;
                    for (var cellY = 0; cellY < frame.VerticalCellCount; cellY++)
                    {
                        var currentCellY = cellY + originCellY;
                        for (var cellX = 0; cellX < frame.HorizontalCellCount; cellX++, frameCellIndex++)
                        {
                            var currentCell = (originCellX + cellX) + (currentCellY * frame.HorizontalCellCount);
                            var nextCell = false;
                            var tmp = 0;
                            if (cellBuffer[currentCell] != null)
                            {
                                if (EqualCellsBitstreamSize > 0)
                                    tmp = ec.GetBit();

                                if (tmp == 0)
                                    pixelMask = (UInt32)pm.GetBits(4);
                                else
                                    nextCell = true;
                            }
                            else pixelMask = 0x0F;

                            if (nextCell)
                                continue;

                            // Decode the pixels
                            var pixelStack = new UInt32[4];
                            lastPixel = 0;
                            int numberOfPixelBits = pixelMaskLookup[pixelMask];
                            var encodingType = ((numberOfPixelBits != 0) && (EncodingTypeBitsreamSize > 0))
                                ? et.GetBit()
                                : 0;
                            int decodedPixel = 0;
                            for (int i = 0; i < numberOfPixelBits; i++)
                            {
                                if (encodingType != 0)
                                {
                                    pixelStack[i] = (UInt32)rp.GetBits(8);
                                }
                                else
                                {
                                    pixelStack[i] = lastPixel;
                                    var pixelDisplacement = pcd.GetBits(4);
                                    pixelStack[i] += (UInt32)pixelDisplacement;
                                    while (pixelDisplacement == 15)
                                    {
                                        pixelDisplacement = pcd.GetBits(4);
                                        pixelStack[i] += (UInt32)pixelDisplacement;
                                    }
                                }
                                if (pixelStack[i] == lastPixel)
                                {
                                    pixelStack[i] = 0;
                                    i = numberOfPixelBits; // Just break here....
                                }
                                else
                                {
                                    lastPixel = pixelStack[i];
                                    decodedPixel++;
                                }
                            }

                            var oldEntry = cellBuffer[currentCell];
                            pbIndex++;
                            var newEntry = PixelBuffer[pbIndex];
                            var curIdx = decodedPixel - 1;
                            
                            for (int i = 0; i < 4; i++)
                            {
                                if ((pixelMask & (1 << i)) != 0)
                                {
                                    if (curIdx >= 0)
                                        newEntry.Value[i] = (byte)pixelStack[curIdx--];
                                    else
                                        newEntry.Value[i] = 0;
                                }
                                else
                                    newEntry.Value[i] = oldEntry.Value[i];
                            }

                            cellBuffer[currentCell] = newEntry;
                            newEntry.Frame = frameIndex;
                            newEntry.FrameCellIndex = cellX + (cellY * HorizontalCellCount);
                        }
                    }
                }
            }

            private void CalculateCellOffsets()
            {
                // Calculate the number of vertical and horizontal cells we need
                HorizontalCellCount = 1 + (Box.Width - 1) / 4;
                VerticalCellCount = 1 + (Box.Height - 1) / 4;

                // Calculate the cell widths
                var cellWidths = new int[HorizontalCellCount];
                if (HorizontalCellCount == 1)
                    cellWidths[0] = Box.Width;
                else
                {
                    for (var i = 0; i < HorizontalCellCount - 1; i++)
                        cellWidths[i] = 4;
                    cellWidths[HorizontalCellCount - 1] = Box.Width - (4 * (HorizontalCellCount - 1));
                }

                // Calculate the cell heights
                var cellHeights = new int[VerticalCellCount];
                if (VerticalCellCount == 1)
                    cellHeights[0] = Box.Height;
                else
                {
                    for (var i = 0; i < VerticalCellCount - 1; i++)
                        cellHeights[i] = 4;
                    cellHeights[VerticalCellCount - 1] = Box.Height - (4 * (VerticalCellCount - 1));
                }

                // Set the cell widths and heights in the cell buffer
                Cells = new Cell[VerticalCellCount * HorizontalCellCount];
                var yOffset = 0;
                for (var y = 0; y < VerticalCellCount; y++)
                {
                    var xOffset = 0;
                    for (var x = 0; x < HorizontalCellCount; x++)
                    {
                        Cells[x + (y * HorizontalCellCount)] = new Cell
                        {
                            Width = cellWidths[x],
                            Height = cellHeights[y],
                            XOffset = xOffset,
                            YOffset = yOffset
                        };
                        xOffset += 4;
                    }
                    yOffset += 4;
                }
            }
        }



        public int Signature { get; private set; }
        public int Version { get; private set; }
        public int NumberOfDirections { get; private set; }
        public int NumberOfFrames { get; private set; }
        public MPQDCCDirection[] Directions { get; private set; }

        public MPQDCC(byte[] data, Palette palette)
        {
            var bm = new BitMuncher(data);
            Signature = bm.GetByte();
            if (Signature != 0x74)
                throw new ApplicationException("Signature expected to be 0x74 but it is not.");

            Version = bm.GetByte();
            NumberOfDirections = bm.GetByte();
            NumberOfFrames = bm.GetInt32();

            if (bm.GetInt32() != 1)
                throw new ApplicationException("This value isn't 1. It has to be 1.");

            var totalSizeCoded = bm.GetInt32();
            var directionOffsets = new int[NumberOfDirections];
            for (var i = 0; i < NumberOfDirections; i++)
                directionOffsets[i] = bm.GetInt32();

            Directions = new MPQDCCDirection[NumberOfDirections];
            for (var i = 0; i < NumberOfDirections; i++)
            {
                Directions[i] = new MPQDCCDirection(new BitMuncher(data, directionOffsets[i] * 8), this);
            }

        }
    }

}
