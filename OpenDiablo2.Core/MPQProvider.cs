﻿using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core
{
    public sealed class MPQProvider : IMPQProvider
    {
        private readonly GlobalConfiguration globalConfiguration;
        private readonly MPQ[] mpqs;
        private Dictionary<string, int> mpqLookup = new Dictionary<string, int>();

        public MPQProvider(GlobalConfiguration globalConfiguration)
        {
            this.globalConfiguration = globalConfiguration;
            this.mpqs = Directory
                .EnumerateFiles(globalConfiguration.BaseDataPath, "*.mpq")
                .Where(x => !Path.GetFileName(x).StartsWith("patch"))
                .Select(file => new MPQ(file))
                .ToArray();

            for(var i = 0; i < mpqs.Count(); i++)
            {
                foreach(var file in mpqs[i].Files)
                {
                    mpqLookup[file.ToLower()] = i;
                }
            }
        }

        public IEnumerable<MPQ> GetMPQs() => mpqs;

        public Stream GetStream(string fileName)
        { 
            return mpqs[mpqLookup[fileName.ToLower()]].OpenFile(fileName);
        }

        public IEnumerable<IEnumerable<string>> GetTextFile(string fileName)
        {
            foreach (var stream in mpqs.Where(x => x.Files.Contains(fileName)).Select(x => x.OpenFile(fileName)))
                yield return new StreamReader(stream).ReadToEnd().Split('\n').Select(x => x.Trim());
        }


    }
}
