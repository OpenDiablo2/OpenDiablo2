using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using System.IO;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.UpdateEquipment)]
    public sealed class MFUpdateEquipment : IMessageFrame
    {
        public ItemInstance ItemInstance { get; internal set; }
        public string Slot { get; private set; }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write(Slot);
            bw.Write(ItemInstance);
        }

        public void LoadFrom(BinaryReader br)
        {
            Slot = br.ReadString();
            ItemInstance = br.ReadItemInstance();
        }

        public MFUpdateEquipment() { }
        public MFUpdateEquipment(string slot, ItemInstance itemInstance)
        {
            Slot = slot;
            ItemInstance = itemInstance;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnUpdateEquipment(clientHash, Slot, ItemInstance);
        }
    }
}
