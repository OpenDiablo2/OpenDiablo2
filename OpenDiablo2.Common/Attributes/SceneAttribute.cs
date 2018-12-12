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
    /// Defines this class as a scene.
    /// </summary>
    [AttributeUsage(AttributeTargets.Class, Inherited = false, AllowMultiple = false)]
    public sealed class SceneAttribute : Attribute
    {
        /// <summary>
        /// Defines the type of scene that this class represents.
        /// </summary>
        public eSceneType SceneType { get; }

        /// <summary>
        /// Defines this class as a scene type
        /// </summary>
        /// <param name="sceneType"></param>
        public SceneAttribute(eSceneType sceneType) => SceneType = sceneType;
    }
}
