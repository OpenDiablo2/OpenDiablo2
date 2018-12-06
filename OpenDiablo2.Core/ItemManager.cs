using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class ItemManager : IItemManager
    {
        private IEngineDataManager engineDataManager;

        public ItemManager(IEngineDataManager engineDataManager)
        {
            this.engineDataManager = engineDataManager;
        }

        public Item getItem(string code)
        {
            Item item = engineDataManager.Items.Where(x => x.Code == code).FirstOrDefault();

            return item;
        }

    }
}
