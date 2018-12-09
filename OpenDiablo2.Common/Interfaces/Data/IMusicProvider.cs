using System;
using System.IO;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ISoundProvider : IDisposable
    {
        void LoadSong(Stream data);
        void PlaySong();
        void StopSong();
        void StopSfx(int channel);
        int PlaySfx(byte[] data);
    }
}
