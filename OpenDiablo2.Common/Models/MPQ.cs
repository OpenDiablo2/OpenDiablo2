/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQ : IDisposable
    {
        private const string HEADER_SIGNATURE = "MPQ\x1A";
        private const string LISTFILE_NAME = "(listfile)";
        private const int MPQ_HASH_FILE_KEY = 3;
        private const int MPQ_HASH_NAME_A = 1;
        private const int MPQ_HASH_NAME_B = 2;

        internal struct HeaderRecord
        {
            public UInt32 HeaderSize;
            public UInt32 ArchiveSize;
            public UInt16 FormatVersion;
            public Byte BlockSize;
            public UInt32 HashTablePos;
            public UInt32 BlockTablePos;
            public UInt32 HashTableSize;
            public UInt32 BlockTableSize;
            // Other properties are for >0 MPQ version
        }

        [Flags]
        internal enum eBlockRecordBitflag : UInt32
        {
            IsFile = 0x80000000, // Block is a file, and follows the file data format; otherwise, block is free space or unused. If the block is not a file, all other flags should be cleared, and FileSize should be 0.
            SingleUnit = 0x01000000, // File is stored as a single unit, rather than split into sectors.
            KeyAdjusted = 0x00020000, // The file's encryption key is adjusted by the block offset and file size (explained in detail in the File Data section). File must be encrypted.
            IsEncrypted = 0x00010000, // File is encrypted.
            IsCompressed = 0x00000200, // File is compressed. File cannot be imploded.
            IsImploded = 0x00000100,  // File is imploded. File cannot be compressed.
            IsPatchFile = 0x00100000, // Marks as a mpq patch file
            IsDeleteFile = 0x02000000, // Marks as a delete request
        }

        internal struct BlockRecord
        {
            public UInt32 BlockOffset;
            public UInt32 BlockSize;
            public UInt32 FileSize;
            public UInt32 Flags;
            public uint EncryptionSeed { get; set; }
            public string FileName { get; internal set; }

            public bool IsFile => (Flags & (UInt32)eBlockRecordBitflag.IsFile) != 0;
            public bool SingleUnit => (Flags & (UInt32)eBlockRecordBitflag.SingleUnit) != 0;
            public bool KeyAdjusted => (Flags & (UInt32)eBlockRecordBitflag.KeyAdjusted) != 0;
            public bool IsEncrypted => (Flags & (UInt32)eBlockRecordBitflag.IsEncrypted) != 0;
            public bool IsCompressed => (Flags & (UInt32)eBlockRecordBitflag.IsCompressed) != 0;
            public bool IsImploded => (Flags & (UInt32)eBlockRecordBitflag.IsImploded) != 0;
            public bool IsPatchFile => (Flags & (UInt32)eBlockRecordBitflag.IsPatchFile) != 0;
            public bool IsDeleteFile => (Flags & (UInt32)eBlockRecordBitflag.IsDeleteFile) != 0;
        }

        internal struct HashRecord
        {
            public UInt32 FilePathHashA;
            public UInt32 FilePathHashB;
            public UInt16 Language;
            public UInt16 Platform;
            public UInt32 FileBlockIndex;
        }

        internal static readonly UInt32[] cryptTable = new UInt32[0x500];
        internal HeaderRecord Header;
        private readonly List<BlockRecord> blockTable = new List<BlockRecord>();
        private readonly List<HashRecord> hashTable = new List<HashRecord>();
        internal Stream fileStream;

        public UInt16 LanguageId { get; internal set; } = 0;
        public const byte Platform = 0;

        public string Path { get; private set; }
        public eMPQFormatVersion FormatVersion => (eMPQFormatVersion)Header.FormatVersion;
        public List<string> Files => GetFilePaths();
        private List<string> FilesOverride = null;

        private List<string> GetFilePaths()
        {
            if(FilesOverride != null)
            {
                return FilesOverride; // if we were passed in an existing listfile,
                // don't query again; use the one we were given instead
            }
            using (var stream = OpenFile(LISTFILE_NAME))
            {
                if (stream == null)
                {
                    return new List<string>();
                }

                var sr = new StreamReader(stream);
                var text = sr.ReadToEnd();

                return text.Split('\n').Where(x => !String.IsNullOrWhiteSpace(x)).Select(x => x.Trim()).ToList();
            }
        }

        static MPQ()
        {
            InitializeCryptTable();
        }

        public MPQ(string path, List<string> listfile = null)
        {
            this.Path = path;
            this.FilesOverride = listfile;

            // If you crash here, you may have Diablo2 open... can't do that :)
            fileStream = new FileStream(path, FileMode.Open);

            using (var br = new BinaryReader(fileStream, Encoding.Default, true))
            {
                var header = Encoding.ASCII.GetString(br.ReadBytes(4));
                if (header != HEADER_SIGNATURE)
                    throw new OpenDiablo2Exception($"Unknown header signature '{header}' detected while processing '{Path}'!");

                ParseMPQHeader(br);
            }

        }

        private static void InitializeCryptTable()
        {
            UInt32 seed = 0x00100001;
            UInt32 index1 = 0;
            UInt32 index2 = 0;
            int i;

            for (index1 = 0; index1 < 0x100; index1++)
            {
                for (index2 = index1, i = 0; i < 5; i++, index2 += 0x100)
                {
                    seed = (seed * 125 + 3) % 0x2AAAAB;
                    var temp = (seed & 0xFFFF) << 0x10;

                    seed = (seed * 125 + 3) % 0x2AAAAB;

                    cryptTable[index2] = temp | (seed & 0xFFFF);
                }
            }
        }

        internal static void DecryptBlock(uint[] data, uint seed1)
        {
            uint seed2 = 0xeeeeeeee;

            for (int i = 0; i < data.Length; i++)
            {
                seed2 += cryptTable[0x400 + (seed1 & 0xff)];
                uint result = data[i];
                result ^= seed1 + seed2;

                seed1 = ((~seed1 << 21) + 0x11111111) | (seed1 >> 11);
                seed2 = result + seed2 + (seed2 << 5) + 3;
                data[i] = result;
            }
        }

        internal static void DecryptBlock(byte[] data, uint seed1)
        {
            uint seed2 = 0xeeeeeeee;

            // NB: If the block is not an even multiple of 4,
            // the remainder is not encrypted
            for (int i = 0; i < data.Length - 3; i += 4)
            {
                seed2 += cryptTable[(int)(0x400 + (seed1 & 0xff))];

                uint result = BitConverter.ToUInt32(data, i);
                result ^= seed1 + seed2;

                seed1 = ((~seed1 << 21) + 0x11111111) | (seed1 >> 11);
                seed2 = result + seed2 + (seed2 << 5) + 3;

                data[i + 0] = (byte)(result & 0xff);
                data[i + 1] = (byte)((result >> 8) & 0xff);
                data[i + 2] = (byte)((result >> 16) & 0xff);
                data[i + 3] = (byte)((result >> 24) & 0xff);
            }
        }


        private void ParseMPQHeader(BinaryReader br)
        {
            Header = new HeaderRecord
            {
                HeaderSize = br.ReadUInt32(),
                ArchiveSize = br.ReadUInt32(),
                FormatVersion = br.ReadUInt16(),
                BlockSize = (byte)br.ReadInt16(),
                HashTablePos = br.ReadUInt32(),
                BlockTablePos = br.ReadUInt32(),
                HashTableSize = br.ReadUInt32(),
                BlockTableSize = br.ReadUInt32()
            };

            if (FormatVersion != eMPQFormatVersion.Format1)
                throw new OpenDiablo2Exception($"Unsupported MPQ format version of {Header.FormatVersion} detected for '{Path}'!");

            if (br.BaseStream.Position != Header.HeaderSize)
                throw new OpenDiablo2Exception($"Invalid header size detected for '{Path}'. Expected to be at offset {Header.HeaderSize} but we are at offset {br.BaseStream.Position} instead!");

            br.BaseStream.Seek(Header.BlockTablePos, SeekOrigin.Begin);

            // Process the block table
            var bData = br.ReadBytes((int)(16 * Header.BlockTableSize));
            DecryptBlock(bData, HashString("(block table)", MPQ_HASH_FILE_KEY));
            using (var ms = new MemoryStream(bData))
            using (var dr = new BinaryReader(new MemoryStream(bData)))
                for (var index = 0; index < Header.BlockTableSize; index++)
                {
                    blockTable.Add(new BlockRecord
                    {
                        BlockOffset = dr.ReadUInt32(),
                        BlockSize = dr.ReadUInt32(),
                        FileSize = dr.ReadUInt32(),
                        Flags = dr.ReadUInt32()
                    });
                }

            // Process the hash table
            br.BaseStream.Seek(Header.HashTablePos, SeekOrigin.Begin);
            bData = br.ReadBytes((int)(16 * Header.HashTableSize));
            DecryptBlock(bData, HashString("(hash table)", MPQ_HASH_FILE_KEY));
            using (var ms = new MemoryStream(bData))
            using (var dr = new BinaryReader(new MemoryStream(bData)))
                for (var index = 0; index < Header.HashTableSize; index++)
                {
                    hashTable.Add(new HashRecord
                    {
                        FilePathHashA = dr.ReadUInt32(),
                        FilePathHashB = dr.ReadUInt32(),
                        Language = dr.ReadUInt16(),
                        Platform = dr.ReadUInt16(),
                        FileBlockIndex = dr.ReadUInt32()
                    });
                }

        }

        private uint CalculateEncryptionSeed(BlockRecord record)
        {
            if (record.FileName == null) return 0;

            uint seed = HashString(System.IO.Path.GetFileName(record.FileName), MPQ_HASH_FILE_KEY);
            if (record.KeyAdjusted)
                seed = (seed + record.BlockOffset) ^ record.FileSize;
            return seed;
        }

        private static UInt32 HashString(string inputString, UInt32 hashType)
        {
            if (hashType > MPQ_HASH_FILE_KEY)
                throw new OpenDiablo2Exception($"Unknown hash type {hashType} for input string {inputString}");

            UInt32 seed1 = 0x7FED7FED;
            UInt32 seed2 = 0xEEEEEEEE;

            foreach (var ch in inputString)
            {
                var chInt = (UInt32)char.ToUpper(ch);
                seed1 = cryptTable[(hashType * 0x100) + chInt] ^ (seed1 + seed2);
                seed2 = chInt + seed1 + seed2 + (seed2 << 5) + 3;
            }
            return seed1;
        }

        /*
        private static UInt32 ComputeFileKey(string filePath, BlockRecord blockRecord, UInt32 archiveOffset)
        {
            var fileName = filePath.Split('\\').Last();

            // Hash the name to get the base key
            var fileKey = HashString(fileName, MPQ_HASH_FILE_KEY);

            // Offset-adjust the key if necessary
            if (blockRecord.KeyAdjusted)
                fileKey = (fileKey + blockRecord.BlockOffset) ^ blockRecord.FileSize;

            return fileKey;
        }
        
        private bool FindFileInHashTable(string filePath, out UInt32 fileHashEntry)
        {
            fileHashEntry = 0;

            // Find the home entry in the hash table for the file
            UInt32 initEntry = HashString(filePath, MPQ_HASH_TABLE_OFFSET) & Header.HashTableSize - 1;

            // Is there anything there at all?
            if (hashTable[(int)initEntry].FileBlockIndex == MPQ_HASH_ENTRY_EMPTY)
                return false;

            // Compute the hashes to compare the hash table entry against
            var nNameHashA = HashString(filePath, MPQ_HASH_NAME_A);
            var nNameHashB = HashString(filePath, MPQ_HASH_NAME_B);
            var iCurEntry = initEntry;

            // Check each entry in the hash table till a termination point is reached
            do
            {
                if (hashTable[(int)iCurEntry].FileBlockIndex != MPQ_HASH_ENTRY_DELETED)
                {
                    if (hashTable[(int)iCurEntry].FilePathHashA == nNameHashA
                        && hashTable[(int)iCurEntry].FilePathHashB == nNameHashB
                        && hashTable[(int)iCurEntry].Language == LanguageId
                        && hashTable[(int)iCurEntry].Platform == (UInt16)PlatformID.Win32S)
                    {
                        fileHashEntry = iCurEntry;

                        return true;
                    }
                }

                iCurEntry = (iCurEntry + 1) & Header.HashTableSize - 1;
            } while (iCurEntry != initEntry && hashTable[(int)iCurEntry].FileBlockIndex != MPQ_HASH_ENTRY_EMPTY);

            return false;
        }
        */
        private bool GetHashRecord(string fileName, out HashRecord hash)
        {
            uint index = HashString(fileName, 0);
            index &= Header.HashTableSize - 1;
            uint name1 = HashString(fileName, MPQ_HASH_NAME_A);
            uint name2 = HashString(fileName, MPQ_HASH_NAME_B);

            for (uint i = index; i < hashTable.Count; ++i)
            {
                hash = hashTable[(int)i];
                if (hash.FilePathHashA == name1 && hash.FilePathHashB == name2)
                    return true;
            }

            for (uint i = 0; i < index; i++)
            {
                hash = hashTable[(int)i];
                if (hash.FilePathHashA == name1 && hash.FilePathHashB == name2)
                    return true;
            }

            hash = new HashRecord();
            return false;
        }

        public MPQStream OpenFile(string filename)
        {
            if (!GetHashRecord(filename, out HashRecord hash))
                throw new FileNotFoundException("File not found: " + filename);

            BlockRecord block = blockTable[(int)hash.FileBlockIndex];
            block.FileName = filename.ToLower();
            block.EncryptionSeed = CalculateEncryptionSeed(block);
            return new MPQStream(this, block);
        }

        public bool HasFile(string filename)
        {
            return GetHashRecord(filename, out HashRecord hash);
        }

        public void Dispose()
        {
            fileStream?.Dispose();
        }
    }
}
