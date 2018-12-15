using System;
using System.IO;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.FocusOnPlayer)]
    public sealed class MFFocusOnPlayer : IMessageFrame
    {
        public Guid PlayerToFocusOn { get; set; } = Guid.Empty;

        public MFFocusOnPlayer() { }
        public MFFocusOnPlayer(Guid playerId)
        {
            this.PlayerToFocusOn = playerId;
        }

        public void LoadFrom(BinaryReader br)
        {
            PlayerToFocusOn = new Guid(br.ReadBytes(16));
        }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write(PlayerToFocusOn.ToByteArray());
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnFocusOnPlayer(clientHash, PlayerToFocusOn);
    }
}
