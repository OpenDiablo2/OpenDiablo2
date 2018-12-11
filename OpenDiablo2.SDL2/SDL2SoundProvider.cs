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
using System.IO;
using System.Threading;
using OpenDiablo2.Common.Interfaces;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2SoundProvider : ISoundProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);
        private IntPtr music = IntPtr.Zero;
        private int musicChannel;
        private byte[] musicBytes; // Cannot be local or GC will destory it with great anger

        public SDL2SoundProvider()
        {
            if (SDL_mixer.Mix_OpenAudio(22050, SDL_mixer.MIX_DEFAULT_FORMAT, 2, 1024) < 0)
                log.Error($"SDL_mixer could not initialize! SDL_mixer Error: {SDL.SDL_GetError()}");
        }

        public void PlaySong()
        {
            if (music == IntPtr.Zero)
                return;

            musicChannel = SDL_mixer.Mix_PlayChannel(-1, music, 1);
            //SDL_mixer.Mix_Volume(musicChannel, 64); // TODO: Customizable volume
            
        }

        public void LoadSong(Stream data)
        {
            if (music != IntPtr.Zero)
                StopSong();

            musicBytes = new byte[data.Length - data.Position];
            data.ReadAsync(musicBytes, 0, (int)(data.Length - data.Position));

            // Wait until SOMETHING gets written out
            while (musicBytes[8] == 0)
                Thread.Sleep(1);

            music = SDL_mixer.Mix_QuickLoad_WAV(musicBytes);
        }

        public void StopSong()
        {
            if (music == IntPtr.Zero)
                return;
            SDL_mixer.Mix_HaltChannel(musicChannel);
            SDL_mixer.Mix_FreeChunk(music);
            music = IntPtr.Zero;
        }

        public void Dispose()
        {
            StopSong();
            SDL_mixer.Mix_CloseAudio();
        }

        public int PlaySfx(byte[] data)
        {
            if (data == null || data.Length == 0)
                return -1;

            var sound = SDL_mixer.Mix_QuickLoad_WAV(data);
            var channel = SDL_mixer.Mix_PlayChannel(-1, sound, 0);
            //SDL_mixer.Mix_Volume(channel, 64); // TODO: Customizable volume
            return channel;
        }

        public void StopSfx(int channel)
        {
            SDL_mixer.Mix_HaltChannel(channel);
        }
    }
}
