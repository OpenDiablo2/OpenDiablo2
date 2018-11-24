using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.GameState_
{
    public sealed class GameState
    {
        private readonly ISceneManager sceneManager;

        public GameState(ISceneManager sceneManager)
        {
            this.sceneManager = sceneManager;
        }

        public void Initialize(string characterName, eHero hero)
        {

            sceneManager.ChangeScene("Game");
        }
    }
}
