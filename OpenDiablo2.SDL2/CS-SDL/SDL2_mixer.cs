#region License
/* SDL2# - C# Wrapper for SDL2
 *
 * Copyright (c) 2013-2016 Ethan Lee.
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from
 * the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not
 * claim that you wrote the original software. If you use this software in a
 * product, an acknowledgment in the product documentation would be
 * appreciated but is not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source distribution.
 *
 * Ethan "flibitijibibo" Lee <flibitijibibo@flibitijibibo.com>
 *
 */
#endregion

#region Using Statements
using System;
using System.Runtime.InteropServices;
#endregion

namespace SDL2
{
	public static class SDL_mixer
	{
		#region SDL2# Variables

		/* Used by DllImport to load the native library. */
		private const string nativeLibName = "SDL2_mixer";

		#endregion

		#region SDL_mixer.h

		/* Similar to the headers, this is the version we're expecting to be
		 * running with. You will likely want to check this somewhere in your
		 * program!
		 */
		public const int SDL_MIXER_MAJOR_VERSION =	2;
		public const int SDL_MIXER_MINOR_VERSION =	0;
		public const int SDL_MIXER_PATCHLEVEL =		2;

		/* In C, you can redefine this value before including SDL_mixer.h.
		 * We're not going to allow this in SDL2#, since the value of this
		 * variable is persistent and not dependent on preprocessor ordering.
		 */
		public const int MIX_CHANNELS = 8;

		public static readonly int MIX_DEFAULT_FREQUENCY = 22050;
		public static readonly ushort MIX_DEFAULT_FORMAT =
			BitConverter.IsLittleEndian ? SDL.AUDIO_S16LSB : SDL.AUDIO_S16MSB;
		public static readonly int MIX_DEFAULT_CHANNELS = 2;
		public static readonly byte MIX_MAX_VOLUME = 128;

		[Flags]
		public enum MIX_InitFlags
		{
			MIX_INIT_FLAC =		0x00000001,
			MIX_INIT_MOD =		0x00000002,
			MIX_INIT_MP3 =		0x00000008,
			MIX_INIT_OGG =		0x00000010,
			MIX_INIT_MID =		0x00000020,
		}

		public enum Mix_Fading
		{
			MIX_NO_FADING,
			MIX_FADING_OUT,
			MIX_FADING_IN
		}

		public enum Mix_MusicType
		{
			MUS_NONE,
			MUS_CMD,
			MUS_WAV,
			MUS_MOD,
			MUS_MID,
			MUS_OGG,
			MUS_MP3,
			MUS_MP3_MAD,
			MUS_FLAC,
			MUS_MODPLUG
		}

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate void MixFuncDelegate(
			IntPtr udata, // void*
			IntPtr stream, // Uint8*
			int len
		);

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate void Mix_EffectFunc_t(
			int chan,
			IntPtr stream, // void*
			int len,
			IntPtr udata // void*
		);

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate void Mix_EffectDone_t(
			int chan,
			IntPtr udata // void*
		);

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate void MusicFinishedDelegate();

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate void ChannelFinishedDelegate(int channel);

		[UnmanagedFunctionPointer(CallingConvention.Cdecl)]
		public delegate int SoundFontDelegate(
			IntPtr a, // const char*
			IntPtr b // void*
		);

		public static void SDL_MIXER_VERSION(out SDL.SDL_version X)
		{
			X.major = SDL_MIXER_MAJOR_VERSION;
			X.minor = SDL_MIXER_MINOR_VERSION;
			X.patch = SDL_MIXER_PATCHLEVEL;
		}

