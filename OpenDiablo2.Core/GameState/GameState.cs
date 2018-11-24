using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core.GameState_
{
    public sealed class GameState
    {
        private readonly ISceneManager sceneManager;
        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;

        public MPQDS1 MapData { get; set; }
        public bool MapDirty { get; set; }
        public int Act { get; private set; }
        public string MapName { get; set; }
        public Palette CurrentPalette => paletteProvider.PaletteTable[$"ACT{Act}"];

        public GameState(ISceneManager sceneManager, IResourceManager resourceManager, IPaletteProvider paletteProvider)
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
        }

        public void Initialize(string characterName, eHero hero)
        {
            sceneManager.ChangeScene("Game");
            ChangeMap(ResourcePaths.MapAct1TownE1, 1);
        }

        public void ChangeMap(string mapName, int act)
        {
            MapName = mapName;
            Act = act;
            MapDirty = true;
            MapData = resourceManager.GetMPQDS1(mapName, -1, act);
        }
    }
}
