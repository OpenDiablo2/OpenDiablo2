using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{
    [Scene("Game")]
    public sealed class Game : IScene
    {
        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;
        private readonly IMapEngine mapEngine;
        private readonly IGameState gameState;
        private readonly IKeyboardInfoProvider keyboardInfoProvider;

        //private ISprite[] testSprite;

        private ISprite panelSprite, healthManaSprite, gameGlobeOverlapSprite;

        public Game(
            IRenderWindow renderWindow, 
            IResourceManager resourceManager, 
            IMapEngine mapEngine, 
            IGameState gameState,
            IKeyboardInfoProvider keyboardInfoProvider
        )
        {
            this.renderWindow = renderWindow;
            this.resourceManager = resourceManager;
            this.mapEngine = mapEngine;
            this.gameState = gameState;
            this.keyboardInfoProvider = keyboardInfoProvider;

            panelSprite = renderWindow.LoadSprite(ResourcePaths.GamePanels, Palettes.Act1);
            healthManaSprite = renderWindow.LoadSprite(ResourcePaths.HealthMana, Palettes.Act1);
            gameGlobeOverlapSprite = renderWindow.LoadSprite(ResourcePaths.GameGlobeOverlap, Palettes.Act1);
        }

        public void Render()
        {
            /*
            if (gameState.MapDirty)
                RedrawMap();

            for (int i = 0; i < gameState.MapData.Width * gameState.MapData.Height; i++)
                renderWindow.Draw(testSprite[i]);
                */
            mapEngine.Render();
            DrawPanel();

        }

        private void DrawPanel()
        {
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

        }

        public void Update(long ms)
        {
            var seconds = (float)ms / 1000f;
            var xMod = 0f;
            var yMod = 0f;

            if (keyboardInfoProvider.KeyIsPressed(80 /*left*/))
            {
                xMod = -8f * seconds;
                yMod = 8f * seconds;
            }

            if (keyboardInfoProvider.KeyIsPressed(79 /*right*/))
            {
                xMod = 8f * seconds;
                yMod = -8f * seconds;
            }

            if (keyboardInfoProvider.KeyIsPressed(81 /*down*/))
            {
                yMod = 10f * seconds;
                xMod = 10f * seconds;
            }

            if (keyboardInfoProvider.KeyIsPressed(82 /*up*/))
            {
                yMod = -10f * seconds;
                xMod = -10f * seconds;
            }

            if (xMod != 0f || yMod != 0f)
                mapEngine.CameraLocation = new PointF(mapEngine.CameraLocation.X + xMod, mapEngine.CameraLocation.Y + yMod);

            mapEngine.Update(ms);
        }

        public void Dispose()
        {

        }
        /*
        private void RedrawMap()
        {
            gameState.MapDirty = false;

            testSprite = new ISprite[gameState.MapData.Width * gameState.MapData.Height];
            var idx = 0;
            for (int y = 0; y < gameState.MapData.Height; y++)
            {
                for (int x = 0; x < gameState.MapData.Width; x++)
                {
                    testSprite[idx] = renderWindow.GenerateMapCell(gameState.MapData, x, y, eRenderCellType.Floor, gameState.CurrentPalette);
                    testSprite[idx].Location = new Point(((x - y) * 80) - 2900, ((x + y) * 40) - 1900);
                    idx++;
                }
            }

        }*/
    }
}
