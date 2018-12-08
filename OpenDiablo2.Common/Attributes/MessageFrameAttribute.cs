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
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Attributes
{
    /// <summary>
    /// Defines the class as a Message Frame. This is used by the client/server logic
    /// to decide how to serialize and deserialize networking objects.
    /// </summary>
    [AttributeUsage(AttributeTargets.All, Inherited = false, AllowMultiple = true)]
    public sealed class MessageFrameAttribute : Attribute
    {
        /// <summary>
        /// The type of message frame this class represents.
        /// </summary>
        public eMessageFrameType FrameType { get; private set; }

        /// <summary>
        /// Define this class as a message frame.
        /// </summary>
        /// <param name="frameType">The type of message frame this class represents.</param>
        public MessageFrameAttribute(eMessageFrameType frameType) => this.FrameType = frameType;


    }
}
