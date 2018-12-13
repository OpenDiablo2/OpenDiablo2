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
        public string Type { get; internal set; } = "FIXME"; // TODO: Fix this please
    }

    public static class ArmorHelper
    {
        public static Armor ToArmor(this string[] row)
            => new Armor
            {
                Name = row[0],
                Code = row[17],
                InvFile = row[33]
            };

        public static void Write(this BinaryWriter binaryWriter, Armor armor)
        {
            (armor as Item).Write(binaryWriter);
            binaryWriter.Write(armor.Type);
        }

        public static Armor ReadItemArmor(this BinaryReader binaryReader)
        {
            var result = new Armor();
            Item.Read(binaryReader, result);
            result.Type = binaryReader.ReadString();
            return result;
        }
    }   
}
