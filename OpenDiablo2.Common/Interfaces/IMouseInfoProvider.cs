using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMouseInfoProvider
    {
        int MouseX { get; }
        int MouseY { get; }
        bool LeftMouseDown { get; }
        bool RightMouseDown { get; }
    }
}
