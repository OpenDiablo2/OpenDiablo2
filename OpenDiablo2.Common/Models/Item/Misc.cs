using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class Misc : Item 
    {
        
    }

    public static class MiscHelper
    {
        public static Misc ToMisc(this string[] row)
            => new Misc
            {
                Name = row[0],
                Code = row[12],
                InvFile = row[21]
            };
    }   
}
