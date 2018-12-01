using System;
using System.Linq;
using System.Text;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.JoinGame)]
    public sealed class MFJoinGame : IMessageFrame
    {
        public string PlayerName { get; set; }
        public eHero HeroType { get; set; }

        public byte[] Data
        {
            get
            {
                return new byte[] { (byte)HeroType }
                    .Concat(BitConverter.GetBytes((UInt16)PlayerName.Length))
                    .Concat(Encoding.UTF8.GetBytes(PlayerName))
                    .ToArray();
            }

            set
            {
                HeroType = (eHero)value[0];
                var playerNameLen = BitConverter.ToUInt16(value, 1);
                PlayerName = Encoding.UTF8.GetString(value, 3, value.Length - 3);

                if (PlayerName.Length != playerNameLen)
                    throw new ApplicationException("Invalid player length!");
            }
        }

        public MFJoinGame() { }
        public MFJoinGame(string playerName, eHero heroType)
        {
            PlayerName = playerName;
            HeroType = heroType;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnJoinGame(clientHash, HeroType, PlayerName);
        }
    }
}
