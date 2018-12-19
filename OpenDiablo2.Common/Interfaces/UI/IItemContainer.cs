using OpenDiablo2.Common.Models;
using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IItemContainer : IDisposable
    {
        ItemInstance ContainedItem { get; }
        Point Location { get; set; }
        string Slot { get; set; }

        void SetContainedItem(ItemInstance containedItem);
        void Render();
        void Update();
    }
}
