using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
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
                Code = row[2],
                WeaponClass = row[34],
                InvFile = row[45]
            };
    }   
}
