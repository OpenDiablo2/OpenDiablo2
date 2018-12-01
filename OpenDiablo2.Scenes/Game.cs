﻿using System;
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

        private IMiniPanel minipanel;
        private ICharacterPanel characterpanel;
        private IInventoryPanel inventorypanel;

        private bool showMinipanel = false;
        private IButton runButton, menuButton;
        private eMovementType lastMovementType = eMovementType.Stopped;
        private int lastDirection = -1;

        const double Rad2Deg = 180.0 / Math.PI;
        const double Deg2Rad = Math.PI / 180.0;

        public Game(
            IRenderWindow renderWindow,
            IResourceManager resourceManager,
            IMapEngine mapEngine,
            IGameState gameState,
            IMouseInfoProvider mouseInfoProvider,
            IKeyboardInfoProvider keyboardInfoProvider,
            ISessionManager sessionManager,
            Func<eButtonType, IButton> createButton,
            Func<IMiniPanel> createMiniPanel,
            Func<ICharacterPanel> createCharacterPanel,
            Func<IInventoryPanel> createInventoryPanel
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
            // Maybe? Not sure. 
            // miniPanel.OnMenuActivate();

            characterpanel = createCharacterPanel();
            inventorypanel = createInventoryPanel();

            runButton = createButton(eButtonType.Run);
            runButton.Location = new Point(256, 570);
            runButton.OnToggle = OnRunToggle;

            menuButton = createButton(eButtonType.Menu);
            menuButton.Location = new Point(393, 561);
            menuButton.OnToggle = OnMenuToggle;
        }

        private void OnMenuToggle(bool isToggled)
        {
            this.showMinipanel = isToggled;
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
            if(gameState.ShowInventoryPanel)
            {
                inventorypanel.Render();
            }

            if (gameState.ShowCharacterPanel)
            {
                characterpanel.Render();
            }

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

            if(showMinipanel)
            {
                minipanel.Render();
            }

            
            
            runButton.Render();
            menuButton.Render();
        }

        public void Update(long ms)
        {
            var seconds = ms / 1000f;

            if(showMinipanel)
            {
                minipanel.Update();
            }

            characterpanel.Update();

            runButton.Update();
            menuButton.Update();

            HandleMovement();
            //var xMod = 0f;
            //var yMod = 0f;


            //if (keyboardInfoProvider.KeyIsPressed(80 /*left*/))
            //{
            //    xMod = -8f * seconds;
            //}

            //if (keyboardInfoProvider.KeyIsPressed(79 /*right*/))
            //{
            //    xMod = 8f * seconds;
            //}

            //if (keyboardInfoProvider.KeyIsPressed(81 /*down*/))
            //{
            //    yMod = 10f * seconds;
            //}

            //if (keyboardInfoProvider.KeyIsPressed(82 /*up*/))
            //{
            //    yMod = -10f * seconds;
            //}

            //if (xMod != 0f || yMod != 0f)
            //{
            //    xMod *= .5f;
            //    yMod *= .5f;
            //    mapEngine.CameraLocation = new PointF(mapEngine.CameraLocation.X + xMod, mapEngine.CameraLocation.Y + yMod);
            //}

            mapEngine.Update(ms);
        }

        private void HandleMovement()
        {
            if (mouseInfoProvider.MouseY > 530) // 550 is what it should be, but the minipanel check needs to happent oo
                return;

            if (gameState.ShowInventoryPanel && mouseInfoProvider.MouseX >= 400)
                return;

            if (gameState.ShowCharacterPanel && mouseInfoProvider.MouseX < 400)
                return;


            // TODO: Filter movement for inventory panel
            var xOffset = (gameState.ShowInventoryPanel ? -200 : 0) + (gameState.ShowCharacterPanel ? 200 : 0);

            var cursorDirection = (int)Math.Round(((Math.Atan2(300 - mouseInfoProvider.MouseY, mouseInfoProvider.MouseX - (400 + xOffset)) * Rad2Deg) + 180) / 32);

            if (mouseInfoProvider.LeftMouseDown && (lastMovementType == eMovementType.Stopped || lastDirection != cursorDirection))
            {
                lastDirection = cursorDirection;
                lastMovementType = runButton.Toggled ? eMovementType.Running : eMovementType.Walking;
                sessionManager.MoveRequest(cursorDirection, lastMovementType);
            }
            else if (!mouseInfoProvider.LeftMouseDown && lastMovementType != eMovementType.Stopped)
            {
                lastDirection = cursorDirection;
                lastMovementType = eMovementType.Stopped;
                sessionManager.MoveRequest(cursorDirection, lastMovementType);
            }
        }

        public void Dispose()
        {

        }
    }
}