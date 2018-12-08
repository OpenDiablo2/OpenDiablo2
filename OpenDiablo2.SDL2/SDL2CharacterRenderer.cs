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
using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2CharacterRenderer : ICharacterRenderer
    {
        sealed class DirectionCacheItem
        {
            public int Direction { get; set; }
            public eMobMode MobMode { get; set; }

            public SDL.SDL_Rect[] SpriteRect { get; set; }
            public IntPtr[] SpriteTexture { get; set; }
            public int FramesToAnimate { get; set; }
            public int AnimationSpeed { get; set; }
            public int RenderFrameIndex { get; set; }
        }

        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        public Guid UID { get; set; }
        public PlayerLocationDetails LocationDetails { get; set; }
        public eHero Hero { get; set; }
        public eWeaponClass WeaponClass { get; set; }
        public eArmorType ArmorType { get; set; }
        public eMobMode MobMode { get; set; }

        private readonly IntPtr renderer;

        private readonly List<DirectionCacheItem> directionCache = new List<DirectionCacheItem>();
        DirectionCacheItem currentDirectionCache;
        private float seconds;

        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;

        private MPQCOF animationData;
        private MPQDCC[] layerData;

        static readonly byte[] directionConversion = new byte[] { 3, 15, 4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14 };

        public SDL2CharacterRenderer(IntPtr renderer, IResourceManager resourceManager, IPaletteProvider paletteProvider)
        {
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.renderer = renderer;
        }

        public void Render(int pixelOffsetX, int pixelOffsetY)
        {
            if (currentDirectionCache == null)
                return;

            var destRect = new SDL.SDL_Rect
            {
                x = pixelOffsetX + currentDirectionCache.SpriteRect[currentDirectionCache.RenderFrameIndex].x,
                y = pixelOffsetY + currentDirectionCache.SpriteRect[currentDirectionCache.RenderFrameIndex].y,
                w = currentDirectionCache.SpriteRect[currentDirectionCache.RenderFrameIndex].w,
                h = currentDirectionCache.SpriteRect[currentDirectionCache.RenderFrameIndex].h
            };

            SDL.SDL_RenderCopy(renderer, currentDirectionCache.SpriteTexture[currentDirectionCache.RenderFrameIndex], IntPtr.Zero, ref destRect);
        }

        public void Update(long ms)
        {
            if (currentDirectionCache == null)
                return;

            seconds += ((float)ms / 1000f);
            var animationSeg = (15f / (float)currentDirectionCache.AnimationSpeed);
            while (seconds >= animationSeg)
            {
                seconds -= animationSeg;
                currentDirectionCache.RenderFrameIndex++;
                if (currentDirectionCache.RenderFrameIndex >= currentDirectionCache.FramesToAnimate)
                    currentDirectionCache.RenderFrameIndex = 0;
            }
        }

        public void Dispose()
        {

        }

        public void ResetAnimationData()
        {
            switch (LocationDetails.MovementType)
            {
                case eMovementType.Stopped:
                    MobMode = eMobMode.PlayerTownNeutral;
                    break;
                case eMovementType.Walking:
                    MobMode = eMobMode.PlayerTownWalk;
                    break;
                case eMovementType.Running:
                    MobMode = eMobMode.PlayerRun;
                    break;
                default:
                    MobMode = eMobMode.PlayerNeutral;
                    break;
            }

            currentDirectionCache = directionCache.FirstOrDefault(x => x.MobMode == MobMode && x.Direction == directionConversion[LocationDetails.MovementDirection]);
            if (currentDirectionCache != null)
            {
                currentDirectionCache.RenderFrameIndex = 0;
                seconds = 0f;
                return;
            }

            animationData = resourceManager.GetPlayerAnimation(Hero, WeaponClass, MobMode);
            if (animationData == null)
                throw new ApplicationException("Could not locate animation for the character!");

            var palette = paletteProvider.PaletteTable["Units"];
            var data = animationData.Layers
                .Select(layer => resourceManager.GetPlayerDCC(layer, ArmorType, palette))
                .ToArray();

            log.Warn($"{data.Where(x => x == null).Count()} animation layers were not found!");

            layerData = data.Where(x => x != null)
                .ToArray();

            CacheFrames();
        }

        private unsafe void CacheFrames()
        {
            var cache = new DirectionCacheItem
            {
                MobMode = MobMode,
                Direction = directionConversion[LocationDetails.MovementDirection]
            };

            var palette = paletteProvider.PaletteTable[Palettes.Units];

            var dirAnimation = animationData.Animations[0];
            cache.FramesToAnimate = dirAnimation.FramesPerDirection;
            cache.AnimationSpeed = dirAnimation.AnimationSpeed;
            cache.RenderFrameIndex = 0;

            var minX = Int32.MaxValue;
            var minY = Int32.MaxValue;
            var maxX = Int32.MinValue;
            var maxY = Int32.MinValue;

            foreach (var layer in layerData)
            {
                minX = Math.Min(minX, layer.Directions[directionConversion[LocationDetails.MovementDirection]].Box.Left);
                minY = Math.Min(minY, layer.Directions[directionConversion[LocationDetails.MovementDirection]].Box.Top);
                maxX = Math.Max(maxX, layer.Directions[directionConversion[LocationDetails.MovementDirection]].Box.Right);
                maxY = Math.Max(maxY, layer.Directions[directionConversion[LocationDetails.MovementDirection]].Box.Bottom);
            }

            var frameW = (maxX - minX) * 2; // Hack
            var frameH = (maxY - minY) * 2;

            cache.SpriteTexture = new IntPtr[cache.FramesToAnimate];
            cache.SpriteRect = new SDL.SDL_Rect[cache.FramesToAnimate];

            for (var frameIndex = 0; frameIndex < cache.FramesToAnimate; frameIndex++)
            {
                var texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, frameW, frameH);

                SDL.SDL_LockTexture(texture, IntPtr.Zero, out IntPtr pixels, out int pitch);
                UInt32* data = (UInt32*)pixels;

                foreach (var layer in layerData)
                {
                    var direction = layer.Directions[directionConversion[LocationDetails.MovementDirection]];
                    var frame = direction.Frames[frameIndex];

                    foreach (var cell in frame.Cells)
                    {
                        if (cell.PixelData == null)
                            continue; // TODO: This isn't good

                        for (int y = 0; y < cell.Height; y++)
                        {
                            for (int x = 0; x < cell.Width; x++)
                            {
                                // Index 0 is always transparent
                                var paletteIndex = cell.PixelData[x + (y * cell.Width)];

                                if (paletteIndex == 0)
                                    continue;

                                var color = palette.Colors[paletteIndex];
                                var relativeX = (frame.XOffset - minX);
                                var relativeY = (frame.YOffset - minY);

                                var offsetX = x + cell.XOffset + (frame.Box.X - minX);
                                var offsetY = y + cell.YOffset + (frame.Box.Y - minY);
                                if (offsetX < 0 || offsetX > frameW || offsetY < 0 || offsetY > frameH)
                                    throw new ApplicationException("There is nothing we can do now.");

                                data[offsetX + (offsetY * (pitch / 4))] = color;
                            }
                        }
                    }
                }

                SDL.SDL_UnlockTexture(texture);
                SDL.SDL_SetTextureBlendMode(texture, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
                // TODO: Temporary code!
                cache.SpriteTexture[frameIndex] = texture;
                cache.SpriteRect[frameIndex] = new SDL.SDL_Rect { x = minX, y = minY, w = frameW, h = frameH };

                directionCache.Add(cache);
                currentDirectionCache = cache;
            }
        }

    }
}
