using OpenDiablo2.Common.Models;
using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IItemContainer : IDisposable
    {
        Item ContainedItem { get; set; }
        Point Location { get; set; }
        void Render();
        void Update();
    }
}
