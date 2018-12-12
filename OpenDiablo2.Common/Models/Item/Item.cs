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
    public abstract class Item
    {
        public string Code { get; internal set; } // Internal code
        public string Name { get; internal set; } // Item name
        public string InvFile { get; internal set; } // Sprite used for the inventory and mouse cursor

        internal void Write(BinaryWriter binaryWriter)
        {
            binaryWriter.Write(Code);
            binaryWriter.Write(Name);
            binaryWriter.Write(InvFile);
        }

        internal static void Read(BinaryReader binaryReader, Item source)
        {
            source.Code = binaryReader.ReadString();
            source.Name = binaryReader.ReadString();
            source.InvFile = binaryReader.ReadString();
        }
    }
}
