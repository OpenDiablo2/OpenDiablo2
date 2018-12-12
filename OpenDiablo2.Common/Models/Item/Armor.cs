using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class Armor : Item 
    {
        public string Type { get; internal set; }
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
    }   
}
