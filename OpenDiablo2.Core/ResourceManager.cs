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

using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Core
{
    public sealed class ResourceManager : IResourceManager
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly ICache cache;
        private readonly IMPQProvider mpqProvider;
        private readonly IEngineDataManager engineDataManager;

        public Dictionary<string, List<AnimationData>> Animations { get; private set; }

        public ResourceManager(ICache cache, IMPQProvider mpqProvider, IEngineDataManager engineDataManager)
        {
            this.cache = cache;
            this.mpqProvider = mpqProvider;
            this.engineDataManager = engineDataManager;

            Animations = AnimationData.LoadFromStream(mpqProvider.GetStream(ResourcePaths.AnimationData));
        }

        public ImageSet GetImageSet(string resourcePath)
            // => cache.AddOrGetExisting($"ImageSet::{resourcePath}", () => ImageSet.LoadFromStream(mpqProvider.GetStream(resourcePath)));
            => ImageSet.LoadFromStream(mpqProvider.GetStream(resourcePath));

        public MPQFont GetMPQFont(string resourcePath)
            => cache.AddOrGetExisting($"Font::{resourcePath}", () => MPQFont.LoadFromStream(mpqProvider.GetStream($"{resourcePath}.DC6"), mpqProvider.GetStream($"{resourcePath}.tbl")));

        public MPQDS1 GetMPQDS1(string resourcePath, LevelPreset level, LevelType levelType)
            => cache.AddOrGetExisting($"DS1::{resourcePath}::{level}::{levelType}", ()
                => new MPQDS1(mpqProvider.GetStream(resourcePath), level, levelType, engineDataManager, this) { MapFile = resourcePath });

        public Palette GetPalette(string paletteFile)
            => cache.AddOrGetExisting($"Palette::{paletteFile}", () =>
            {
                var paletteNameParts = paletteFile.Split('\\');
                var paletteName = paletteNameParts[paletteNameParts.Count() - 2];
                return Palette.LoadFromStream(mpqProvider.GetStream(paletteFile), paletteName);
            });

        public MPQDT1 GetMPQDT1(string resourcePath)
            => cache.AddOrGetExisting($"DT1::{resourcePath}", () => new MPQDT1(mpqProvider.GetStream(resourcePath)));

        public MPQCOF GetPlayerAnimation(eHero hero, eMobMode mobMode, PlayerEquipment equipment)
            => cache.AddOrGetExisting($"COF::{hero}{mobMode.ToToken()}{equipment.HashKey}", () =>
            {
                var path = $"{ResourcePaths.PlayerAnimationBase}\\{hero.ToToken()}\\COF\\{hero.ToToken()}{mobMode.ToToken()}{equipment.WeaponClass.ToToken()}.cof";
                return MPQCOF.Load(mpqProvider.GetStream(path), Animations, hero, mobMode, equipment);
            }, new System.Runtime.Caching.CacheItemPolicy { Priority = System.Runtime.Caching.CacheItemPriority.NotRemovable });

        public MPQDCC GetPlayerDCC(MPQCOF.COFLayer cofLayer, PlayerEquipment equipment, Palette palette)
        {
            // TODO: Smarter hashing maybe
            return cache.AddOrGetExisting($"PlayerDCC::{cofLayer.CompositType.ToToken()}{cofLayer.COF.MobMode.ToToken()}{equipment.HashKey}", () =>
             {
                 byte[] binaryData;
                 
                 var streamPath = mpqProvider.GetCharacterDccPath(cofLayer.COF.Hero, cofLayer.COF.MobMode, cofLayer.CompositType, equipment);

                 // If stream path is null, there is nothing to load for this layer (this is NOT an error!)
                 if (streamPath == null)
                     return null;

                 using (var stream = mpqProvider.GetStream(streamPath))
                 {
                     if (stream == null)
                     {
                         log.Error($"Could not load Player DCC: {streamPath}");
                         return null;
                     }

                     binaryData = new byte[stream.Length];
                     stream.Read(binaryData, 0, (int)stream.Length);
                 }
                 return new MPQDCC(binaryData, palette);
             });
        }
    }
}
