using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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
