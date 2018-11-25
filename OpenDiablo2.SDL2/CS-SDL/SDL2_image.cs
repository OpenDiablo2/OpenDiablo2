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
	public static class SDL_image
	{
		#region SDL2# Variables

		/* Used by DllImport to load the native library. */
		private const string nativeLibName = "SDL2_image";

		#endregion

		#region SDL_image.h

		/* Similar to the headers, this is the version we're expecting to be
		 * running with. You will likely want to check this somewhere in your
		 * program!
		 */
		public const int SDL_IMAGE_MAJOR_VERSION =	2;
		public const int SDL_IMAGE_MINOR_VERSION =	0;
		public const int SDL_IMAGE_PATCHLEVEL =		2;

		[Flags]
		public enum IMG_InitFlags
		{
			IMG_INIT_JPG =	0x00000001,
			IMG_INIT_PNG =	0x00000002,
			IMG_INIT_TIF =	0x00000004,
			IMG_INIT_WEBP =	0x00000008
		}

		public static void SDL_IMAGE_VERSION(out SDL.SDL_version X)
		{
			X.major = SDL_IMAGE_MAJOR_VERSION;
			X.minor = SDL_IMAGE_MINOR_VERSION;
			X.patch = SDL_IMAGE_PATCHLEVEL;
		}

		[DllImport(nativeLibName, EntryPoint = "IMG_Linked_Version", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_IMG_Linked_Version();
		public static SDL.SDL_version IMG_Linked_Version()
		{
			SDL.SDL_version result;
			IntPtr result_ptr = INTERNAL_IMG_Linked_Version();
			result = (SDL.SDL_version) Marshal.PtrToStructure(
				result_ptr,
				typeof(SDL.SDL_version)
			);
			return result;
		}

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int IMG_Init(IMG_InitFlags flags);

		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern void IMG_Quit();

		/* IntPtr refers to an SDL_Surface* */
		[DllImport(nativeLibName, EntryPoint = "IMG_Load", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_IMG_Load(
			byte[] file
		);
		public static IntPtr IMG_Load(string file)
		{
			return INTERNAL_IMG_Load(SDL.UTF8_ToNative(file));
		}

		/* src refers to an SDL_RWops*, IntPtr to an SDL_Surface* */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr IMG_Load_RW(
			IntPtr src,
			int freesrc
		);

		/* src refers to an SDL_RWops*, IntPtr to an SDL_Surface* */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, EntryPoint = "IMG_LoadTyped_RW", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_IMG_LoadTyped_RW(
			IntPtr src,
			int freesrc,
			byte[] type
		);
		public static IntPtr IMG_LoadTyped_RW(
			IntPtr src,
			int freesrc,
			string type
		) {
			return INTERNAL_IMG_LoadTyped_RW(
				src,
				freesrc,
				SDL.UTF8_ToNative(type)
			);
		}

		/* IntPtr refers to an SDL_Texture*, renderer to an SDL_Renderer* */
		[DllImport(nativeLibName, EntryPoint = "IMG_LoadTexture", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_IMG_LoadTexture(
			IntPtr renderer,
			byte[] file
		);
		public static IntPtr IMG_LoadTexture(
			IntPtr renderer,
			string file
		) {
			return INTERNAL_IMG_LoadTexture(
				renderer,
				SDL.UTF8_ToNative(file)
			);
		}

		/* renderer refers to an SDL_Renderer*.
		 * src refers to an SDL_RWops*.
		 * IntPtr to an SDL_Texture*.
		 */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr IMG_LoadTexture_RW(
			IntPtr renderer,
			IntPtr src,
			int freesrc
		);

		/* renderer refers to an SDL_Renderer*.
		 * src refers to an SDL_RWops*.
		 * IntPtr to an SDL_Texture*.
		 */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, EntryPoint = "IMG_LoadTextureTyped_RW", CallingConvention = CallingConvention.Cdecl)]
		private static extern IntPtr INTERNAL_IMG_LoadTextureTyped_RW(
			IntPtr renderer,
			IntPtr src,
			int freesrc,
			byte[] type
		);
		public static IntPtr IMG_LoadTextureTyped_RW(
			IntPtr renderer,
			IntPtr src,
			int freesrc,
			string type
		) {
			return INTERNAL_IMG_LoadTextureTyped_RW(
				renderer,
				src,
				freesrc,
				SDL.UTF8_ToNative(type)
			);
		}

		/* IntPtr refers to an SDL_Surface* */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern IntPtr IMG_ReadXPMFromArray(
			[In()] [MarshalAs(UnmanagedType.LPArray, ArraySubType = UnmanagedType.LPStr)]
				string[] xpm
		);

		/* surface refers to an SDL_Surface* */
		[DllImport(nativeLibName, EntryPoint = "IMG_SavePNG", CallingConvention = CallingConvention.Cdecl)]
		private static extern int INTERNAL_IMG_SavePNG(
			IntPtr surface,
			byte[] file
		);
		public static int IMG_SavePNG(IntPtr surface, string file)
		{
			return INTERNAL_IMG_SavePNG(
				surface,
				SDL.UTF8_ToNative(file)
			);
		}

		/* surface refers to an SDL_Surface*, dst to an SDL_RWops* */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int IMG_SavePNG_RW(
			IntPtr surface,
			IntPtr dst,
			int freedst
		);

		/* surface refers to an SDL_Surface* */
		[DllImport(nativeLibName, EntryPoint = "IMG_SaveJPG", CallingConvention = CallingConvention.Cdecl)]
		private static extern int INTERNAL_IMG_SaveJPG(
			IntPtr surface,
			byte[] file,
			int quality
		);
		public static int IMG_SaveJPG(IntPtr surface, string file, int quality)
		{
			return INTERNAL_IMG_SaveJPG(
				surface,
				SDL.UTF8_ToNative(file),
				quality
			);
		}

		/* surface refers to an SDL_Surface*, dst to an SDL_RWops* */
		/* THIS IS A PUBLIC RWops FUNCTION! */
		[DllImport(nativeLibName, CallingConvention = CallingConvention.Cdecl)]
		public static extern int IMG_SaveJPG_RW(
			IntPtr surface,
			IntPtr dst,
			int freedst,
			int quality
		);

		#endregion
	}
}
