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
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Models
{
    public sealed class PlayerLocationDetails
    {
        public int PlayerId { get; set; }
        public float PlayerX { get; set; }
        public float PlayerY { get; set; }
        public eMovementType MovementType { get; set; }
        public int MovementDirection { get; set; }
        public float MovementSpeed { get; set; }
        // TODO: They may not be on the same 'anchor map'...

        public byte[] GetBytes()
        {
            var result = new List<byte>();
            result.AddRange(BitConverter.GetBytes(PlayerId));
            result.AddRange(BitConverter.GetBytes(PlayerX));
            result.AddRange(BitConverter.GetBytes(PlayerY));
            result.AddRange(BitConverter.GetBytes(MovementDirection));
            result.AddRange(BitConverter.GetBytes((byte)MovementType));
            result.AddRange(BitConverter.GetBytes(MovementSpeed));
            return result.ToArray();
        }

        public static PlayerLocationDetails FromBytes(byte[] data, int offset = 0)
        {
            var result = new PlayerLocationDetails
            {
                PlayerId = BitConverter.ToInt32(data, offset + 0),
                PlayerX = BitConverter.ToSingle(data, offset + 4),
                PlayerY = BitConverter.ToSingle(data, offset + 8),
                MovementDirection = BitConverter.ToInt32(data, offset + 12),
                MovementType = (eMovementType)data[offset + 16],
                MovementSpeed = BitConverter.ToSingle(data, offset + 18)
            };
            return result;
        }
        public static int SizeInBytes => 22;
    }

    public static class PlayerLocationDetailsExtensions
    {
        public static PlayerLocationDetails ToPlayerLocationDetails(this PlayerState source)
        {
            var result = new PlayerLocationDetails
            {
                PlayerId = source.Id,
                PlayerX = source.GetPosition().X,
                PlayerY = source.GetPosition().Y,
                MovementType = source.MovementType,
                MovementDirection = source.MovementDirection,
                MovementSpeed = (source.MovementType == eMovementType.Running ? source.GetRunVelocity() : source.GetWalkVeloicty()) / 4f
            };
            return result;
        }
    }
}
