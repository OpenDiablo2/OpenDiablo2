using OpenDiablo2.Common.Interfaces;
using SDL2;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2MusicPlayer : IMusicProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);
        private IntPtr music = IntPtr.Zero;

        public SDL2MusicPlayer()
        {
            if (SDL_mixer.Mix_OpenAudio(22050, SDL_mixer.MIX_DEFAULT_FORMAT, 2, 2048) < 0)
            {
                log.Error($"SDL_mixer could not initialize! SDL_mixer Error: {SDL.SDL_GetError()}");
                return;
            }
        }

        public void PlaySong()
        {
            SDL_mixer.Mix_PlayChannel(-1, music, 1);
        }

        public void LoadSong(Stream data)
        {
            if (music != IntPtr.Zero)
                StopSong();

            var br = new BinaryReader(data);
            var bytes = br.ReadBytes((int)(data.Length - data.Position));
            music = SDL_mixer.Mix_QuickLoad_WAV(bytes);
        }

        public void StopSong()
        {
            if (music == IntPtr.Zero)
                return;
            SDL_mixer.Mix_FreeChunk(music);
        }

        public void Dispose()
        {
            StopSong();
            SDL_mixer.Mix_CloseAudio();
        }

        /*

            musicStream = data;
            SDL.SDL_AudioSpec want = new SDL.SDL_AudioSpec
            {
                freq = 22050,
                format = SDL.AUDIO_S16LSB,
                channels = 2,
                samples = 4096,
                callback = AudioCallback
            };

            SDL.SDL_OpenAudio(ref want, out audioSpec);
            SDL.SDL_PauseAudio(0);
         */

    }
}
