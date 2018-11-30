using System;
using System.IO;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMusicProvider : IDisposable
    {
        void LoadSong(Stream data);
        void PlaySong();
        void StopSong();
    }
}
