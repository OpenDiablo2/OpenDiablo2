using System;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces.Drawing
{
    public interface ICharacterRenderer : IDisposable
    {
        Guid UID { get; set; }
        PlayerLocationDetails LocationDetails { get; set; }
        eHero Hero { get; set; }
        eWeaponClass WeaponClass { get; set; }
        eArmorType ArmorType { get; set; }
        eMobMode MobMode { get; set; }
        string ShieldCode { get; set; }

        void Update(long ms);
        void Render(int pixelOffsetX, int pixelOffsetY);
        void ResetAnimationData();
    }
}
