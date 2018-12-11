using System.Collections.Generic;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IItemManager
    {
        Item getItem(string code);
        ItemInstance getItemInstance(string code);
    }
}
