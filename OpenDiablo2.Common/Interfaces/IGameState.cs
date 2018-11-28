using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IGameState
    {
        MPQDS1 MapData { get; }
        int Act { get; }
        int Seed { get; }
        string MapName { get; }
        Palette CurrentPalette { get; }

        void Initialize(string text, eHero value);
    }
}
