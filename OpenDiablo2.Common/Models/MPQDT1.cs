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
        public UInt32 FileOffset { get; internal set; }
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
                stream.Seek(tileHeaderOffset + (96 * tileIndex), SeekOrigin.Begin);
                Tiles[tileIndex] = new MPQDT1Tile();
                var tile = Tiles[tileIndex];

                tile.Direction = br.ReadUInt32();
                tile.RoofHeight = br.ReadUInt16();
                tile.SoundIndex = br.ReadByte();
                tile.Animated = br.ReadByte() == 1;
                tile.Height = br.ReadInt32();
                tile.Width = br.ReadInt32();
                br.ReadBytes(4);
                tile.Orientation = br.ReadUInt32();
                tile.MainIndex = br.ReadUInt32();
                tile.SubIndex = br.ReadUInt32();
                tile.RarityOrFrameIndex = br.ReadUInt32();
                br.ReadBytes(4);
                for (int i = 0; i < 25; i++)
                    tile.SubTileFlags[i] = br.ReadByte();
                br.ReadBytes(7);
                tile.BlockHeadersPointer = br.ReadUInt32();
                tile.BlockDataLength = br.ReadUInt32();
                tile.NumberOfBlocks = br.ReadUInt32();
                br.ReadBytes(12);

                if (tile.BlockHeadersPointer == 0 || tile.Width == 0 || tile.Height == 0)
                {
                    tile.Blocks = new MPQDT1Block[0];
                    continue;
                }

                tile.Blocks = new MPQDT1Block[tile.NumberOfBlocks];
                for (int blockIndex = 0; blockIndex < tile.NumberOfBlocks; blockIndex++)
                {
                    stream.Seek(tile.BlockHeadersPointer + (20 * blockIndex), SeekOrigin.Begin);
                    tile.Blocks[blockIndex] = new MPQDT1Block();
                    var block = tile.Blocks[blockIndex];

                    block.PositionX = br.ReadUInt16();
                    block.PositionY = br.ReadUInt16();
                    br.ReadBytes(2);
                    block.GridX = br.ReadByte();
                    block.GridX = br.ReadByte();
                    block.Format = br.ReadUInt16();
                    block.Length = br.ReadUInt32();
                    br.ReadBytes(2);
                    block.FileOffset = br.ReadUInt32();
                    
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
                        // RLE block
                        /* TODO: BROKEN
                        var length = block.Length;
                        byte b1;
                        byte b2;
                        int x = 0;
                        int y = 0;
                        int width = (block.Format >> 8);
                        int height = (block.Format & 0xFF);
                        block.PixelData = new Int16[width * height];
                        while (length > 0)
                        {
                            b1 = br.ReadByte();
                            b2 = br.ReadByte();
                            length -= 2;
                            if (b1 != 0 || b2 != 0)
                            {
                                x += b1;
                                length -= b2;
                                while (b2 > 0)
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
                        */
                    }
                }
            }
        }
    }
}
