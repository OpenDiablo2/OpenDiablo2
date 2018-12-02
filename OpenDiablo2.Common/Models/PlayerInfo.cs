using System;
using System.Collections.Generic;
using System.Text;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Models
{
    public sealed class PlayerInfo
    {
        public string Name { get; set; }
        public eHero Hero { get; set; }
        public PlayerLocationDetails LocationDetails { get; set; }

        public byte[] GetBytes()
        {
            var result = new List<byte>();
            var nameBytes = Encoding.UTF8.GetBytes(Name);
            result.Add((byte)Hero);
            result.AddRange(BitConverter.GetBytes((Int32)nameBytes.Length));
            result.AddRange(nameBytes);
            result.AddRange(LocationDetails.GetBytes());
            return result.ToArray();
        }

        public static PlayerInfo FromBytes(byte[] data, int offset = 0)
        {
            var result = new PlayerInfo();
            result.Hero = (eHero)data[offset];
            var nameLength = BitConverter.ToInt32(data, offset + 1);
            result.Name = Encoding.UTF8.GetString(data, offset + 5, nameLength);
            result.LocationDetails = PlayerLocationDetails.FromBytes(data, offset + 5 + nameLength);
            return result;
        }

        public int SizeInBytes => 5 + Encoding.UTF8.GetByteCount(Name) + PlayerLocationDetails.SizeInBytes;
    }


    public static class PlayerInfoExtensions
    {
        public static PlayerInfo ToPlayerInfo(this PlayerState source)
            => new PlayerInfo
            {
                Hero = source.HeroType,
                LocationDetails = new PlayerLocationDetails
                {
                    PlayerId = source.Id,
                    PlayerX = source.GetPosition().X,
                    PlayerY = source.GetPosition().Y
                },
                Name = source.Name
            };
    }
}
