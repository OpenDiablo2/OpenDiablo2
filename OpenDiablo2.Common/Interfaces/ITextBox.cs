using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ITextBox : IDisposable
    {
        string Text { get; set; }
        Point Location { get; set; }
        void Render();
        void Update(long ms);
    }
}
