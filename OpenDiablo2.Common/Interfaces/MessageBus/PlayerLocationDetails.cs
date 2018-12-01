using System;
using System.Collections.Generic;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Interfaces.MessageBus
{
    public sealed class PlayerLocationDetails
    {
        public int PlayerId { get; set; }
        public float PlayerX { get; set; }
        public float PlayerY { get; set; }
        // TODO: They may not be on the same 'anchor map'...

        public byte[] GetBytes()
        {
            var result = new List<byte>();
            result.AddRange(BitConverter.GetBytes(PlayerId));
            result.AddRange(BitConverter.GetBytes(PlayerX));
            result.AddRange(BitConverter.GetBytes(PlayerY));
            return result.ToArray();
        }

        public static PlayerLocationDetails FromBytes(byte[] data, int offset = 0)
            => new PlayerLocationDetails
            {
                PlayerId = BitConverter.ToInt32(data, offset + 0),
                PlayerX = BitConverter.ToSingle(data, offset + 4),
                PlayerY = BitConverter.ToSingle(data, offset + 8)
            };

        public static int SizeInBytes => 12;
    }

    public static class PlayerLocationDetailsExtensions
    {
        public static PlayerLocationDetails ToPlayerLocationDetails(this PlayerState source)
            => new PlayerLocationDetails
            {
                PlayerId = source.Id,
                PlayerX = source.GetPosition().X,
                PlayerY = source.GetPosition().Y
            };
    }
}
