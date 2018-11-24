using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core
{
    struct HashTableEntry
    {
        public bool IsActive { get; set; }
        public UInt16 Index { get; set; }
        public UInt32 HashValue { get; set; }
        public UInt32 IndexString { get; set; }
        public UInt32 NameString { get; set; }
        public UInt16 NameLength { get; set; }
    }

    public sealed class TextDictionary : ITextDictionary
    {
        private readonly IMPQProvider mpqProvider;

        private Dictionary<string, string> lookupTable = new Dictionary<string, string>();

        public TextDictionary(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;
            LoadDictionary();
            LoadExpansionTable();
        }

        private void LoadExpansionTable()
        {
            using (var stream = mpqProvider.GetStream(ResourcePaths.ExpansionStringTable))
            using (var br = new BinaryReader(stream))
            {
                br.ReadBytes(2); // CRC
                var numberOfElements = br.ReadUInt16();
                var hashTableSize = br.ReadUInt32();
                br.ReadByte(); // Version (always 0)
                var stringOffset = br.ReadUInt32();
                var numberOfLoopsOffset = br.ReadUInt32();
                var fileSize = br.ReadUInt32();

                var elementIndexes = new List<UInt16>();
                for (var elementIndex = 0; elementIndex < numberOfElements; elementIndex++)
                {
                    elementIndexes.Add(br.ReadUInt16());
                }

                var hashEntries = new List<HashTableEntry>();
                for (var hashEntryIndex = 0; hashEntryIndex < hashTableSize; hashEntryIndex++)
                {
                    hashEntries.Add(new HashTableEntry
                    {
                        IsActive = br.ReadByte() == 1,
                        Index = br.ReadUInt16(),
                        HashValue = br.ReadUInt32(),
                        IndexString = br.ReadUInt32(),
                        NameString = br.ReadUInt32(),
                        NameLength = br.ReadUInt16()
                    });
                }

                foreach (var hashEntry in hashEntries.Where(x => x.IsActive))
                {
                    stream.Seek(hashEntry.NameString, SeekOrigin.Begin);
                    var value = Encoding.ASCII.GetString(br.ReadBytes(hashEntry.NameLength - 1));

                    stream.Seek(hashEntry.IndexString, SeekOrigin.Begin);

                    var key = "";
                    while(true)
                    {
                        var b = br.ReadByte();
                        if (b == 0)
                            break;
                        key += (char)b;
                    }

                    lookupTable[key] = value;
                    
                }
            }
        }

        private void LoadDictionary()
        {
            var text = mpqProvider.GetTextFile(ResourcePaths.EnglishTable).First();

            var rowstoLoad = text.Where(x => x.Split(',').Count() == 3).Select(x => x.Split(',').Select(z => z.Trim()).ToArray());
            foreach (var row in rowstoLoad)
                lookupTable[row[1]] = !(row[2].StartsWith("\"") && row[2].EndsWith("\"")) ? row[2] : row[2].Substring(1, row[2].Length - 2);

        }

        public string Translate(string key) => lookupTable[key];
    }
}
