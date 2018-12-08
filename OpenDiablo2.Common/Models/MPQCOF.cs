using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQCOF
    {
        public class COFLayer
        {
            public MPQCOF COF { get; internal set; }
            public eCompositType CompositType { get; internal set; }
            public byte Shadow { get; internal set; }
            public bool IsTransparent { get; internal set; }
            public eDrawEffect DrawEffect { get; internal set; }
            public eWeaponClass WeaponClass { get; internal set; }

            public string GetDCCPath(eArmorType armorType)
            {
                var result = $"{ResourcePaths.PlayerAnimationBase}\\{COF.Hero.ToToken()}\\{CompositType.ToToken()}\\{COF.Hero.ToToken()}{CompositType.ToToken()}{armorType.ToToken()}{COF.MobMode.ToToken()}{COF.WeaponClass.ToToken()}.dcc";
                return result;
            }

        }

        public eHero Hero { get; private set; }
        public eWeaponClass WeaponClass { get; private set; }
        public eMobMode MobMode { get; private set; }
        public List<AnimationData> Animations { get; private set; }

        public IEnumerable<COFLayer> Layers { get; private set; }
        public IEnumerable<eAnimationFrame> AnimationFrames { get; private set; }

        public static MPQCOF Load(Stream stream, Dictionary<string, List<AnimationData>> animations, eHero hero, eWeaponClass weaponClass, eMobMode mobMode)
        {
            var result = new MPQCOF
            {
                WeaponClass = weaponClass,
                MobMode = mobMode,
                Hero = hero
            };

            var br = new BinaryReader(stream);

            var numLayers = br.ReadByte();
            var framesPerDir = br.ReadByte();
            br.ReadByte(); // Number of directions

            br.ReadBytes(25); // Skip 25 unknown bytes...

            var layers = new List<COFLayer>();
            for (var layerIdx = 0; layerIdx < numLayers; layerIdx++)
            {
                var layer = new COFLayer
                {
                    COF = result,
                    CompositType = (eCompositType)br.ReadByte(),
                    Shadow = br.ReadByte()
                };
                br.ReadByte(); // Unknown
                layer.IsTransparent = br.ReadByte() != 0;
                layer.DrawEffect = (eDrawEffect)br.ReadByte();
                layers.Add(layer);
                layer.WeaponClass = Encoding.ASCII.GetString(br.ReadBytes(4)).Trim('\0').ToWeaponClass();
            }
            result.Layers = layers;
            result.AnimationFrames = br.ReadBytes(framesPerDir).Select(x => (eAnimationFrame)x);

            var cofName = $"{hero.ToToken()}{mobMode.ToToken()}{weaponClass.ToToken()}".ToUpper();
            result.Animations = animations[cofName];
            br.Dispose();
            return result;
        }
    }
}
