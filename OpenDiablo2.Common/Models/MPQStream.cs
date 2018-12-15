using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using OpenDiablo2.Common.Exceptions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static OpenDiablo2.Common.Models.MPQ;


namespace OpenDiablo2.Common.Models
{
    public sealed class MPQStream : Stream
    {
        private readonly MPQ mpq;
        private readonly BlockRecord blockRecord;
        private uint[] blockPositions;
        private long position;
        private byte[] _currentData;
        private int _currentBlockIndex = -1;
        private readonly int blockSize;

        internal MPQStream(MPQ mpq, BlockRecord blockRecord)
        {
            this.mpq = mpq;
            this.blockRecord = blockRecord;
            this.blockSize = 0x200 << mpq.Header.BlockSize;

            if ((blockRecord.IsCompressed || blockRecord.IsImploded) && !blockRecord.SingleUnit)
                LoadBlockOffsets();

        }

        private void LoadBlockOffsets()
        {

            int blockposcount = (int)((blockRecord.FileSize + blockSize - 1) / blockSize) + 1;
            blockPositions = new uint[blockposcount];

            lock (mpq.fileStream)
            {
                mpq.fileStream.Seek(blockRecord.BlockOffset, SeekOrigin.Begin);
                var br = new BinaryReader(mpq.fileStream);
                for (int i = 0; i < blockposcount; i++)
                    blockPositions[i] = br.ReadUInt32();
                
            }

            uint blockpossize = (uint)blockposcount * 4;
            
            if (blockRecord.IsEncrypted)
            {
                MPQ.DecryptBlock(blockPositions, blockRecord.EncryptionSeed - 1);

                if (blockPositions[0] != blockpossize)
                    throw new OpenDiablo2Exception("Decryption failed");
                if (blockPositions[1] > blockSize + blockpossize)
                    throw new OpenDiablo2Exception("Decryption failed");
            }

        }

        public override bool CanRead => true;

        public override bool CanSeek => true;

        public override bool CanWrite => false;

        public override long Length => blockRecord.FileSize;

        public override long Position
        {
            get => position;
            set => Seek(value, SeekOrigin.Begin);
        }

        public override void Flush() { }

        public override int Read(byte[] buffer, int offset, int count)
        {
            if (blockRecord.SingleUnit)
                return ReadInternalSingleUnit(buffer, offset, count);

            int toread = count;
            int readtotal = 0;

            while (toread > 0)
            {
                int read = ReadInternal(buffer, offset, toread);
                if (read == 0) break;
                readtotal += read;
                offset += read;
                toread -= read;
            }
            return readtotal;
        }

        private int ReadInternalSingleUnit(byte[] buffer, int offset, int count)
        {
            if (position >= Length)
                return 0;

            if (_currentData == null)
                LoadSingleUnit();

            int bytestocopy = Math.Min((int)(_currentData.Length - position), count);

            Array.Copy(_currentData, position, buffer, offset, bytestocopy);

            position += bytestocopy;
            return bytestocopy;
        }

        private void LoadSingleUnit()
        {
            // Read the entire file into memory
            byte[] filedata = new byte[blockSize];
            lock (mpq.fileStream)
            {
                mpq.fileStream.Seek(mpq.Header.HeaderSize + blockRecord.BlockOffset, SeekOrigin.Begin);
                int read = mpq.fileStream.Read(filedata, 0, filedata.Length);
                if (read != filedata.Length)
                    throw new OpenDiablo2Exception("Insufficient data or invalid data length");
            }

            if (blockSize == blockRecord.FileSize)
                _currentData = filedata;
            else
                _currentData = DecompressMulti(filedata, (int)blockRecord.FileSize);
        }

        private int ReadInternal(byte[] buffer, int offset, int count)
        {
            // OW: avoid reading past the contents of the file
            if (position >= Length)
                return 0;

            BufferData();

            int localposition = (int)(position % blockSize);
            int bytestocopy = Math.Min(_currentData.Length - localposition, count);
            if (bytestocopy <= 0) return 0;

            Array.Copy(_currentData, localposition, buffer, offset, bytestocopy);

            position += bytestocopy;
            return bytestocopy;
        }

        public override int ReadByte()
        {
            if (position >= Length) return -1;

            if (blockRecord.SingleUnit)
                return ReadByteSingleUnit();

            BufferData();

            int localposition = (int)(position % blockSize);
            position++;
            return _currentData[localposition];
        }

        private int ReadByteSingleUnit()
        {
            if (_currentData == null)
                LoadSingleUnit();

            return _currentData[position++];
        }

        private void BufferData()
        {
            int requiredblock = (int)(position / blockSize);
            if (requiredblock != _currentBlockIndex)
            {
                int expectedlength = (int)Math.Min(Length - (requiredblock * blockSize), blockSize);
                _currentData = LoadBlock(requiredblock, expectedlength);
                _currentBlockIndex = requiredblock;
            }
        }

