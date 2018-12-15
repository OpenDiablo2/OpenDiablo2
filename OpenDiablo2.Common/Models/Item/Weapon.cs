using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class Weapon : Item 
    {
        public string WeaponClass { get; internal set; }
    }

    public static class WeaponHelper
    {
        public static Weapon ToWeapon(this string[] row)
            => new Weapon
            {
                Name = row[0],
                Code = row[3],
                WeaponClass = row[37],
                InvFile = row[48]
            };

        public static void Write(this BinaryWriter binaryWriter, Weapon weapon)
        {
            (weapon as Item).Write(binaryWriter);
            binaryWriter.Write(weapon.WeaponClass);
        }

        public static Weapon ReadItemWeapon(this BinaryReader binaryReader)
        {
            var result = new Weapon();
            Item.Read(binaryReader, result);
            result.WeaponClass = binaryReader.ReadString();
            return result;
        }
    }   
}
