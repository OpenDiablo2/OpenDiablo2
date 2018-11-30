using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public abstract class Item
    {
        public string Code { get; internal set; }
        public string Name { get; internal set; }
        public string InvFile { get; internal set; }
    }
}
