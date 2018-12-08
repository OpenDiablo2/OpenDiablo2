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

using OpenDiablo2.Common.Enums;
using System.Drawing;

namespace OpenDiablo2.Common.Extensions
{
    public static class EnumExtensions
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        public static Point GetOffset(this ePanelFrameType value)
        {
            switch (value)
            {
                case ePanelFrameType.Left:
                    return new Point(80, 63);
                case ePanelFrameType.Right:
                    return new Point(400, 63);
                case ePanelFrameType.Center:
                    return new Point(0, 0);
            }

            log.Warn($"Unknown panel positon, {value}");
            return default(Point);
        }
    }
}
