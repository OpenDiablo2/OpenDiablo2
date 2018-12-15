using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System.IO;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.JoinGame)]
    public sealed class MFJoinGame : IMessageFrame
    {
        public string PlayerName { get; set; }
        public eHero HeroType { get; set; }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write((byte)HeroType);
            bw.Write(PlayerName);
        }

        public void LoadFrom(BinaryReader br)
        {
            HeroType = (eHero)br.ReadByte();
            PlayerName = br.ReadString();
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
