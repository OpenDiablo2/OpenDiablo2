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
 
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace OpenDiablo2.Core
{
    public sealed class MPQProvider : IMPQProvider
    {
        private readonly IList<MPQ> mpqs;
        private readonly Dictionary<string, int> mpqLookup = new Dictionary<string, int>();

        public MPQProvider(GlobalConfiguration globalConfiguration)
        {
            // TODO: Make this less dumb. We need to an external file to configure mpq load order.
            
            this.mpqs = Directory
                .EnumerateFiles(globalConfiguration.BaseDataPath, "*.mpq")
                .Where(x => !Path.GetFileName(x).StartsWith("patch"))
                .Select(file => new MPQ(file))
                .ToArray();

            

            // Load the base game files
            for(var i = 0; i < mpqs.Count(); i++)
            {
                if (Path.GetFileName(mpqs[i].Path).StartsWith("d2exp") || Path.GetFileName(mpqs[i].Path).StartsWith("d2x"))
                    continue;

                foreach(var file in mpqs[i].Files)
                {
                    mpqLookup[file.ToLower()] = i;
                }
            }

            // Load the expansion game files
            for (var i = 0; i < mpqs.Count(); i++)
            {
                if (!Path.GetFileName(mpqs[i].Path).StartsWith("d2exp") && !Path.GetFileName(mpqs[i].Path).StartsWith("d2x"))
                    continue;

                foreach (var file in mpqs[i].Files)
                {
                    mpqLookup[file.ToLower()] = i;
                }
            }
        }

        public byte[] GetBytes(string fileName)
        {
            var stream = GetStream(fileName);
            var result = new byte[stream.Length];
            stream.Read(result, 0, (int)stream.Length);
            return result;
        }

        public IEnumerator<MPQ> GetEnumerator()
        {
            return mpqs.GetEnumerator();
        }

        IEnumerator IEnumerable.GetEnumerator()
        {
            return GetEnumerator();
        }

        public Stream GetStream(string fileName)
        {
            if (!mpqLookup.ContainsKey(fileName.ToLower()))
                return null;

            return mpqs[mpqLookup[fileName.ToLower()]].OpenFile(fileName);
        }
        
        public IEnumerable<string> GetTextFile(string fileName)
            => new StreamReader(mpqs[mpqLookup[fileName.ToLower()]].OpenFile(fileName)).ReadToEnd().Split('\n');
    }
}
