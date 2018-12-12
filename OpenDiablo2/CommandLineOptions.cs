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

using CommandLine;

namespace OpenDiablo2
{
    public sealed class CommandLineOptions
    {
        /// <summary>
        /// The root path of the data files
        /// </summary>
        [Option('p', "datapath", Required = false, HelpText = "Specifies the root data path")]
        public string DataPath { get; set; }

        /// <summary>
        /// When true, the hardware cursor is used instead of the software one.
        /// </summary>
        [Option("hwmouse", Default = false, Required = false, HelpText = "Use the hardware mouse instead of software")]
        public bool HardwareMouse { get; set; }

        /// <summary>
        /// When hardware cursor mode is enabled, this changes the scale of the HW cursor
        /// </summary>
        [Option("mousescale", Default = 1, Required = false, HelpText = "When hardware mouse is enabled, this defines the pixel scale of the mouse. No effect for software mode")]
        public int MouseScale { get; set; }

        /// <summary>
        /// When true, the game runs in full screen.
        /// </summary>
        [Option('f', "fullscreen", Default = false, Required = false, HelpText = "When set, the game will start in full screen mode")]
        public bool FullScreen { get; set; }
    }
}
