using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQDT1Block
    {
        public UInt16 PositionX { get; internal set; }
        public UInt16 PositionY { get; internal set; }
        public byte GridX { get; internal set; }
        public byte GridY { get; internal set; }
        public UInt16 Format { get; internal set; }
        public UInt32 Length { get; internal set; }
        public Int32 FileOffset { get; internal set; }
        public Int16[] PixelData { get; internal set; }
    }

    public sealed class MPQDT1Tile
    {
        public UInt32 Direction { get; internal set; }
        public UInt16 RoofHeight { get; internal set; }
        public byte SoundIndex { get; internal set; }
        public bool Animated { get; internal set; }
        public int Height { get; internal set; }
        public int Width { get; internal set; }
        public UInt32 Orientation { get; internal set; }
        public UInt32 MainIndex { get; internal set; }
        public UInt32 SubIndex { get; internal set; }
        public UInt32 RarityOrFrameIndex { get; internal set; }// Frame index if animated floor tile
        public byte[] SubTileFlags { get; internal set; } = new byte[25];
        public UInt32 BlockHeadersPointer { get; internal set; } // Pointer to block headers for this tile
        public UInt32 BlockDataLength { get; internal set; } // Block headers + block data of this tile
        public UInt32 NumberOfBlocks { get; internal set; }
        public MPQDT1Block[] Blocks { get; internal set; }
    }

    public sealed class MPQDT1
    {
        public UInt32 NumberOfTiles { get; private set; }
        public MPQDT1Tile[] Tiles { get; private set; }

        public MPQDT1(Stream stream)
        {
            var br = new BinaryReader(stream);

            stream.Seek(268, SeekOrigin.Begin); // Skip useless header info
            NumberOfTiles = br.ReadUInt32();
            var tileHeaderOffset = br.ReadUInt32();
            stream.Seek(tileHeaderOffset, SeekOrigin.Begin);
            Tiles = new MPQDT1Tile[NumberOfTiles];

            for (int tileIndex = 0; tileIndex < NumberOfTiles; tileIndex++)
            {
                Tiles[tileIndex] = new MPQDT1Tile();
                Tiles[tileIndex].Direction = br.ReadUInt32();
                Tiles[tileIndex].RoofHeight = br.ReadUInt16();
                Tiles[tileIndex].SoundIndex = br.ReadByte();
                Tiles[tileIndex].Animated = br.ReadByte() == 1;
                Tiles[tileIndex].Height = br.ReadInt32();
                Tiles[tileIndex].Width = br.ReadInt32();
                br.ReadBytes(4);
                Tiles[tileIndex].Orientation = br.ReadUInt32();
                Tiles[tileIndex].MainIndex = br.ReadUInt32();
                Tiles[tileIndex].SubIndex = br.ReadUInt32();
                Tiles[tileIndex].RarityOrFrameIndex = br.ReadUInt32();
                br.ReadBytes(4);
                for (int i = 0; i < 25; i++)
                    Tiles[tileIndex].SubTileFlags[i] = br.ReadByte();
                br.ReadBytes(7);
                Tiles[tileIndex].BlockHeadersPointer = br.ReadUInt32();
                Tiles[tileIndex].BlockDataLength = br.ReadUInt32();
                Tiles[tileIndex].NumberOfBlocks = br.ReadUInt32();
                br.ReadBytes(12);
            }


            for (int tileIndex = 0; tileIndex < NumberOfTiles; tileIndex++)
            {
                var tile = Tiles[tileIndex];
                if (tile.BlockHeadersPointer == 0)
                {
                    tile.Blocks = new MPQDT1Block[0];
                    continue;
                }
                stream.Seek(tile.BlockHeadersPointer, SeekOrigin.Begin);

                tile.Blocks = new MPQDT1Block[tile.NumberOfBlocks];
                for (int blockIndex = 0; blockIndex < tile.NumberOfBlocks; blockIndex++)
                {
                    tile.Blocks[blockIndex] = new MPQDT1Block();
                    tile.Blocks[blockIndex].PositionX = br.ReadUInt16();
                    tile.Blocks[blockIndex].PositionY = br.ReadUInt16();
                    br.ReadBytes(2);
                    tile.Blocks[blockIndex].GridX = br.ReadByte();
                    tile.Blocks[blockIndex].GridX = br.ReadByte();
                    tile.Blocks[blockIndex].Format = br.ReadUInt16();
                    tile.Blocks[blockIndex].Length = br.ReadUInt32();
                    br.ReadBytes(2);
                    tile.Blocks[blockIndex].FileOffset = br.ReadInt32();
                }

                for (int blockIndex = 0; blockIndex < tile.NumberOfBlocks; blockIndex++)
                {
                    var block = tile.Blocks[blockIndex];
                    stream.Seek(tile.BlockHeadersPointer + block.FileOffset, SeekOrigin.Begin);

                    if (block.Format == 1)
                    {
                        // 3D isometric block
                        if (block.Length != 256)
                            throw new ApplicationException($"Expected exactly 256 bytes of data, but got {block.Length} instead!");

                        int x = 0;
                        int y = 0;
                        int n = 0;
                        int[] xjump = { 14, 12, 10, 8, 6, 4, 2, 0, 2, 4, 6, 8, 10, 12, 14 };
                        int[] nbpix = { 4, 8, 12, 16, 20, 24, 28, 32, 28, 24, 20, 16, 12, 8, 4 };
                        var length = 256;
                        block.PixelData = new Int16[32 * 16];
                        while (length > 0)
                        {
                            x = xjump[y];
                            n = nbpix[y];
                            length -= n;
                            while (n > 0)
                            {
                                block.PixelData[x + (y * 32)] = br.ReadByte();
                                x++;
                                n--;
                            }
                            y++;
                        }

                    }
                    else
                    {
                        // TODO: This doesn't work.. memory pointer issues?
                        continue;

                        // RLE block
                        var length = block.Length;
                        byte b1;
                        byte b2;
                        int x = 0;
                        int y = 0;
                        block.PixelData = new Int16[32 * 16];
                        while (length > 0)
                        {
                            b1 = br.ReadByte();
                            b2 = br.ReadByte();
                            length -= 2;
                            if (b1 != 0 || b2 != 0)
                            {
                                x += b1;
                                length -= b2;
                                while(b2 > 0)
                                {
                                    block.PixelData[x + (y * 32)] = br.ReadByte();
                                    br.ReadByte();
                                    x++;
                                    b2--;
                                }
                            }
                            else
                            {
                                x = 0;
                                y++;
                            }
                        }
                    }
                }
            }
        }
    }
}
