using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{
    [Scene("Game")]
    public sealed class Game : IScene
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;
        private readonly IMapEngine mapEngine;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IGameState gameState;
        private readonly ISessionManager sessionManager;
        private readonly IKeyboardInfoProvider keyboardInfoProvider;
        private readonly IGameHUD gameHUD;

        //private ISprite[] testSprite;
        
        private eMovementType lastMovementType = eMovementType.Stopped;
        private byte lastDirection = 255;

        const double Rad2Deg = 180.0 / Math.PI;

        public Game(
            IRenderWindow renderWindow,
            IResourceManager resourceManager,
            IMapEngine mapEngine,
            IGameState gameState,
            IMouseInfoProvider mouseInfoProvider,
            IKeyboardInfoProvider keyboardInfoProvider,
            IItemManager itemManager,
            ISessionManager sessionManager,
            IGameHUD gameHUD
        )
        {
            this.renderWindow = renderWindow;
            this.resourceManager = resourceManager;
            this.mapEngine = mapEngine;
            this.gameState = gameState;
            this.mouseInfoProvider = mouseInfoProvider;
            this.keyboardInfoProvider = keyboardInfoProvider;
            this.sessionManager = sessionManager;
            this.gameHUD = gameHUD;

            /*var item = itemManager.getItem("hdm");
            var cursorsprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(item.InvFile), Palettes.Units);
            
            renderWindow.MouseCursor = renderWindow.LoadCursor(cursorsprite, 0, new Point(cursorsprite.FrameSize.Width/2, cursorsprite.FrameSize.Height / 2));*/
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
            var seconds = ms / 1000f;

            HandleMovement();

            mapEngine.Update(ms);
            gameHUD.Update();
        }

        private void HandleMovement()
        {
            // todo; if clicked on hud, then we don't move. But when clicked on map and move cursor over hud, then it's fine
            if (gameHUD.IsMouseOver())
                return;

            var mx = (mouseInfoProvider.MouseX - 400) - gameState.CameraOffset;
            var my = (mouseInfoProvider.MouseY - 300);

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
