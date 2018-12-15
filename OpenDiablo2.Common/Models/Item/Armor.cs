using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class Armor : Item
    {
        public string Type { get; internal set; }

        public Dictionary<eCompositType, eArmorType> ArmorTypes = new Dictionary<eCompositType, eArmorType>() {
            {eCompositType.RightArm, eArmorType.Lite},
            {eCompositType.LeftArm, eArmorType.Lite},
            {eCompositType.Torso, eArmorType.Lite},
            {eCompositType.Legs, eArmorType.Lite},
            {eCompositType.Special1, eArmorType.Lite},
            {eCompositType.Special2, eArmorType.Lite},
        };
    }

    public static class ArmorHelper
    {
        public static Armor ToArmor(this string[] row)
        {

            var Armor = new Armor
            {
                Name = row[0],
                Code = row[17],
                InvFile = row[34],
                Type = row[48]
            };

            if(Armor.Type == "tors")
            {
                Armor.ArmorTypes[eCompositType.RightArm] = (eArmorType)Enum.Parse(typeof(eArmorType), row[37]);
                Armor.ArmorTypes[eCompositType.LeftArm] = (eArmorType)Enum.Parse(typeof(eArmorType), row[38]);
                Armor.ArmorTypes[eCompositType.Torso] = (eArmorType)Enum.Parse(typeof(eArmorType), row[39]);
                Armor.ArmorTypes[eCompositType.Legs] = (eArmorType)Enum.Parse(typeof(eArmorType), row[40]);
                Armor.ArmorTypes[eCompositType.Special1] = (eArmorType)Enum.Parse(typeof(eArmorType), row[41]);
                Armor.ArmorTypes[eCompositType.Special2] = (eArmorType)Enum.Parse(typeof(eArmorType), row[42]);
            }
            
            return Armor;
         }

        public static void Write(this BinaryWriter binaryWriter, Armor armor)
        {
            (armor as Item).Write(binaryWriter);
            binaryWriter.Write(armor.Type);

            // Assuming order will be fine
            foreach(var armorType in armor.ArmorTypes)
            {
                binaryWriter.Write((byte)armorType.Value);
            }
        }

        public static Armor ReadItemArmor(this BinaryReader binaryReader)
        {
            var result = new Armor();
            Item.Read(binaryReader, result);
            result.Type = binaryReader.ReadString();

            result.ArmorTypes[eCompositType.RightArm] = (eArmorType)binaryReader.ReadByte();
            result.ArmorTypes[eCompositType.LeftArm] = (eArmorType)binaryReader.ReadByte();
            result.ArmorTypes[eCompositType.Torso] = (eArmorType)binaryReader.ReadByte();
            result.ArmorTypes[eCompositType.Legs] = (eArmorType)binaryReader.ReadByte();
            result.ArmorTypes[eCompositType.Special1] = (eArmorType)binaryReader.ReadByte();
            result.ArmorTypes[eCompositType.Special2] = (eArmorType)binaryReader.ReadByte();
            
            return result;
        }
    }   
}
