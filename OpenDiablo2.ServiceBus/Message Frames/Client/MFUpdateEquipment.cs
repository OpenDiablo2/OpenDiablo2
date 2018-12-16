using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models.Mobs;
using System.IO;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.JoinGame)]
    public sealed class MFUpdateEquipment : IMessageFrame
    {
        public PlayerEquipment PlayerEquipment { get; internal set; }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write(PlayerEquipment);
        }

        public void LoadFrom(BinaryReader br)
        {
            PlayerEquipment = br.ReadPlayerEquipment();
        }

        public MFUpdateEquipment() { }
        public MFUpdateEquipment(PlayerEquipment playerEquipment)
        {
            PlayerEquipment = playerEquipment;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnUpdateEquipment(clientHash, PlayerEquipment);
        }
    }
}
