using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public sealed class PlayerEquipment
    {
        public ItemInstance Head { get; set; }
        public ItemInstance Neck { get; set; }
        public ItemInstance Torso { get; set; }
        public ItemInstance RightArm { get; set; }
        public ItemInstance LeftArm { get; set; }
        public ItemInstance RightRing { get; set; }
        public ItemInstance LeftRing { get; set; }
        public ItemInstance Feet { get; set; }
        public ItemInstance Belt { get; set; }
        public ItemInstance Gloves { get; set; }

        public eWeaponClass WeaponClass
        {
            get
            {
                if (LeftArm?.Item is Weapon)
                    return ((Weapon)LeftArm.Item).WeaponClass.ToWeaponClass();
                else if (RightArm?.Item is Weapon)
                    return ((Weapon)RightArm.Item).WeaponClass.ToWeaponClass();
                else
                    return eWeaponClass.HandToHand;
            }
        }

        public string HashKey
            => $"{Head?.Item.Name}{Neck?.Item.Name}{Torso?.Item.Name}{RightArm?.Item.Name}{LeftArm?.Item.Name}" +
            $"{RightRing?.Item.Name}{LeftRing?.Item.Name}{Feet?.Item.Name}{Belt?.Item.Name}{Gloves?.Item.Name}";

        public void EquipItem(string slot, ItemInstance item)
        {
            switch (slot.ToLower())
            {
                case "head":
                    Head = item;
                    break;
                case "neck":
                    Neck = item;
                    break;
                case "tors":
                    Torso = item;
                    break;
                case "rarm":
                    RightArm = item;
                    break;
                case "larm":
                    LeftArm = item;
                    break;
                case "rrin":
                    RightRing = item;
                    break;
                case "lrin":
                    LeftRing = item;
                    break;
                case "belt":
                    Belt = item;
                    break;
                case "feet":
                    Feet = item;
                    break;
                case "glov":
                    Gloves = item;
                    break;
                default:
                    throw new OpenDiablo2Exception($"Unknown slot name '{slot}'!");
            }
        }
    }

    public static class PlayerEquipmentHelper
    {
        public static void Write(this BinaryWriter binaryWriter, PlayerEquipment source)
        {
            binaryWriter.Write(source.Head);
            binaryWriter.Write(source.Neck);
            binaryWriter.Write(source.Torso);
            binaryWriter.Write(source.RightArm);
            binaryWriter.Write(source.LeftArm);
            binaryWriter.Write(source.RightRing);
            binaryWriter.Write(source.LeftRing);
            binaryWriter.Write(source.Belt);
            binaryWriter.Write(source.Feet);
            binaryWriter.Write(source.Gloves);
        }

        public static PlayerEquipment ReadPlayerEquipment(this BinaryReader binaryReader)
        {
            return new PlayerEquipment
            {
                Head = binaryReader.ReadItemInstance(),
                Neck = binaryReader.ReadItemInstance(),
                Torso = binaryReader.ReadItemInstance(),
                RightArm = binaryReader.ReadItemInstance(),
                LeftArm = binaryReader.ReadItemInstance(),
                RightRing = binaryReader.ReadItemInstance(),
                LeftRing = binaryReader.ReadItemInstance(),
                Belt = binaryReader.ReadItemInstance(),
                Feet = binaryReader.ReadItemInstance(),
                Gloves = binaryReader.ReadItemInstance()
            };
        }
    }
}
