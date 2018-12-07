using System;
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2CharacterRenderer : ICharacterRenderer
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        public Guid UID { get; set; }
        public PlayerLocationDetails LocationDetails { get; set; }
        public eHero Hero { get; set; }
        public eWeaponClass WeaponClass { get; set; }
        public eArmorType ArmorType { get; set; }
        public eMobMode MobMode { get; set; }

        private IntPtr renderer;
        private IntPtr tempTexture;

        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;

        private MPQCOF animationData;
        private MPQDCC[] layerData;

        public SDL2CharacterRenderer(IntPtr renderer, IResourceManager resourceManager, IPaletteProvider paletteProvider)
        {
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.renderer = renderer;
        }

        public void Render(int pixelOffsetX, int pixelOffsetY)
        {

        }

        public void Update(long ms)
        {

        }

        public void Dispose()
        {

        }

        public void ResetAnimationData()
        {

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
            var dirIndex = 0; // TODO: Specify the real direction
            var frameIndex = 0;

            var dirAnimation = animationData.Animations[dirIndex];
            var framesToAnimate = dirAnimation.FramesPerDirection;

            foreach (var layer in layerData)
            {
                var direction = layer.Directions[dirIndex];
                var frame = direction.Frames[0];
                var texture = SDL.SDL_CreateTexture(
                    renderer,
                    SDL.SDL_PIXELFORMAT_ARGB8888,
                    (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING,
                    frame.Width,
                    frame.Height
                );


                IntPtr pixels;
                int pitch;

                SDL.SDL_LockTexture(texture, IntPtr.Zero, out pixels, out pitch);
                
                SDL.SDL_UnlockTexture(texture);


                tempTexture = texture;
            }
        }
    }
}
