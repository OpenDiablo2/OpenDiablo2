using System;
using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.PlayerInfo)]
    public sealed class MFPlayerInfo : IMessageFrame
    {
        public IEnumerable<PlayerInfo> PlayerInfos { get; set; } = new List<PlayerInfo>();

        public byte[] Data
        {
            get
            {
                var result = BitConverter.GetBytes(PlayerInfos.Count())
                .Concat(PlayerInfos.SelectMany(x => x.GetBytes()))
                .ToArray();
                return result;
            }

            set
            {
                var count = BitConverter.ToInt32(value, 0);
                var playerInfos = new List<PlayerInfo>();
                var offset = 4;
                for (var i = 0; i < count; i++)
                {
                    var playerInfo = PlayerInfo.FromBytes(value, offset);
                    playerInfos.Add(playerInfo);
                    offset += playerInfo.SizeInBytes;
                }

                PlayerInfos = playerInfos;
            }
        }

        public MFPlayerInfo() { }
        public MFPlayerInfo(IEnumerable<PlayerInfo> playerInfo)
        {
            this.PlayerInfos = playerInfo;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnPlayerInfo(clientHash, PlayerInfos);
    }

}
