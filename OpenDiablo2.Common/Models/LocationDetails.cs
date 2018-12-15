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
    public sealed class LocationDetails : IMobLocation
    {
        public Guid UID { get; set; }
        public float X { get; set; }
        public float Y { get; set; }
        public eMovementType MovementType { get; set; }
        public int MovementDirection { get; set; }
        public float MovementSpeed { get; set; }
        public List<PointF> Waypoints { get; set; } = new List<PointF>();
        // TODO: They may not be on the same 'anchor map'...

        public void WriteBytes(BinaryWriter bw)
        {
            bw.Write(UID.ToByteArray());
            bw.Write((Single)X);
            bw.Write((Single)Y);
            bw.Write((Int32)MovementDirection);
            bw.Write((byte)MovementType);
            bw.Write((Single)MovementSpeed);
            bw.Write((Int16)(Waypoints?.Count ?? 0));
            for (var i = 0; i < Waypoints.Count; i++)
            {
                bw.Write((Single)Waypoints[i].X);
                bw.Write((Single)Waypoints[i].Y);
            }
        }

        public static LocationDetails FromBytes(BinaryReader br)
        {
            var result = new LocationDetails
            {
                UID = new Guid(br.ReadBytes(16)),
                X = br.ReadSingle(),
                Y = br.ReadSingle(),
                MovementDirection = br.ReadInt32(),
                MovementType = (eMovementType)br.ReadByte(),
                MovementSpeed = br.ReadSingle()
            };
            var numWaypoints = br.ReadInt16();
            result.Waypoints = new List<PointF>(numWaypoints);
            for(var i = 0; i < numWaypoints; i++)
            {
                result.Waypoints.Add(new PointF
                {
                    X = br.ReadSingle(),
                    Y = br.ReadSingle()
                });
            }
                
            return result;
        }
    }

    public static class PlayerLocationDetailsExtensions
    {
        public static LocationDetails ToPlayerLocationDetails(this PlayerState source)
        {
            var result = new LocationDetails
            {
                UID = source.UID,
                X = source.GetPosition().X,
                Y = source.GetPosition().Y,
                Waypoints = source.Waypoints,
                MovementType = source.MovementType,
                MovementSpeed = (source.MovementType == eMovementType.Running ? source.GetRunVelocity() : source.GetWalkVeloicty()) / 4f
            };
            return result;
        }
    }
}
