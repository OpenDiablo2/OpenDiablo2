using System;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Interfaces.Drawing
{
    public interface ICharacterRenderer : IDisposable
    {
        Guid UID { get; set; }

        void Update(long ms);
        void Render(int pixelOffsetX, int pixelOffsetY);
        void ResetAnimationData();
        void ResetCache();
    }
}
