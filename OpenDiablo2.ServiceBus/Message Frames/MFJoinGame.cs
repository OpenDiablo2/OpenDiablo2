using System;
using System.Linq;
using System.Text;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames
{
    [MessageFrame(eMessageFrameType.JoinGame)]
    public sealed class MFJoinGame : IMessageFrame
    {
        public Guid PlayerId { get; set; }
        public string PlayerName { get; set; }
        public byte[] Data
        {
            get
            {
                return PlayerId.ToByteArray()
                    .Concat(BitConverter.GetBytes((UInt16)PlayerName.Length))
                    .Concat(Encoding.UTF8.GetBytes(PlayerName))
                    .ToArray();
            }

            set
            {

                PlayerId = new Guid(value.Take(16).ToArray());
                var playerNameLen = BitConverter.ToUInt16(value, 16);
                PlayerName = Encoding.UTF8.GetString(value, 18, value.Length - 18);

                if (PlayerName.Length != playerNameLen)
                    throw new ApplicationException("Invalid player length!");
            }
        }

        public MFJoinGame() { }
        public MFJoinGame(string playerName)
        {
            PlayerId = Guid.NewGuid();
            PlayerName = playerName;
        }

        public void Process(object sender, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnJoinGame(sender, PlayerId, PlayerName);
        }
    }
}
