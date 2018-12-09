using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    /**
     * Base Item class, contains common attributes for all item types.
     **/
    public class ItemInstance : IItemInstance
    {
        public Item Item { get; internal set; }
        public string Name { get; internal set; }
        public int Level { get; internal set; }
        public bool Identified { get; internal set; }

        public ItemInstance(Item item)
        {
            Item = item;
        }
    }
}
