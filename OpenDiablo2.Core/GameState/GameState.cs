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
        MPQDS1 mapData;

        public GameState(ISceneManager sceneManager, IResourceManager resourceManager)
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
        }

        public void Initialize(string characterName, eHero hero)
        {
            sceneManager.ChangeScene("Game");
            mapData = resourceManager.GetMPQDS1(ResourcePaths.MapAct1TownE1, -1, 1);
        }
    }
}
