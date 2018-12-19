using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models.Mobs;
using System;
using System.IO;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.ChangeEquipment)]
    public sealed class MFChangeEquipment : IMessageFrame
    {
        public PlayerEquipment PlayerEquipment { get; internal set; }
        public Guid PlayerId { get; internal set; }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write(PlayerId.ToByteArray());
            bw.Write(PlayerEquipment);
        }

        public void LoadFrom(BinaryReader br)
        {
            PlayerId = new Guid(br.ReadBytes(16));
            PlayerEquipment = br.ReadPlayerEquipment();
        }

        public MFChangeEquipment() { }
        public MFChangeEquipment(Guid playerId, PlayerEquipment playerEquipment)
        {
            PlayerEquipment = playerEquipment;
            PlayerId = playerId;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnChangeEquipment(clientHash, PlayerId, PlayerEquipment);
        }
    }
}
