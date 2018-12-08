using System;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.FocusOnPlayer)]
    public sealed class MFFocusOnPlayer : IMessageFrame
    {
        public int PlayerToFocusOn { get; set; } = 0;

        public byte[] Data
        {
            get => BitConverter.GetBytes(PlayerToFocusOn);
            set => PlayerToFocusOn = BitConverter.ToInt32(value, 0);
        }

        public MFFocusOnPlayer() { }
        public MFFocusOnPlayer(int playerId)
        {
            this.PlayerToFocusOn = playerId;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnFocusOnPlayer(clientHash, PlayerToFocusOn);
    }
}
