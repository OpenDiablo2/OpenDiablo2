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
    public sealed class GameState : IGameState
    {
        private readonly ISceneManager sceneManager;
        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;
        private readonly Func<IMapEngine> getMapEngine;

        public MPQDS1 MapData { get; private set; }
        public int Act { get; private set; }
        public string MapName { get; private set; }
        public Palette CurrentPalette => paletteProvider.PaletteTable[$"ACT{Act}"];

        public GameState(
            ISceneManager sceneManager,
            IResourceManager resourceManager,
            IPaletteProvider paletteProvider,
            Func<IMapEngine> getMapEngine
            )
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.getMapEngine = getMapEngine;
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
            MapData = resourceManager.GetMPQDS1(mapName, -1, act);
            getMapEngine().NotifyMapChanged();
        }
    }
}
