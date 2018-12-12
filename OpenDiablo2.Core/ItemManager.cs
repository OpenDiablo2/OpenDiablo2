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

using System.Linq;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class ItemManager : IItemManager
    {
        private readonly IEngineDataManager engineDataManager;

        public ItemManager(IEngineDataManager engineDataManager)
        {
            this.engineDataManager = engineDataManager;
        }

        public Item getItem(string code)
        {
            Item item = engineDataManager.Items.FirstOrDefault(x => x.Code == code);

            return item;
        }

        public ItemInstance getItemInstance(string code)
        {
            return new ItemInstance(getItem(code));
        }

    }
}
