using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core
{
    public sealed class TextDictionary : ITextDictionary
    {
        private readonly IMPQProvider mpqProvider;

        private Dictionary<string, string> lookupTable = new Dictionary<string, string>();

        public TextDictionary(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;
            LoadDictionary();
        }

        private void LoadDictionary()
        {
            var text = mpqProvider.GetTextFile(ResourcePaths.EnglishTable).First();

            var rowstoLoad = text.Where(x => x.Split(',').Count() == 3).Select(x => x.Split(',').Select(z => z.Trim()).ToArray());
            foreach(var row in rowstoLoad)
                lookupTable[row[1]] = !(row[2].StartsWith("\"") && row[2].EndsWith("\"")) ? row[2] : row[2].Substring(1, row[2].Length - 2);

        }

        public string Translate(string key) => lookupTable[key];
    }
}
