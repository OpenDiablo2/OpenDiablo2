using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.IO;
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
            Name = item.Name;
        }
    }

    public static class ItemInstanceHelper 
    {
        public static void Write(this BinaryWriter br, ItemInstance itemInstance)
        {
            if (itemInstance == null)
            {
                br.Write(false);
                return;
            }

            br.Write(true);

            if (itemInstance.Item is Weapon)
            {
                br.Write((byte)0);
                br.Write((Weapon)itemInstance.Item);
            }
            else if (itemInstance.Item is Armor)
            {
                br.Write((byte)1);
                br.Write((Armor)itemInstance.Item);
            }
            else if (itemInstance.Item is Misc)
            {
                br.Write((byte)2);
                br.Write((Misc)itemInstance.Item);
            }
            else throw new OpenDiablo2Exception("Unknown item type.");

            br.Write(itemInstance.Name);
            br.Write((Int16)itemInstance.Level);
            br.Write(itemInstance.Identified);
        }

        public static ItemInstance ReadItemInstance(this BinaryReader binaryReader)
        {
            if (!binaryReader.ReadBoolean())
                return null;

            Item item;
            switch(binaryReader.ReadByte())
            {
                case 0:
                    item = binaryReader.ReadItemWeapon();
                    break;
                case 1:
                    item = binaryReader.ReadItemArmor();
                    break;
                case 2:
                    item = binaryReader.ReadItemMisc();
                    break;
                default:
                    throw new OpenDiablo2Exception("Unknown item type.");
            }

            return new ItemInstance(item)
            {
                Name = binaryReader.ReadString(),
                Level = binaryReader.ReadInt16(),
                Identified = binaryReader.ReadBoolean()
            };
        }
    }
}
