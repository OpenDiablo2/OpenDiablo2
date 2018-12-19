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
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2CharacterRenderer : ICharacterRenderer
    {
        sealed class DirectionCacheItem
        {
            public int Direction { get; set; }
            public eMobMode MobMode { get; set; }
            public string EquipmentHash { get; set; }

            public SDL.SDL_Rect[] SpriteRect { get; set; }
            public IntPtr[] SpriteTexture { get; set; }
            public int FramesToAnimate { get; set; }
            public int AnimationSpeed { get; set; }
        }

        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        public Guid UID { get; set; }
        public string ShieldCode { get; set; }
        public string WeaponCode { get; set; }

        private PlayerInfo currentPlayer = null;
        private PlayerInfo GetThisMob()
        {
            if (currentPlayer == null)
                currentPlayer = gameState.PlayerInfos.First(x => x.UID == UID);
            return currentPlayer;
        }


        private eMobMode MobMode
        {
            get => GetThisMob().MobMode;
            set => GetThisMob().MobMode = value;
        }
        private IMobLocation MobLocation => GetThisMob();
        private eHero Hero => GetThisMob().Hero;
        private PlayerEquipment Equipment => GetThisMob().Equipment;



        private readonly IntPtr renderer;

        private readonly List<DirectionCacheItem> directionCache = new List<DirectionCacheItem>();
        DirectionCacheItem currentDirectionCache;
        private float seconds;
        private int renderFrameIndex = 0;

        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;
        private readonly IGameState gameState;
        private int lastDirection = -1;
        private eMovementType lastMovementType = eMovementType.Stopped;
        private string lastEquipmentHash = string.Empty;

        private MPQCOF animationData;

        static readonly byte[] directionConversion = new byte[] { 3, 15, 4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14 };

        public SDL2CharacterRenderer(IntPtr renderer, IResourceManager resourceManager, IPaletteProvider paletteProvider, IGameState gameState)
        {
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.renderer = renderer;
            this.gameState = gameState;
        }

        public void Render(int pixelOffsetX, int pixelOffsetY)
        {
            if (currentDirectionCache == null)
                return;

            var destRect = new SDL.SDL_Rect
            {
                x = pixelOffsetX + currentDirectionCache.SpriteRect[renderFrameIndex].x,
                y = pixelOffsetY + currentDirectionCache.SpriteRect[renderFrameIndex].y,
                w = currentDirectionCache.SpriteRect[renderFrameIndex].w,
                h = currentDirectionCache.SpriteRect[renderFrameIndex].h
            };

            SDL.SDL_RenderCopy(renderer, currentDirectionCache.SpriteTexture[renderFrameIndex], IntPtr.Zero, ref destRect);
        }

        public void Update(long ms)
        {
            if (currentDirectionCache == null)
                return;


            if ((lastDirection != MobLocation.MovementDirection) || (lastMovementType != MobLocation.MovementType) || (lastEquipmentHash != Equipment.HashKey))
            {
                lastMovementType = MobLocation.MovementType;
                lastDirection = MobLocation.MovementDirection;
                lastEquipmentHash = Equipment.HashKey;
                ResetAnimationData();
            }

            seconds += ms / 1000f;
            var animationSeg = 15f / currentDirectionCache.AnimationSpeed;
            while (seconds >= animationSeg)
            {
                seconds -= animationSeg;
                renderFrameIndex++;
                while (renderFrameIndex >= currentDirectionCache.FramesToAnimate)
                    renderFrameIndex -= currentDirectionCache.FramesToAnimate;
            }
        }

        public void Dispose()
        {

        }

        public void ResetCache()
        {
            directionCache.Clear();
        }

        public void ResetAnimationData()
        {

            var lastMobMode = MobMode;
            switch (MobLocation.MovementType)
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
            if (lastMobMode != MobMode)
                renderFrameIndex = 0;

            currentDirectionCache = directionCache.FirstOrDefault(x => x.MobMode == MobMode && x.Direction == directionConversion[MobLocation.MovementDirection] && x.EquipmentHash == Equipment.HashKey);

            if (currentDirectionCache != null)
                return;

            animationData = resourceManager.GetPlayerAnimation(Hero, MobMode, Equipment);
            if (animationData == null)
                throw new OpenDiablo2Exception("Could not locate animation for the character!");

            var palette = paletteProvider.PaletteTable["Units"];
            CacheFrames(animationData.Layers.Select(layer => resourceManager.GetPlayerDCC(layer, Equipment, palette)).ToArray());
        }

        private unsafe void CacheFrames(MPQDCC[] layerData)
        {
            var directionCache = new DirectionCacheItem
            {
                MobMode = MobMode,
                Direction = directionConversion[MobLocation.MovementDirection]
            };
          
            var palette = paletteProvider.PaletteTable[Palettes.Units];

            var dirAnimation = animationData.Animations[0];
            directionCache.FramesToAnimate = dirAnimation.FramesPerDirection;
            directionCache.AnimationSpeed = dirAnimation.AnimationSpeed;

            var minX = Int32.MaxValue;
            var minY = Int32.MaxValue;
            var maxX = Int32.MinValue;
            var maxY = Int32.MinValue;

            var layersIgnored = 0;
            foreach (var layer in layerData)
            {
                if (layer == null)
                {
                    layersIgnored++;
                    continue;
                }

                minX = Math.Min(minX, layer.Directions[directionConversion[MobLocation.MovementDirection]].Box.Left);
                minY = Math.Min(minY, layer.Directions[directionConversion[MobLocation.MovementDirection]].Box.Top);
                maxX = Math.Max(maxX, layer.Directions[directionConversion[MobLocation.MovementDirection]].Box.Right);
                maxY = Math.Max(maxY, layer.Directions[directionConversion[MobLocation.MovementDirection]].Box.Bottom);
            }

            if (layersIgnored > 0)
                log.Warn($"{layersIgnored} animation layer(s) were not found!");

            var frameW = (maxX - minX);
            var frameH = (maxY - minY);

            directionCache.SpriteTexture = new IntPtr[directionCache.FramesToAnimate];
            directionCache.SpriteRect = new SDL.SDL_Rect[directionCache.FramesToAnimate];

            for (var frameIndex = 0; frameIndex < directionCache.FramesToAnimate; frameIndex++)
            {
                var texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, frameW, frameH);

                SDL.SDL_LockTexture(texture, IntPtr.Zero, out var pixels, out var pitch);
                var data = (UInt32*)pixels;

                var priorityBase = (directionConversion[MobLocation.MovementDirection] * animationData.FramesPerDirection * animationData.NumberOfLayers)
                        + (frameIndex * animationData.NumberOfLayers);
                for (var i = 0; i < animationData.NumberOfLayers; i++)
                {
                    var comp = animationData.Priority[priorityBase + i];
                    if (!animationData.CompositLayers.ContainsKey(comp))
                        continue;

                    var layer = layerData[animationData.CompositLayers[comp]];

                    if (layer == null)
                        continue; // TODO: This is most likely not ok

                    var direction = layer.Directions[directionConversion[MobLocation.MovementDirection]];
                    var frame = direction.Frames[frameIndex];


                    for (var y = 0; y < direction.Box.Height; y++)
                    {
                        for (var x = 0; x < direction.Box.Width; x++)
                        {
                            var paletteIndex = frame.PixelData[x + (y * direction.Box.Width)];

                            if (paletteIndex == 0)
                                continue;

                            var color = palette.Colors[paletteIndex];
                            var actualX = x + direction.Box.X - minX;
                            var actualY = y + direction.Box.Y - minY;

                            data[actualX + (actualY * (pitch / 4))] = color;

                        }
                    }

                }

                SDL.SDL_UnlockTexture(texture);
                SDL.SDL_SetTextureBlendMode(texture, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

                directionCache.SpriteTexture[frameIndex] = texture;
                directionCache.SpriteRect[frameIndex] = new SDL.SDL_Rect { x = minX, y = minY, w = frameW, h = frameH };
                directionCache.EquipmentHash = Equipment.HashKey;
                this.directionCache.Add(directionCache);
            }

            currentDirectionCache = directionCache;
        }

    }
}
