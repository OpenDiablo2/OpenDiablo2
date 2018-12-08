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

        //private ISprite[] testSprite;

        private ISprite panelSprite, healthManaSprite, gameGlobeOverlapSprite;

        private readonly IMiniPanel minipanel;
        
        private IButton runButton, menuButton;
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
            Func<eButtonType, IButton> createButton,
            Func<IMiniPanel> createMiniPanel
        )
        {
            this.renderWindow = renderWindow;
            this.resourceManager = resourceManager;
            this.mapEngine = mapEngine;
            this.gameState = gameState;
            this.mouseInfoProvider = mouseInfoProvider;
            this.keyboardInfoProvider = keyboardInfoProvider;
            this.sessionManager = sessionManager;


            panelSprite = renderWindow.LoadSprite(ResourcePaths.GamePanels, Palettes.Act1);
            healthManaSprite = renderWindow.LoadSprite(ResourcePaths.HealthMana, Palettes.Act1);
            gameGlobeOverlapSprite = renderWindow.LoadSprite(ResourcePaths.GameGlobeOverlap, Palettes.Act1);

            minipanel = createMiniPanel();

            runButton = createButton(eButtonType.Run);
            runButton.Location = new Point(256, 570);
            runButton.OnToggle = OnRunToggle;

            // move to minipanel?
            menuButton = createButton(eButtonType.Menu);
            menuButton.Location = new Point(393, 561);
            menuButton.OnToggle = minipanel.OnMenuToggle;

            /*var item = itemManager.getItem("hdm");
            var cursorsprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(item.InvFile), Palettes.Units);
            
            renderWindow.MouseCursor = renderWindow.LoadCursor(cursorsprite, 0, new Point(cursorsprite.FrameSize.Width/2, cursorsprite.FrameSize.Height / 2));*/
        }

        private void OnRunToggle(bool isToggled)
        {
            log.Debug("Run Toggle: " + isToggled);
        }

        public void Render()
        {
            // TODO: Maybe show some sort of connecting/loading message?
            if (mapEngine.FocusedPlayerId == 0)
                return;

            mapEngine.Render();
            DrawPanel();

        }

        private void DrawPanel()
        {
            minipanel.Render();

            // Render the background bottom bar
            renderWindow.Draw(panelSprite, 0, new Point(0, 600));
            renderWindow.Draw(panelSprite, 1, new Point(166, 600));
            renderWindow.Draw(panelSprite, 2, new Point(294, 600));
            renderWindow.Draw(panelSprite, 3, new Point(422, 600));
            renderWindow.Draw(panelSprite, 4, new Point(550, 600));
            renderWindow.Draw(panelSprite, 5, new Point(685, 600));

            // Render the health bar
            renderWindow.Draw(healthManaSprite, 0, new Point(30, 587));
            renderWindow.Draw(gameGlobeOverlapSprite, 0, new Point(28, 595));

            // Render the mana bar
            renderWindow.Draw(healthManaSprite, 1, new Point(692, 588));
            renderWindow.Draw(gameGlobeOverlapSprite, 1, new Point(693, 591));

            
            
            runButton.Render();
            menuButton.Render();
        }

        public void Update(long ms)
        {
            var seconds = ms / 1000f;

            minipanel.Update();

            runButton.Update();
            menuButton.Update();

            HandleMovement();

            mapEngine.Update(ms);
        }

        private void HandleMovement()
        {
            if (mouseInfoProvider.MouseY > 530) // 550 is what it should be, but the minipanel check needs to happent oo
                return;

            if (minipanel.IsRightPanelVisible && mouseInfoProvider.MouseX >= 400)
                return;

            if (minipanel.IsLeftPanelVisible && mouseInfoProvider.MouseX < 400)
                return;


            // TODO: Filter movement for inventory panel
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
                lastMovementType = runButton.Toggled ? eMovementType.Running : eMovementType.Walking;
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
