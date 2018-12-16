using System;
using System.Collections.Generic;
using System.IO;
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

        public void LoadFrom(BinaryReader br)
        {
            var count = br.ReadUInt16();
            var playerInfos = new PlayerInfo[count];
            for (var i = 0; i < count; i++)
                playerInfos[i] = PlayerInfo.FromBytes(br);

            PlayerInfos = playerInfos;
        }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write((UInt16)PlayerInfos.Count());
            foreach (var playerInfo in PlayerInfos)
                playerInfo.WriteBytes(bw);
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
