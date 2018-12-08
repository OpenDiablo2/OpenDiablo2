using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Exceptions;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQDCC
    {
        public sealed class PixelBufferEntry
        {
            public byte[] Value;
            public int Frame;
            public int FrameCellIndex;
        }

        public sealed class Cell
        {
            public int Width;
            public int Height;
            public int XOffset;
            public int YOffset;

            public int LastWidth;
            public int LastHeight;
            public int LastXOffset;
            public int LastYOffset;

            public byte[] PixelData;
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
                Width = (int)bits.GetBits(direction.WidthBits);
                Height = (int)bits.GetBits(direction.HeightBits);
                XOffset = bits.GetSignedBits(direction.XOffsetBits);
                YOffset = bits.GetSignedBits(direction.YOffsetBits);
                NumberOfOptionalBytes = (int)bits.GetBits(direction.OptionalDataBits);
                NumberOfCodedBytes = (int)bits.GetBits(direction.CodedBytesBits);
                FrameIsBottomUp = bits.GetBit() == 1;

                Box = new Rectangle(
                    XOffset,
                    YOffset - Height + 1,
                    Width,
                    Height
                );

            }

            public void CalculateCells(MPQDCCDirection direction)
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
            public byte[] PaletteEntries { get; private set; }
            public Rectangle Box { get; private set; }
            public Cell[] Cells { get; private set; }
            public int HorizontalCellCount { get; private set; }
            public int VerticalCellCount { get; private set; }
            public PixelBufferEntry[] PixelBuffer { get; private set; }

            private static readonly byte[] crazyBitTable = { 0, 1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 24, 26, 28, 30, 32 };
            public MPQDCCDirection(BitMuncher bm, MPQDCC file)
            {
                OutSizeCoded = (int)bm.GetUInt32();
                CompressionFlags = (int)bm.GetBits(2);
                Variable0Bits = crazyBitTable[bm.GetBits(4)];
                WidthBits = crazyBitTable[bm.GetBits(4)];
                HeightBits = crazyBitTable[bm.GetBits(4)];
                XOffsetBits = crazyBitTable[bm.GetBits(4)];
                YOffsetBits = crazyBitTable[bm.GetBits(4)];
                OptionalDataBits = crazyBitTable[bm.GetBits(4)];
                CodedBytesBits = crazyBitTable[bm.GetBits(4)];

                Frames = new MPQDCCDirectionFrame[file.FramesPerDirection];

                var minx = long.MaxValue;
                var miny = long.MaxValue;
                var maxx = long.MinValue;
                var maxy = long.MinValue;
                // Load the frame headers
                for (var frameIdx = 0; frameIdx < file.FramesPerDirection; frameIdx++)
                {
                    Frames[frameIdx] = new MPQDCCDirectionFrame(bm, this);

                    minx = Math.Min(Frames[frameIdx].Box.X, minx);
                    miny = Math.Min(Frames[frameIdx].Box.Y, miny);
                    maxx = Math.Max(Frames[frameIdx].Box.Right, maxx);
                    maxy = Math.Max(Frames[frameIdx].Box.Bottom, maxy);
                }

                Box = new Rectangle
                {
                    X = (int)minx,
                    Y = (int)miny,
                    Width = (int)(maxx - minx),
                    Height = (int)(maxy - miny)
                };

                if (OptionalDataBits > 0)
                    throw new OpenDiablo2Exception("Optional bits in DCC data is not currently supported.");

                if ((CompressionFlags & 0x2) > 0)
                    EqualCellsBitstreamSize = (int)bm.GetBits(20);

                PixelMaskBitstreamSize = (int)bm.GetBits(20);

                if ((CompressionFlags & 0x1) > 0)
                {
                    EncodingTypeBitsreamSize = (int)bm.GetBits(20);
                    RawPixelCodesBitstreamSize = (int)bm.GetBits(20);
                }


                // PixelValuesKey 
                var paletteEntries = new List<bool>();
                for (var i = 0; i < 256; i++)
                    paletteEntries.Add(bm.GetBit() != 0);

                PaletteEntries = new byte[paletteEntries.Count(x => x == true)];
                var paletteOffset = 0;
                for (var i = 0; i < 256; i++)
                {
                    if (!paletteEntries[i])
                        continue;

                    PaletteEntries[paletteOffset++] = (byte)i;
                }


                // HERE BE GIANTS:
                // Because of the way this thing mashes bits together, BIT offset matters
                // here. For example, if you are on byte offset 3, bit offset 6, and
                // the EqualCellsBitstreamSize is 20 bytes, then the next bit stream
                // will be located at byte 23, bit offset 6!

                var equalCellsBitstream = new BitMuncher(bm);
                bm.SkipBits(EqualCellsBitstreamSize);

                var pixelMaskBitstream = new BitMuncher(bm);
                bm.SkipBits(PixelMaskBitstreamSize);

                var encodingTypeBitsream = new BitMuncher(bm);
                bm.SkipBits(EncodingTypeBitsreamSize);

                var rawPixelCodesBitstream = new BitMuncher(bm);
                bm.SkipBits(RawPixelCodesBitstreamSize);

                var pixelCodeandDisplacement = new BitMuncher(bm);

                // Calculate the cells for the direction
                CaculateCells();

                // Caculate the cells for each of the frames
                foreach (var frame in Frames)
                    frame.CalculateCells(this);

                // Fill in the pixel buffer
                FillPixelBuffer(pixelCodeandDisplacement, equalCellsBitstream, pixelMaskBitstream, encodingTypeBitsream, rawPixelCodesBitstream);

                // Generate the actual frame pixel data
                GenerateFrames(pixelCodeandDisplacement);

                // Verify that everything we expected to read was actually read (sanity check)...
                if (equalCellsBitstream.BitsRead != EqualCellsBitstreamSize)
                    throw new OpenDiablo2Exception("Did not read the correct number of bits!");

                if (pixelMaskBitstream.BitsRead != PixelMaskBitstreamSize)
                    throw new OpenDiablo2Exception("Did not read the correct number of bits!");

                if (encodingTypeBitsream.BitsRead != EncodingTypeBitsreamSize)
                    throw new OpenDiablo2Exception("Did not read the correct number of bits!");

                if (rawPixelCodesBitstream.BitsRead != RawPixelCodesBitstreamSize)
                    throw new OpenDiablo2Exception("Did not read the correct number of bits!");

                bm.SkipBits(pixelCodeandDisplacement.BitsRead);
            }

            private void GenerateFrames(BitMuncher pcd)
            {
                var pbIdx = 0;

                foreach (var cell in Cells)
                {
                    cell.LastWidth = -1;
                    cell.LastHeight = -1;
                    cell.PixelData = new byte[cell.Width * cell.Height];
                }

                

                    var frameIndex = -1;
                foreach (var frame in Frames)
                {
                    frameIndex++;
                    var numberOfCells = frame.HorizontalCellCount * frame.VerticalCellCount;

                    var c = -1;
                    foreach (var cell in frame.Cells)
                    {
                        c++;
                        var cellX = cell.XOffset / 4;
                        var cellY = cell.YOffset / 4;
                        var cellIndex = cellX + (cellY * HorizontalCellCount);
                        var bufferCell = Cells[cellIndex];

                        var pbe = PixelBuffer[pbIdx];
                        if ((pbe.Frame != frameIndex) || (pbe.FrameCellIndex != c))
                        {
                            // This buffer cell has an EqualCell bit set to 1, so copy the frame cell or clear it

                            if ((cell.Width != bufferCell.LastWidth) || (cell.Height != bufferCell.LastHeight))
                            {
                                // Different sizes
                                /// TODO: Clear the pixels of the frame cell
                                for (var i = 0; i < bufferCell.PixelData.Length; i++)
                                    bufferCell.PixelData[i] = 0x00;

                            }
                            else
                            {
                                // Same sizes

                                // Copy the old frame cell into the new position
                                // blit(dir->bmp, dir->bmp, buff_cell->last_x0, buff_cell->last_y0, cell->x0, cell->y0, cell->w, cell->h );
                                for (var i = 0; i < bufferCell.PixelData.Length; i++)
                                    bufferCell.PixelData[i] = cell.PixelData[i];


                                bufferCell.LastWidth = cell.LastWidth;
                                bufferCell.LastHeight = cell.LastHeight;
                                // Copy it again into the final frame image
                                // blit(cell->bmp, frm_bmp, 0, 0, cell->x0, cell->y0, cell->w, cell->h );
                            }
                        }
                        else
                        {
                            if (pbe.Value[0] == pbe.Value[1])
                            {
                                // Clear the frame
                            }
                            else
                            {
                                // Fill the frame cell with the pixels
                                var bitsToRead = (pbe.Value[1] == pbe.Value[2]) ? 1 : 2;
                                cell.PixelData = new byte[cell.Width * cell.Height];

                                for (var y = 0; y < cell.Height; y++)
                                {
                                    for (var x = 0; x < cell.Width; x++)
                                    {
                                        var paletteIndex = pcd.GetBits(bitsToRead);
                                        cell.PixelData[x + (y * cell.Width)] = pbe.Value[paletteIndex];
                                    }
                                }
                            }

                            // Copy the frame cell into the frame
                            //blit(cell->bmp, frm_bmp, 0, 0, cell->x0, cell->y0, cell->w, cell->h );
                            pbIdx++;
                        }
                    }
                }
            }

            private static readonly int[] pixelMaskLookup = new int[] { 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4 };
            private void FillPixelBuffer(BitMuncher pcd, BitMuncher ec, BitMuncher pm, BitMuncher et, BitMuncher rp)
            {
                UInt32 lastPixel = 0;

                var maxCellX = Frames.Sum(x => x.HorizontalCellCount);
                var maxCellY = Frames.Sum(x => x.VerticalCellCount);

                PixelBuffer = new PixelBufferEntry[maxCellX * maxCellY];
                for (var i = 0; i < maxCellX * maxCellY; i++)
                    PixelBuffer[i] = new PixelBufferEntry { Frame = -1, FrameCellIndex = -1, Value = new byte[4] };

                var cellBuffer = new PixelBufferEntry[HorizontalCellCount * VerticalCellCount];

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
                            var currentCell = (originCellX + cellX) + (currentCellY * HorizontalCellCount);
                            var nextCell = false;
                            var tmp = 0;
                            if (cellBuffer[currentCell] != null)
                            {
                                if (EqualCellsBitstreamSize > 0)
                                    tmp = (int)ec.GetBit();
                                else
                                    tmp = 0;

                                if (tmp == 0)
                                    pixelMask = pm.GetBits(4);
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
                            int encodingType = 0;
                            if ((numberOfPixelBits != 0) && (EncodingTypeBitsreamSize > 0))
                                encodingType = (int)et.GetBit();
                            else
                                encodingType = 0;

                            int decodedPixel = 0;
                            for (int i = 0; i < numberOfPixelBits; i++)
                            {
                                if (encodingType != 0)
                                {
                                    pixelStack[i] = rp.GetBits(8);
                                }
                                else
                                {
                                    pixelStack[i] = lastPixel;
                                    var pixelDisplacement = pcd.GetBits(4);
                                    pixelStack[i] += pixelDisplacement;
                                    while (pixelDisplacement == 15)
                                    {
                                        pixelDisplacement = pcd.GetBits(4);
                                        pixelStack[i] += pixelDisplacement;
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
                            newEntry.FrameCellIndex = cellX + (cellY * frame.HorizontalCellCount);
                        }
                    }
                }


                // Convert the palette entry index into actual palette entries
                for (var i = 0; i < pbIndex; i++)
                {
                    for (var x = 0; x < 4; x++)
                    {
                        var y = PixelBuffer[i].Value[x];
                        PixelBuffer[i].Value[x] = PaletteEntries[y];
                    }
                }
            }

            private void CaculateCells()
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
        public int FramesPerDirection { get; private set; }
        public MPQDCCDirection[] Directions { get; private set; }

        public MPQDCC(byte[] data, Palette palette)
        {
            var bm = new BitMuncher(data);
            Signature = bm.GetByte();
            if (Signature != 0x74)
                throw new OpenDiablo2Exception("Signature expected to be 0x74 but it is not.");

            Version = bm.GetByte();
            NumberOfDirections = bm.GetByte();
            FramesPerDirection = bm.GetInt32();

            if (bm.GetInt32() != 1)
                throw new OpenDiablo2Exception("This value isn't 1. It has to be 1.");

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
