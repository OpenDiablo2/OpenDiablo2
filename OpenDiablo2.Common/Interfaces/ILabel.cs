using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ILabel : IDisposable
    {
        string Text { get; set; }
        Point Position { get; set; }
        Size TextArea { get; set; }
        Color TextColor { get; set; }
    }
}
