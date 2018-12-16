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
using System.Drawing;
using System.IO;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Models
{
    public sealed class PlayerInfo : IMobLocation
    {
        public Guid UID { get; set; }
        public string Name { get; set; }
        public eHero Hero { get; set; }
        public eMobMode MobMode { get; set; }
        public PlayerEquipment Equipment { get; set; }
        public float X { get; set; }
        public float Y { get; set; }
        public float MovementSpeed { get; set; }
        public int MovementDirection { get; set; }
        public List<PointF> Waypoints { get; set; }
        public eMovementType MovementType { get; set; }

        public void WriteBytes(BinaryWriter writer)
        {
            writer.Write(UID.ToByteArray());
            writer.Write(Name);
            writer.Write((byte)Hero);
            writer.Write((byte)MobMode);
            writer.Write(Equipment);
            writer.Write((Single)X);
            writer.Write((Single)Y);
            writer.Write((Single)MovementSpeed);
            writer.Write((byte)MovementDirection);
            writer.Write((byte)MovementType);
            writer.Write((UInt16)Waypoints.Count);
            foreach (var waypoint in Waypoints)
            {
                writer.Write((Single)waypoint.X);
                writer.Write((Single)waypoint.Y);
            }
        }

        public static PlayerInfo FromBytes(BinaryReader br)
        {
            var result = new PlayerInfo
            {
                UID = new Guid(br.ReadBytes(16)),
                Name = br.ReadString(),
                Hero = (eHero)br.ReadByte(),
                MobMode = (eMobMode)br.ReadByte(),
                Equipment = br.ReadPlayerEquipment(),
                X = br.ReadSingle(),
                Y = br.ReadSingle(),
                MovementSpeed = br.ReadSingle(),
                MovementDirection = br.ReadByte(),
                MovementType = (eMovementType)br.ReadByte()
            };

            var numWaypoints = br.ReadUInt16();
            result.Waypoints = new List<PointF>();
            for (var i = 0; i < numWaypoints; i++)
                result.Waypoints.Add(new PointF
                {
                    X = br.ReadSingle(),
                    Y= br.ReadSingle()
                });

            return result;
        }
    }


    public static class PlayerInfoExtensions
    {
        // Map the player state to a PlayerInfo network package object.
        public static PlayerInfo ToPlayerInfo(this PlayerState source)
            => new PlayerInfo
            {
                UID = source.UID,
                Name = source.Name,
                Hero = source.HeroType,
                MobMode = source.MobMode,
                Equipment = source.Equipment,
                X = source.X,
                Y = source.Y,
                MovementSpeed = source.MovementSpeed,
                MovementDirection = source.MovementDirection,
                MovementType = source.MovementType,
                Waypoints = source.Waypoints
            };
    }
}
