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
        
    }

    public static class WeaponHelper
    {
        public static Weapon ToWeapon(this string[] row)
            => new Weapon
            {
                Name = row[0],
                Code = row[30],
                InvFile = row[45]
            };
    }   
}
