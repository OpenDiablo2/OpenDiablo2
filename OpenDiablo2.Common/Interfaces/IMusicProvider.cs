using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMusicProvider : IDisposable
    {
        void LoadSong(Stream data);
        void PlaySong();
        void StopSong();
    }
}
