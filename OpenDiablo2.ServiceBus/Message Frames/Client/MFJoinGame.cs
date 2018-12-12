using System;
using System.Collections;
using System.IO;
using System.Linq;
using System.Runtime.Serialization.Formatters.Binary;
using System.Text;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
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
                using (var stream = new MemoryStream())
                using (var writer = new BinaryWriter(stream)) {
                    writer.Write((byte)HeroType);
                    writer.Write(PlayerName);

                    return stream.ToArray();
                }
            }
                
            set
            {
                using(var stream = new MemoryStream(value))
                using(var reader = new BinaryReader(stream))
                {
                    HeroType = (eHero)reader.ReadByte();
                    PlayerName = reader.ReadString();
                }
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
