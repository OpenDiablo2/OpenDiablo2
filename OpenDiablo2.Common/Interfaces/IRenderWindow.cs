using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IRenderWindow : IDisposable
    {
        bool IsRunning { get; }
        void Update();
        void Clear();
        void Sync();
        ISprite LoadSprite(ImageSet source);
        void Draw(ISprite sprite);
    }
}
