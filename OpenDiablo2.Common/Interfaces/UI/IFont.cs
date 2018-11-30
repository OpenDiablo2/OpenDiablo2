using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IFont : IDisposable
    {
        Size CalculateSize(string text);
    }
}
