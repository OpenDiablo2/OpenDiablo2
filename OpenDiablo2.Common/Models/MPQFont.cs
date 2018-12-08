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
using System.Text;
using OpenDiablo2.Common.Exceptions;

namespace OpenDiablo2.Common.Models
{
    public sealed class MPQFont
    {
        public ImageSet FontImageSet { get; internal set; }
        public Dictionary<char, Size> CharacterMetric { get; internal set; } = new Dictionary<char, Size>();

        public static MPQFont LoadFromStream(Stream fontStream, Stream tableStream)
        {
            var result = new MPQFont
            {
                FontImageSet = ImageSet.LoadFromStream(fontStream)
            };

            var br = new BinaryReader(tableStream);
            var wooCheck = Encoding.UTF8.GetString(br.ReadBytes(4));
            if (wooCheck != "Woo!")
                throw new OpenDiablo2Exception("Error loading font. Missing the Woo!");
            br.ReadBytes(8);

            while (tableStream.Position < tableStream.Length)
            {
                br.ReadBytes(3);
                var size = new Size(br.ReadByte(), br.ReadByte());
                br.ReadBytes(3);
                var charCode = (char)br.ReadByte();
                result.CharacterMetric[charCode] = size;
                br.ReadBytes(5);

            }

            return result;
        }
    }
}
