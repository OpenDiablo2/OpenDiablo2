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
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Common.Services
{
    public sealed class MobMovementService
    {
        const double Rad2Deg = 180.0 / Math.PI;
        IMobLocation mobLocation;

        public MobMovementService(IMobLocation mobLocation)
        {
            this.mobLocation = mobLocation;
        }

        /// <summary>Moves the mob along the set of waypoints</summary>
        /// <param name="seconds">The amount of time to process</param>
        public void CalculateMovement(float seconds)
        {
            if (mobLocation.Waypoints.Count == 0)
                return;

            if (!mobLocation.Waypoints.Any())
                return;

            // Determine how far we can move this round
            var moveRate = seconds * mobLocation.MovementSpeed;

            // Keep iterating waypoints while we can
            while (moveRate > 0 && mobLocation.Waypoints.Any())
            {
                // Determine the next target location
                var targetX = mobLocation.Waypoints.First().X;
                var targetY = mobLocation.Waypoints.First().Y;

                // If we're already on this spot, we're done
                if (targetX == mobLocation.X && targetY == mobLocation.Y)
                {
                    mobLocation.Waypoints.RemoveAt(0);
                    continue;
                }

                // Calculate our direction
                var cursorDirection = Math.Round(Math.Atan2(targetY - mobLocation.Y, targetX - mobLocation.X) * Rad2Deg);
                if (cursorDirection < 0)
                    cursorDirection += 360;
                var actualDirection = (byte)(cursorDirection / 22);
                if (actualDirection >= 16)
                    actualDirection -= 16;
                mobLocation.MovementDirection = actualDirection;

                // Determine how far we are away from the waypoint
                var maxDistance = (float)Math.Sqrt(Math.Pow((targetX - mobLocation.X), 2) + Math.Pow((targetY - mobLocation.Y), 2));

                // If we can move beyond the distance we need to to hit the final waypoint, we are doing moving
                if (moveRate >= maxDistance && mobLocation.Waypoints.Count == 1)
                {
                    mobLocation.X = mobLocation.Waypoints.First().X;
                    mobLocation.Y = mobLocation.Waypoints.First().Y;
                    mobLocation.Waypoints.Clear();
                    break;
                }

                // Calculate how far we're actually going to move
                var distance = (float)Math.Min(maxDistance, moveRate);

                // Calculate the point to land on between source and target locations
                float tx = targetX - mobLocation.X;
                float ty = targetY - mobLocation.Y;
                mobLocation.X = (mobLocation.X + distance * tx / maxDistance);
                mobLocation.Y = (mobLocation.Y + distance * ty / maxDistance);

                // Subtract out the distance we moved (in case we need to move through another node)
                moveRate = (float)Math.Max(0, moveRate - distance);

                // If we are on the waypoint, remove it from the list
                if (targetX == mobLocation.X && targetY == mobLocation.Y)
                    mobLocation.Waypoints.RemoveAt(0);
            }

            // Stop walking if we have run out of waypoints
            if (mobLocation.Waypoints.Count == 0)
                mobLocation.MovementType = eMovementType.Stopped;
        }
    }
}
