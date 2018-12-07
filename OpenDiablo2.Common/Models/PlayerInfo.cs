using System;
using System.Collections.Generic;
using System.Text;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Models
{
    public sealed class PlayerInfo
    {
        public Guid UID { get; set; }
        public string Name { get; set; }
        public eHero Hero { get; set; }
        public eWeaponClass WeaponClass { get; set; }
        public eArmorType ArmorType { get; set; }
        public eMobMode MobMode { get; set; }
        public PlayerLocationDetails LocationDetails { get; set; }

        public byte[] GetBytes()
        {
            var result = new List<byte>();
            var nameBytes = Encoding.UTF8.GetBytes(Name);
            result.Add((byte)Hero);
            result.Add((byte)WeaponClass);
            result.Add((byte)ArmorType);
            result.Add((byte)MobMode);
            result.AddRange(BitConverter.GetBytes((Int32)nameBytes.Length));
            result.AddRange(nameBytes);
            result.AddRange(LocationDetails.GetBytes());
            result.AddRange(UID.ToByteArray());
            return result.ToArray();
        }

        public static PlayerInfo FromBytes(byte[] data, int offset = 0)
        {
            var result = new PlayerInfo();
            result.Hero = (eHero)data[offset];
            result.WeaponClass= (eWeaponClass)data[offset + 1];
            result.ArmorType = (eArmorType)data[offset + 2];
            result.MobMode = (eMobMode)data[offset + 3];
            var nameLength = BitConverter.ToInt32(data, offset + 4);
            result.Name = Encoding.UTF8.GetString(data, offset + 8, nameLength);
            result.LocationDetails = PlayerLocationDetails.FromBytes(data, offset + 8 + nameLength);
            var uidBytes = new byte[16];
            Array.Copy(data, offset + 8 + nameLength + PlayerLocationDetails.SizeInBytes + 1, uidBytes, 0, 16);
            result.UID = new Guid(uidBytes);
            return result;
        }

        public int SizeInBytes => 5 + Encoding.UTF8.GetByteCount(Name) + PlayerLocationDetails.SizeInBytes;
    }


    public static class PlayerInfoExtensions
    {
        // Map the player state to a PlayerInfo network package object.
        public static PlayerInfo ToPlayerInfo(this PlayerState source)
            => new PlayerInfo
            {
                UID = source.UID,
                Hero = source.HeroType,
                LocationDetails = new PlayerLocationDetails
                {
                    PlayerId = source.Id,
                    PlayerX = source.GetPosition().X,
                    PlayerY = source.GetPosition().Y,
                },
                Name = source.Name,
                WeaponClass = source.WeaponClass,
                ArmorType = source.ArmorType,
                MobMode = source.MobMode
            };
    }
}
