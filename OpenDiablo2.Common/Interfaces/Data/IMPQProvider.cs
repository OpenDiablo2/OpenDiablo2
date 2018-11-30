using System;
using System.Collections.Generic;
using System.IO;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMPQProvider
    {
        IEnumerable<MPQ> GetMPQs();
        IEnumerable<String> GetTextFile(string fileName);
        Stream GetStream(string fileName);
    }
}