		[DllImport(nativeLibName, EntryPoint = "MIX_Linked_Version", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_MIX_Linked_Version();
		public static SDL.SDL_version MIX_Linked_Version()
		{
			SDL.SDL_version result;
			IntPtr result_ptr = INTERNAL_MIX_Linked_Version();
			result = (SDL.SDL_version) Marshal.PtrToStructure(
				result_ptr,
				typeof(SDL.SDL_version)
			);
			return result;
		}

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_Init(MIX_InitFlags flags);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_Quit();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_OpenAudio(
			int frequency,
			ushort format,
			int channels,
			int chunksize
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_AllocateChannels(int numchans);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_QuerySpec(
			out int frequency,
			out ushort format,
			out int channels
		);

		/* src refers to an SDL_RWops*, IntPtr to a Mix_Chunk* */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr Mix_LoadWAV_RW(
			IntPtr src,
			int freesrc
		);
		
		/* IntPtr refers to a Mix_Chunk* */
		/* This is an RWops macro in the C header. */
		public static IntPtr Mix_LoadWAV(string file)
		{
			IntPtr rwops = SDL.INTERNAL_SDL_RWFromFile(file, "rb");
			return Mix_LoadWAV_RW(rwops, 1);
		}

		/* IntPtr refers to a Mix_Music* */
		[DllImport(nativeLibName, EntryPoint = "Mix_LoadMUS", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_Mix_LoadMUS(
			byte[] file
		);
		public static IntPtr Mix_LoadMUS(string file)
		{
			return INTERNAL_Mix_LoadMUS(SDL.UTF8_ToNative(file));
		}

		/* IntPtr refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr Mix_QuickLoad_WAV(
			[In()] [MarshalAs(UnmanagedType.LPArray, ArraySubType = UnmanagedType.U1)]
				byte[] mem
		);

		/* IntPtr refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr Mix_QuickLoad_RAW(
			[In()] [MarshalAs(UnmanagedType.LPArray, ArraySubType = UnmanagedType.U1, SizeParamIndex = 1)]
				byte[] mem,
			uint len
		);

		/* chunk refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_FreeChunk(IntPtr chunk);

		/* music refers to a Mix_Music* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_FreeMusic(IntPtr music);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GetNumChunkDecoders();

		[DllImport(nativeLibName, EntryPoint = "Mix_GetChunkDecoder", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_Mix_GetChunkDecoder(int index);
		public static string Mix_GetChunkDecoder(int index)
		{
			return SDL.UTF8_ToManaged(
				INTERNAL_Mix_GetChunkDecoder(index)
			);
		}

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GetNumMusicDecoders();

		[DllImport(nativeLibName, EntryPoint = "Mix_GetMusicDecoder", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_Mix_GetMusicDecoder(int index);
		public static string Mix_GetMusicDecoder(int index)
		{
			return SDL.UTF8_ToManaged(
				INTERNAL_Mix_GetMusicDecoder(index)
			);
		}

		/* music refers to a Mix_Music* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern Mix_MusicType Mix_GetMusicType(IntPtr music);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_SetPostMix(
			MixFuncDelegate mix_func,
			IntPtr arg // void*
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_HookMusic(
			MixFuncDelegate mix_func,
			IntPtr arg // void*
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_HookMusicFinished(
			MusicFinishedDelegate music_finished
		);

		/* IntPtr refers to a void* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr Mix_GetMusicHookData();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_ChannelFinished(
			ChannelFinishedDelegate channel_finished
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_RegisterEffect(
			int chan,
			Mix_EffectFunc_t f,
			Mix_EffectDone_t d,
			IntPtr arg // void*
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_UnregisterEffect(
			int channel,
			Mix_EffectFunc_t f
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_UnregisterAllEffects(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetPanning(
			int channel,
			byte left,
			byte right
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetPosition(
			int channel,
			short angle,
			byte distance
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetDistance(int channel, byte distance);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetReverseStereo(int channel, int flip);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_ReserveChannels(int num);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupChannel(int which, int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupChannels(int from, int to, int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupAvailable(int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupCount(int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupOldest(int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GroupNewer(int tag);

		/* chunk refers to a Mix_Chunk* */
		public static int Mix_PlayChannel(
			int channel,
			IntPtr chunk,
			int loops
		) {
			return Mix_PlayChannelTimed(channel, chunk, loops, -1);
		}

		/* chunk refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_PlayChannelTimed(
			int channel,
			IntPtr chunk,
			int loops,
			int ticks
		);

		/* music refers to a Mix_Music* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_PlayMusic(IntPtr music, int loops);

		/* music refers to a Mix_Music* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeInMusic(
			IntPtr music,
			int loops,
			int ms
		);

		/* music refers to a Mix_Music* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeInMusicPos(
			IntPtr music,
			int loops,
			int ms,
			double position
		);

		/* chunk refers to a Mix_Chunk* */
		public static int Mix_FadeInChannel(
			int channel,
			IntPtr chunk,
			int loops,
			int ms
		) {
			return Mix_FadeInChannelTimed(channel, chunk, loops, ms, -1);
		}

		/* chunk refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeInChannelTimed(
			int channel,
			IntPtr chunk,
			int loops,
			int ms,
			int ticks
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_Volume(int channel, int volume);

		/* chunk refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_VolumeChunk(
			IntPtr chunk,
			int volume
		);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_VolumeMusic(int volume);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_HaltChannel(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_HaltGroup(int tag);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_HaltMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_ExpireChannel(int channel, int ticks);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeOutChannel(int which, int ms);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeOutGroup(int tag, int ms);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_FadeOutMusic(int ms);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern Mix_Fading Mix_FadingMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern Mix_Fading Mix_FadingChannel(int which);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_Pause(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_Resume(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_Paused(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_PauseMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_ResumeMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_RewindMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_PausedMusic();

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetMusicPosition(double position);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_Playing(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_PlayingMusic();

		[DllImport(nativeLibName, EntryPoint = "Mix_SetMusicCMD", CallingConvention = CallingConvention.Cdecl)]
		private static extern int INTERNAL_Mix_SetMusicCMD(
			byte[] command
		);
		public static int Mix_SetMusicCMD(string command)
		{
			return INTERNAL_Mix_SetMusicCMD(
				SDL.UTF8_ToNative(command)
			);
		}

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_SetSynchroValue(int value);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_GetSynchroValue();

		[DllImport(nativeLibName, EntryPoint = "Mix_SetSoundFonts", CallingConvention = CallingConvention.Cdecl)]
		private static extern int INTERNAL_Mix_SetSoundFonts(
			byte[] paths
		);
		public static int Mix_SetSoundFonts(string paths)
		{
			return INTERNAL_Mix_SetSoundFonts(
				SDL.UTF8_ToNative(paths)
			);
		}

		[DllImport(nativeLibName, EntryPoint = "Mix_GetSoundFonts", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_Mix_GetSoundFonts();
		public static string Mix_GetSoundFonts()
		{
			return SDL.UTF8_ToManaged(INTERNAL_Mix_GetSoundFonts());
		}

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int Mix_EachSoundFont(
			SoundFontDelegate function,
			IntPtr data // void*
		);

		/* IntPtr refers to a Mix_Chunk* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr Mix_GetChunk(int channel);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void Mix_CloseAudio();

		#endregion
	}
}
