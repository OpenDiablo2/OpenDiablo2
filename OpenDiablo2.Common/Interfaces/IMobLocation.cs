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

using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMobLocation
    {
        float X { get; set; }
        float Y { get; set; }
        float MovementSpeed { get; set; }
        int MovementDirection { get; set; }
        List<PointF> Waypoints { get; set; }
        eMovementType MovementType { get; set; }
    }

    public static class MobLocationHelper
    {
        public static void CopyMobLocationDetailsTo(this IMobLocation source, IMobLocation dest)
        {
            dest.X = source.X;
            dest.Y = source.Y;
            dest.MovementSpeed = source.MovementSpeed;
            dest.MovementDirection = source.MovementDirection;
            dest.Waypoints = source.Waypoints; // TODO: do we need to do a literaly copy here?
            dest.MovementType = source.MovementType;
        }
    }

}