        private byte[] LoadBlock(int blockIndex, int expectedLength)
        {
            uint offset;
            int toread;
            uint encryptionseed;

            if (blockRecord.IsCompressed || blockRecord.IsImploded)
            {
                offset = blockPositions[blockIndex];
                toread = (int)(blockPositions[blockIndex + 1] - offset);
            }
            else
            {
                offset = (uint)(blockIndex * blockSize);
                toread = expectedLength;
            }
            offset += blockRecord.BlockOffset;

            byte[] data = new byte[toread];
            lock (mpq.fileStream)
            {
                mpq.fileStream.Seek(offset, SeekOrigin.Begin);
                int read = mpq.fileStream.Read(data, 0, toread);
                if (read != toread)
                    throw new OpenDiablo2Exception("Insufficient data or invalid data length");
            }

            if (blockRecord.IsEncrypted && blockRecord.FileSize > 3)
            {
                if (blockRecord.EncryptionSeed == 0)
                    throw new OpenDiablo2Exception("Unable to determine encryption key");

                encryptionseed = (uint)(blockIndex + blockRecord.EncryptionSeed);
                MPQ.DecryptBlock(data, encryptionseed);
            }

            if (blockRecord.IsCompressed && (toread != expectedLength))
            {
                //if ((blockRecord.Flags & MpqFileFlags.CompressedMulti) != 0)
                if (!blockRecord.SingleUnit)
                    data = DecompressMulti(data, expectedLength);
                else
                    data = PKDecompress(new MemoryStream(data), expectedLength);
            }

            if (blockRecord.IsImploded && (toread != expectedLength))
            {
                data = PKDecompress(new MemoryStream(data), expectedLength);
            }

            return data;
        }

        private static byte[] DecompressMulti(byte[] input, int outputLength)
        {
            Stream sinput = new MemoryStream(input);

            byte comptype = (byte)sinput.ReadByte();

            switch (comptype)
            {
                case 1: // Huffman
                    return MpqHuffman.Decompress(sinput).ToArray();
                case 2: // ZLib/Deflate
                    return ZlibDecompress(sinput, outputLength);
                case 8: // PKLib/Impode
                    return PKDecompress(sinput, outputLength);
                case 0x10: // BZip2
                    return BZip2Decompress(sinput, outputLength);
                case 0x80: // IMA ADPCM Stereo
                    return MpqWavCompression.Decompress(sinput, 2);
                case 0x40: // IMA ADPCM Mono
                    return MpqWavCompression.Decompress(sinput, 1);

                case 0x12:
                    throw new OpenDiablo2Exception("LZMA compression is not yet supported");
                // Combos
                case 0x22:
                    // TODO: sparse then zlib
                    throw new OpenDiablo2Exception("Sparse compression + Deflate compression is not yet supported");
                case 0x30:
                    // TODO: sparse then bzip2
                    throw new OpenDiablo2Exception("Sparse compression + BZip2 compression is not yet supported");
                case 0x41:
                    sinput = MpqHuffman.Decompress(sinput);
                    return MpqWavCompression.Decompress(sinput, 1);
                case 0x48:
                    {
                        byte[] result = PKDecompress(sinput, outputLength);
                        return MpqWavCompression.Decompress(new MemoryStream(result), 1);
                    }
                case 0x81:
                    sinput = MpqHuffman.Decompress(sinput);
                    return MpqWavCompression.Decompress(sinput, 2);
                case 0x88:
                    {
                        byte[] result = PKDecompress(sinput, outputLength);
                        return MpqWavCompression.Decompress(new MemoryStream(result), 2);
                    }
                default:
                    throw new OpenDiablo2Exception("Compression is not yet supported: 0x" + comptype.ToString("X"));
            }
        }

        private static byte[] BZip2Decompress(Stream data, int expectedLength)
        {
            using (var output = new MemoryStream(expectedLength))
            {
                new Ionic.BZip2.BZip2InputStream(data)
                    .CopyTo(output);
                return output.ToArray();
            }
        }

        private static byte[] PKDecompress(Stream data, int expectedLength)
        {
            PKLibDecompress pk = new PKLibDecompress(data);
            return pk.Explode(expectedLength);
        }

        private static byte[] ZlibDecompress(Stream data, int expectedLength)
        {
            // This assumes that Zlib won't be used in combination with another compression type
            byte[] Output = new byte[expectedLength];
            Stream s = new InflaterInputStream(data);
            int Offset = 0;
            while (expectedLength > 0)
            {
                int size = s.Read(Output, Offset, expectedLength);
                if (size == 0) break;
                Offset += size;
                expectedLength -= size;
            }
            return Output;
        }

        public override long Seek(long offset, SeekOrigin origin)
        {
            long target;

            switch (origin)
            {
                case SeekOrigin.Begin:
                    target = offset;
                    break;
                case SeekOrigin.Current:
                    target = Position + offset;
                    break;
                case SeekOrigin.End:
                    target = Length + offset;
                    break;
                default:
                    throw new ArgumentException("Invalid SeekOrigin", "origin");
            }

            if (target < 0)
                throw new ArgumentOutOfRangeException("offset", "Attmpted to Seek before the beginning of the stream");
            if (target >= Length)
                throw new ArgumentOutOfRangeException("offset", "Attmpted to Seek beyond the end of the stream");

            position = target;

            return position;
        }

        internal static uint DetectFileSeed(uint value0, uint value1, uint decrypted)
        {
            uint temp = (value0 ^ decrypted) - 0xeeeeeeee;

            for (int i = 0; i < 0x100; i++)
            {
                uint seed1 = temp - MPQ.cryptTable[0x400 + i];
                uint seed2 = 0xeeeeeeee + MPQ.cryptTable[0x400 + (seed1 & 0xff)];
                uint result = value0 ^ (seed1 + seed2);

                if (result != decrypted)
                    continue;

                uint saveseed1 = seed1;

                // Test this result against the 2nd value
                seed1 = ((~seed1 << 21) + 0x11111111) | (seed1 >> 11);
                seed2 = result + seed2 + (seed2 << 5) + 3;

                seed2 += MPQ.cryptTable[0x400 + (seed1 & 0xff)];
                result = value1 ^ (seed1 + seed2);

                if ((result & 0xfffc0000) == 0)
                    return saveseed1;
            }
            return 0;
        }

        public override void SetLength(long value) => throw new NotImplementedException();
        public override void Write(byte[] buffer, int offset, int count) => throw new NotImplementedException();
    }
}
