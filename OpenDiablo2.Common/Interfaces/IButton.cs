using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnActivateDelegate();
    public delegate void OnToggleDelegate(bool isToggled);

    public interface IButton : IDisposable
    {
        OnActivateDelegate OnActivate { get; set; }
        bool Enabled { get; set; }
        Point Location { get; set; }
        OnToggleDelegate OnToggle { get; set; }
        string Text { get; set; }

        void Update();
        void Render();
    }
}
