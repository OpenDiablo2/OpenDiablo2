/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System;

namespace OpenDiablo2.Scenes
{
    [Scene(eSceneType.Game)]
    public sealed class Game : IScene
    {
        private readonly IRenderWindow renderWindow;
        private readonly IMapEngine mapEngine;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IGameState gameState;
        private readonly ISessionManager sessionManager;
        private readonly IGameHUD gameHUD;

        private eMovementType lastMovementType = eMovementType.Stopped;
        private byte lastDirection = 255;

        const double Rad2Deg = 180.0 / Math.PI;

        public Game(
            IRenderWindow renderWindow,
            IMapEngine mapEngine,
            IGameState gameState,
            IMouseInfoProvider mouseInfoProvider,
            IItemManager itemManager,
            ISessionManager sessionManager,
            ISoundProvider soundProvider,
            IMPQProvider mpqProvider,
            IGameHUD gameHUD
        )
        {
            this.renderWindow = renderWindow;
            this.mapEngine = mapEngine;
            this.gameState = gameState;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sessionManager = sessionManager;
            this.gameHUD = gameHUD;

            //var item = itemManager.getItem("hdm");
        }

        public void Render()
        {
            // TODO: Maybe show some sort of connecting/loading message?
            if (mapEngine.FocusedPlayerId == 0)
                return;

            mapEngine.Render();
            gameHUD.Render();
        }

        public void Update(long ms)
        {
            HandleMovement();

            mapEngine.Update(ms);
            gameHUD.Update();
        }

        private void HandleMovement()
        {
            // todo; if clicked on hud, then we don't move. But when clicked on map and move cursor over hud, then it's fine
            if (gameHUD.IsMouseOver())
                return;

            var mx = mouseInfoProvider.MouseX - 400 - gameState.CameraOffset;
            var my = mouseInfoProvider.MouseY - 300;

            var tx = (mx / 60f + my / 40f) / 2f;
            var ty = (my / 40f - (mx / 60f)) / 2f;
            var cursorDirection = (int)Math.Round(Math.Atan2(ty, tx) * Rad2Deg);
            if (cursorDirection < 0)
                cursorDirection += 360;
            var actualDirection = (byte)(cursorDirection / 22);
            if (actualDirection >= 16)
                actualDirection -= 16;

            if (mouseInfoProvider.LeftMouseDown && (lastMovementType == eMovementType.Stopped || lastDirection != actualDirection))
            {
                lastDirection = actualDirection;
                lastMovementType = gameHUD.IsRunningEnabled ? eMovementType.Running : eMovementType.Walking;
                sessionManager.MoveRequest(actualDirection, lastMovementType);
            }
            else if (!mouseInfoProvider.LeftMouseDown && lastMovementType != eMovementType.Stopped)
            {
                lastDirection = actualDirection;
                lastMovementType = eMovementType.Stopped;
                sessionManager.MoveRequest(actualDirection, lastMovementType);
            }
        }

        public void Dispose()
        {

        }
    }
}
