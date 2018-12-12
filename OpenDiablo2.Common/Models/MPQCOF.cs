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
            public string ShieldCode { get; internal set; }
            public string WeaponCode { get; internal set; }

            // TODO: Move logic somewhere else.
            // TODO: Consider two hand weapons. 
            public string GetDCCPath(eArmorType armorType)
            {
                string result = null;
                var weaponClass = COF.WeaponClass;
                if(CompositType != eCompositType.RightArm && CompositType != eCompositType.RightHand)
                {
                    weaponClass = eWeaponClass.HandToHand;
                }

                if(CompositType == eCompositType.Shield)
                {
                    result = $"{ResourcePaths.PlayerAnimationBase}\\{COF.Hero.ToToken()}\\{CompositType.ToToken()}\\{COF.Hero.ToToken()}{CompositType.ToToken()}{ShieldCode}{COF.MobMode.ToToken()}{weaponClass.ToToken()}.dcc";
                }
                else if (CompositType == eCompositType.RightHand)
                {
                    result = $"{ResourcePaths.PlayerAnimationBase}\\{COF.Hero.ToToken()}\\{CompositType.ToToken()}\\{COF.Hero.ToToken()}{CompositType.ToToken()}{WeaponCode}{COF.MobMode.ToToken()}{weaponClass.ToToken()}.dcc";
                }
                else
                {
                    result = $"{ResourcePaths.PlayerAnimationBase}\\{COF.Hero.ToToken()}\\{CompositType.ToToken()}\\{COF.Hero.ToToken()}{CompositType.ToToken()}{armorType.ToToken()}{COF.MobMode.ToToken()}{weaponClass.ToToken()}.dcc";
                }
                
                return result;
            }

        }

        public eHero Hero { get; private set; }
        public eWeaponClass WeaponClass { get; private set; }
        public eMobMode MobMode { get; private set; }
        public List<AnimationData> Animations { get; private set; }

        public COFLayer[] Layers { get; private set; }
        public Dictionary<eCompositType, int> CompositLayers { get; private set; }
        public IEnumerable<eAnimationFrame> AnimationFrames { get; private set; }
        public eCompositType[] Priority { get; private set; }
        public int NumberOfDirections { get; internal set; }
        public int FramesPerDirection { get; internal set; }
        public int NumberOfLayers { get; internal set; }

        public static MPQCOF Load(Stream stream, Dictionary<string, List<AnimationData>> animations, eHero hero, eWeaponClass weaponClass, eMobMode mobMode, string ShieldCode, string weaponCode)
        {
            var result = new MPQCOF
            {
                WeaponClass = weaponClass,
                MobMode = mobMode,
                Hero = hero
            };

            var br = new BinaryReader(stream);

            result.NumberOfLayers = br.ReadByte();
            result.FramesPerDirection = br.ReadByte();
            result.NumberOfDirections = br.ReadByte(); // Number of directions

            br.ReadBytes(25); // Skip 25 unknown bytes...

            var layers = new List<COFLayer>();
            result.CompositLayers = new Dictionary<eCompositType, int>();

            for (var layerIdx = 0; layerIdx < result.NumberOfLayers; layerIdx++)
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
                layer.WeaponClass = Encoding.ASCII.GetString(br.ReadBytes(4)).Trim('\0').ToWeaponClass();
                layer.ShieldCode = ShieldCode;
                layer.WeaponCode = weaponCode;
                layers.Add(layer);
                result.CompositLayers[layer.CompositType] = layerIdx;
            }
            result.Layers = layers.ToArray();
            result.AnimationFrames = br.ReadBytes(result.FramesPerDirection).Select(x => (eAnimationFrame)x);
            result.Priority = br.ReadBytes(result.FramesPerDirection * result.NumberOfLayers * result.NumberOfDirections).Select(x => (eCompositType)x).ToArray();

            var cofName = $"{hero.ToToken()}{mobMode.ToToken()}{weaponClass.ToToken()}".ToUpper();
            result.Animations = animations[cofName];
            br.Dispose();
            return result;
        }
    }
}
