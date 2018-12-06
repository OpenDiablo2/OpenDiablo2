using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class ResourceManager : IResourceManager
    {
        private readonly IMPQProvider mpqProvider;
        private readonly IEngineDataManager engineDataManager;

        private Dictionary<string, ImageSet> ImageSets = new Dictionary<string, ImageSet>();
        private Dictionary<string, MPQFont> MPQFonts = new Dictionary<string, MPQFont>();
        private Dictionary<string, Palette> Palettes = new Dictionary<string, Palette>();
        private Dictionary<string, MPQDT1> DTs = new Dictionary<string, MPQDT1>();
        private Dictionary<string, MPQCOF> PlayerCOFs = new Dictionary<string, MPQCOF>();

        public Dictionary<string, List<AnimationData>> Animations { get; private set; } = new Dictionary<string, List<AnimationData>>();

        public ResourceManager(IMPQProvider mpqProvider, IEngineDataManager engineDataManager)
        {
            this.mpqProvider = mpqProvider;
            this.engineDataManager = engineDataManager;

            Animations = AnimationData.LoadFromStream(mpqProvider.GetStream(ResourcePaths.AnimationData));
        }

        public ImageSet GetImageSet(string resourcePath)
        {
            if (!ImageSets.ContainsKey(resourcePath))
                ImageSets[resourcePath] = ImageSet.LoadFromStream(mpqProvider.GetStream(resourcePath));

            return ImageSets[resourcePath];
        }

        public MPQFont GetMPQFont(string resourcePath)
        {
            if (!MPQFonts.ContainsKey(resourcePath))
                MPQFonts[resourcePath] = MPQFont.LoadFromStream(mpqProvider.GetStream($"{resourcePath}.DC6"), mpqProvider.GetStream($"{resourcePath}.tbl"));

            return MPQFonts[resourcePath];
        }

        public MPQDS1 GetMPQDS1(string resourcePath, LevelPreset level, LevelDetail levelDetail, LevelType levelType)
        {
            var mapName = resourcePath.Replace("data\\global\\tiles\\", "").Replace("\\", "/");
            return new MPQDS1(mpqProvider.GetStream(resourcePath), level, levelDetail, levelType, engineDataManager, this)
            {
                MapFile = resourcePath
            };
        }

        public Palette GetPalette(string paletteFile)
        {
            if (!Palettes.ContainsKey(paletteFile))
            {
                var paletteNameParts = paletteFile.Split('\\');
                var paletteName = paletteNameParts[paletteNameParts.Count() - 2];
                Palettes[paletteFile] = Palette.LoadFromStream(mpqProvider.GetStream(paletteFile), paletteName);
            }

            return Palettes[paletteFile];


        }

        public MPQDT1 GetMPQDT1(string resourcePath)
        {
            if (!DTs.ContainsKey(resourcePath))
                DTs[resourcePath] = new MPQDT1(mpqProvider.GetStream(resourcePath));

            return DTs[resourcePath];
        }

        public MPQCOF GetPlayerAnimation(eHero hero, eWeaponClass weaponClass, eMobMode mobMode)
        {
            var key = $"{hero.ToToken()}{mobMode.ToToken()}{weaponClass.ToToken()}";
            if (PlayerCOFs.ContainsKey(key))
                return PlayerCOFs[key];

            var path = $"{ResourcePaths.PlayerAnimationBase}\\{hero.ToToken()}\\COF\\{hero.ToToken()}{mobMode.ToToken()}{weaponClass.ToToken()}.cof";
            var result = MPQCOF.Load(mpqProvider.GetStream(path), Animations, hero, weaponClass, mobMode);
            PlayerCOFs[key] = result;

            return result;
        }

        public MPQDCC GetPlayerDCC(MPQCOF.COFLayer cofLayer, eArmorType armorType, Palette palette)
        {
            byte[] binaryData;
            using (var stream = mpqProvider.GetStream(cofLayer.GetDCCPath(armorType)))
            {
                binaryData = new byte[stream.Length];
                stream.Read(binaryData, 0, (int)stream.Length);
            }
            var result = new MPQDCC(binaryData, palette);
            return result;
        }
    }
}
