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
        private readonly IEngineDataManager engineDataManager;
        private readonly Func<IMapEngine> getMapEngine;

        public MPQDS1 MapData { get; private set; }
        public int Act { get; private set; }
        public string MapName { get; private set; }
        public Palette CurrentPalette => paletteProvider.PaletteTable[$"ACT{Act}"];

        public int Seed { get; internal set; }

        public GameState(
            ISceneManager sceneManager,
            IResourceManager resourceManager,
            IPaletteProvider paletteProvider,
            IEngineDataManager engineDataManager,
            Func<IMapEngine> getMapEngine
            )
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.getMapEngine = getMapEngine;
            this.engineDataManager = engineDataManager;

            
        }

        public void Initialize(string characterName, eHero hero)
        {
            var random = new Random();
            Seed = random.Next();

            sceneManager.ChangeScene("Game");
            ChangeMap(eLevelId.Act2_Town);
        }

        public void ChangeMap(eLevelId levelId)
        {
            var level = engineDataManager.LevelPresets.First(x => x.LevelId == (int)levelId);
            var levelDetails = engineDataManager.LevelDetails.First(x => x.Id == level.LevelId);
            var levelType = engineDataManager.LevelTypes.First(x => x.Id == levelDetails.LevelType);


            var mapNames = new List<string>();
            if (level.File1 != "0") mapNames.Add(level.File1);
            if (level.File2 != "0") mapNames.Add(level.File2);
            if (level.File3 != "0") mapNames.Add(level.File3);
            if (level.File4 != "0") mapNames.Add(level.File4);
            if (level.File5 != "0") mapNames.Add(level.File5);
            if (level.File6 != "0") mapNames.Add(level.File6);

            var random = new Random(Seed);
            var mapName = "data\\global\\tiles\\" + mapNames[random.Next(mapNames.Count())].Replace("/", "\\");
            MapName = level.Name;
            Act = levelType.Act;
            MapData = resourceManager.GetMPQDS1(mapName, level, levelDetails, levelType);

            getMapEngine().NotifyMapChanged();
        }
    }
}
