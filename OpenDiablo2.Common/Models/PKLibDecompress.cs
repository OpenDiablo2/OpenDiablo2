//
// MpqLibDecompress.cs
//
// Authors:
//		Foole (fooleau@gmail.com)
//
// (C) 2006 Foole (fooleau@gmail.com)
// Based on code from StormLib by Ladislav Zezula
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
// 
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
using System;
using System.IO;

namespace OpenDiablo2.Common.Models
{
    enum CompressionType
    {
        Binary = 0,
        Ascii = 1
    }
    /// <summary>
    /// A decompressor for PKLib implode/explode
    /// </summary>
    public class PKLibDecompress
    {
        private readonly BitStream _bitstream;
        private readonly CompressionType _compressionType;
        private readonly int _dictSizeBits;  // Dictionary size in bits

        private static readonly byte[] sPosition1;
        private static readonly byte[] sPosition2;

        private static readonly byte[] sLenBits =
        {
            3, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 7, 7
        };

        private static readonly byte[] sLenCode =
        {
            5, 3, 1, 6, 10, 2, 12, 20, 4, 24, 8, 48, 16, 32, 64, 0
        };

        private static readonly byte[] sExLenBits =
        {
            0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8
        };

        private static readonly UInt16[] sLenBase =
        {
            0x0000, 0x0001, 0x0002, 0x0003, 0x0004, 0x0005, 0x0006, 0x0007,
            0x0008, 0x000A, 0x000E, 0x0016, 0x0026, 0x0046, 0x0086, 0x0106
        };

        private static readonly byte[] sDistBits =
        {
            2, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6,
            6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
            7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
            8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8
        };

        private static readonly byte[] sDistCode =
        {
            0x03, 0x0D, 0x05, 0x19, 0x09, 0x11, 0x01, 0x3E, 0x1E, 0x2E, 0x0E, 0x36, 0x16, 0x26, 0x06, 0x3A,
            0x1A, 0x2A, 0x0A, 0x32, 0x12, 0x22, 0x42, 0x02, 0x7C, 0x3C, 0x5C, 0x1C, 0x6C, 0x2C, 0x4C, 0x0C,
            0x74, 0x34, 0x54, 0x14, 0x64, 0x24, 0x44, 0x04, 0x78, 0x38, 0x58, 0x18, 0x68, 0x28, 0x48, 0x08,
            0xF0, 0x70, 0xB0, 0x30, 0xD0, 0x50, 0x90, 0x10, 0xE0, 0x60, 0xA0, 0x20, 0xC0, 0x40, 0x80, 0x00
        };

#pragma warning disable S3963 // "static" fields should be initialized inline (it'l complain even if you put it in static constructor)
        static PKLibDecompress()
        {
            sPosition1 = GenerateDecodeTable(sDistBits, sDistCode);
            sPosition2 = GenerateDecodeTable(sLenBits, sLenCode);
        }
#pragma warning restore S3963 // "static" fields should be initialized inline

        public PKLibDecompress(Stream input)
        {
            _bitstream = new BitStream(input);

            _compressionType = (CompressionType)input.ReadByte();
            if (_compressionType != CompressionType.Binary && _compressionType != CompressionType.Ascii)
                throw new InvalidDataException("Invalid compression type: " + _compressionType);

            _dictSizeBits = input.ReadByte();
            // This is 6 in test cases
            if (4 > _dictSizeBits || _dictSizeBits > 6)
                throw new InvalidDataException("Invalid dictionary size: " + _dictSizeBits);
        }

        public byte[] Explode(int expectedSize)
        {
            byte[] outputbuffer = new byte[expectedSize];
            Stream outputstream = new MemoryStream(outputbuffer);

            int instruction;
            while ((instruction = DecodeLit()) != -1)
            {
                if (instruction < 0x100)
                {
                    outputstream.WriteByte((byte)instruction);
                }
                else
                {
                    // If instruction is greater than 0x100, it means "Repeat n - 0xFE bytes"
                    int copylength = instruction - 0xFE;
                    int moveback = DecodeDist(copylength);
                    if (moveback == 0) break;

                    int source = (int)outputstream.Position - moveback;
                    // We can't just outputstream.Write the section of the array
                    // because it might overlap with what is currently being written
                    while (copylength-- > 0)
                        outputstream.WriteByte(outputbuffer[source++]);
                }
            }

            if (outputstream.Position == expectedSize)
            {
                return outputbuffer;
            }
            else
            {
                // Resize the array
                byte[] result = new byte[outputstream.Position];
                Array.Copy(outputbuffer, 0, result, 0, result.Length);
                return result;
            }
        }

        // Return values:
        // 0x000 - 0x0FF : One byte from compressed file.
        // 0x100 - 0x305 : Copy previous block (0x100 = 1 byte)
        // -1            : EOF
        private int DecodeLit()
        {
            switch (_bitstream.ReadBits(1))
            {
                case -1:
                    return -1;

                case 1:
                    // The next bits are position in buffers
                    int pos = sPosition2[_bitstream.PeekByte()];

                    // Skip the bits we just used
                    if (_bitstream.ReadBits(sLenBits[pos]) == -1) return -1;

                    int nbits = sExLenBits[pos];
                    if (nbits != 0)
                    {
                        // TODO: Verify this conversion
                        int val2 = _bitstream.ReadBits(nbits);
                        if (val2 == -1 && (pos + val2 != 0x10e)) return -1;

                        pos = sLenBase[pos] + val2;
                    }
                    return pos + 0x100; // Return number of bytes to repeat

                case 0:
                    if (_compressionType == CompressionType.Binary)
                        return _bitstream.ReadBits(8);

                    // TODO: Text mode
                    throw new NotImplementedException("Text mode is not yet implemented");
                default:
                    return 0;
            }
        }

        private int DecodeDist(int length)
        {
            if (!_bitstream.EnsureBits(8)) return 0;
            int pos = sPosition1[_bitstream.PeekByte()];
            byte skip = sDistBits[pos];     // Number of bits to skip

            // Skip the appropriate number of bits
            if (_bitstream.ReadBits(skip) == -1) return 0;

            if (length == 2)
            {
                if (!_bitstream.EnsureBits(2)) return 0;
                pos = (pos << 2) | _bitstream.ReadBits(2);
            }
            else
            {
                if (!_bitstream.EnsureBits(_dictSizeBits)) return 0;
                pos = (pos << _dictSizeBits) | _bitstream.ReadBits(_dictSizeBits);
            }

            return pos + 1;
        }

        private static byte[] GenerateDecodeTable(byte[] bits, byte[] codes)
        {
            byte[] result = new byte[256];

            for (int i = bits.Length - 1; i >= 0; i--)
            {
                UInt32 idx1 = codes[i];
                UInt32 idx2 = (UInt32)1 << bits[i];

                do
                {
                    result[idx1] = (byte)i;
                    idx1 += idx2;
                } while (idx1 < 0x100);
            }
            return result;
        }
    }
}