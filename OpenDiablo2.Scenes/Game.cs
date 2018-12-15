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
using System.Linq;

namespace OpenDiablo2.Scenes
{
    [Scene(eSceneType.Game)]
    public sealed class Game : IScene
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IMapRenderer _mapRenderer;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IGameState gameState;
        private readonly ISessionManager sessionManager;
        private readonly IGameHUD gameHUD;

        private eMovementType lastMovementType = eMovementType.Stopped;
        private byte lastDirection = 255;
        private bool clickedOnHud = false;
        private float lastMoveSend = 0f;
        private float holdMoveTime = 0f;

        public Game(
            IRenderWindow renderWindow,
            IMapRenderer mapRenderer,
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
            this._mapRenderer = mapRenderer;
            this.gameState = gameState;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sessionManager = sessionManager;
            this.gameHUD = gameHUD;

            //var item = itemManager.getItem("hdm");
        }

        public void Render()
        {
            // TODO: Maybe show some sort of connecting/loading message?
            if (_mapRenderer.FocusedPlayerId == Guid.Empty)
                return;

            _mapRenderer.Render();
            gameHUD.Render();
        }

        public void Update(long ms)
        {
            HandleMovement(ms);

            _mapRenderer.Update(ms);
            gameHUD.Update();
        }
        
        private void HandleMovement(long ms)
        {
            if (mouseInfoProvider.ReserveMouse)
                return;

            if(gameHUD.IsMouseOver() && lastMovementType == eMovementType.Stopped)
                clickedOnHud = true;
            else if (!mouseInfoProvider.LeftMouseDown)
                clickedOnHud = false;

            if (clickedOnHud)
                return;

            /*
            var mx = mouseInfoProvider.MouseX - 400 - gameState.CameraOffset;
            var my = mouseInfoProvider.MouseY - 300;

            var tx = (mx / 60f + my / 40f) / 2f;
            var ty = (my / 40f - (mx / 60f)) / 2f;
            */

            if (mouseInfoProvider.LeftMouseDown)
            {
                lastMoveSend += (ms / 1000f);
                holdMoveTime += (ms / 1000f);
                if (lastMoveSend < .25f)
                    return;
                lastMoveSend = 0f;
                var selectedCell = _mapRenderer.GetCellPositionAt(mouseInfoProvider.MouseX, mouseInfoProvider.MouseY);
#if DEBUG
                log.Debug($"Move to cell: ({selectedCell.X}, {selectedCell.Y})");
#endif
                sessionManager.MoveRequest(selectedCell, gameHUD.IsRunningEnabled ? eMovementType.Running : eMovementType.Walking);
            }
            else
            {
                lastMoveSend = 1f; // Next click will always send a mouse move request (TODO: Should this be limited as well?)

                if (holdMoveTime > 0.2f)
                {
                    var player = gameState.PlayerInfos.First(x => x.UID == _mapRenderer.FocusedPlayerId);
                    sessionManager.MoveRequest(new System.Drawing.PointF(player.X, player.Y), eMovementType.Stopped);
                }
                holdMoveTime = 0f;
            }

            /*
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
            */
        }

        public void Dispose()
        {

        }
    }
}
