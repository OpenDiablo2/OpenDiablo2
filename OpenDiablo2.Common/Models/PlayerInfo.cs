/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.IO;
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
        public eMobMode MobMode { get; set; }
        public PlayerLocationDetails LocationDetails { get; set; }
        public PlayerEquipment Equipment { get; set; }

        public byte[] GetBytes()
        {
            using (var stream = new MemoryStream())
            using (var writer = new BinaryWriter(stream))
            {
                writer.Write((byte)Hero);
                writer.Write((byte)MobMode);
                writer.Write(Name);
                writer.Write(Equipment);
                writer.Write(LocationDetails.GetBytes());
                writer.Write(UID.ToByteArray());

                return stream.ToArray();
            }
        }

        public static PlayerInfo FromBytes(byte[] data, int offset = 0)
        {
            using (var stream = new MemoryStream(data))
            using (var reader = new BinaryReader(stream))
            {
                reader.ReadBytes(offset); // Skip

                var result = new PlayerInfo
                {
                    Hero = (eHero)reader.ReadByte(),
                    MobMode = (eMobMode)reader.ReadByte(),
                    Name = reader.ReadString(),
                    Equipment = reader.ReadPlayerEquipment(),
                    LocationDetails = PlayerLocationDetails.FromBytes(reader.ReadBytes(PlayerLocationDetails.SizeInBytes)),
                    UID = new Guid(reader.ReadBytes(16))
                };

                return result;
            }
        }

        public int SizeInBytes => 8 + Encoding.UTF8.GetByteCount(Name) + PlayerLocationDetails.SizeInBytes + 16;
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
                MobMode = source.MobMode,
                Equipment = source.Equipment
            };
    }
}
