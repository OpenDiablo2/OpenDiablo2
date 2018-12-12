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
using System.Collections.Immutable;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public class ItemContainerLayout
    {
        public string ResourceName { get; internal set; }
        public string PaletteName { get; internal set; } = Palettes.Units;
        public int BaseFrame { get; internal set; } = 0;

        public static ImmutableDictionary<eItemContainerType, ItemContainerLayout> Values { get; } = new Dictionary<eItemContainerType, ItemContainerLayout>
        {
            {eItemContainerType.Helm,  new ItemContainerLayout { ResourceName = ResourcePaths.HelmGlovePlaceholder,  BaseFrame = 1 } },
            {eItemContainerType.Amulet, new ItemContainerLayout{ ResourceName = ResourcePaths.RingAmuletPlaceholder  } },
            {eItemContainerType.Armor, new ItemContainerLayout { ResourceName = ResourcePaths.ArmorPlaceholder } },
            {eItemContainerType.Weapon, new ItemContainerLayout { ResourceName = ResourcePaths.WeaponsPlaceholder } },
            {eItemContainerType.Belt, new ItemContainerLayout { ResourceName = ResourcePaths.BeltPlaceholder } },
            {eItemContainerType.Ring, new ItemContainerLayout { ResourceName = ResourcePaths.RingAmuletPlaceholder, BaseFrame = 1 } },
            {eItemContainerType.Glove, new ItemContainerLayout { ResourceName = ResourcePaths.HelmGlovePlaceholder } },
            {eItemContainerType.Boots, new ItemContainerLayout { ResourceName = ResourcePaths.BootsPlaceholder } },
        }.ToImmutableDictionary();
    }

}
