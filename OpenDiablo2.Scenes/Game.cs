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
using OpenDiablo2.Core.GameState_;

namespace OpenDiablo2.Scenes
{
    [Scene("Game")]
    public sealed class Game : IScene
    {
        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;
        private GameState gameState;

        private ISprite testSprite;

        private ISprite panelSprite, healthManaSprite, gameGlobeOverlapSprite;

        public Game(IRenderWindow renderWindow, IResourceManager resourceManager, GameState gameState)
        {
            this.renderWindow = renderWindow;
            this.resourceManager = resourceManager;
            this.gameState = gameState;

            panelSprite = renderWindow.LoadSprite(ResourcePaths.GamePanels, Palettes.Act1);
            healthManaSprite = renderWindow.LoadSprite(ResourcePaths.HealthMana, Palettes.Act1);
            gameGlobeOverlapSprite = renderWindow.LoadSprite(ResourcePaths.GameGlobeOverlap, Palettes.Act1);
        }

        public void Render()
        {
            if (gameState.MapDirty)
                RedrawMap();

            renderWindow.Draw(testSprite);

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
            renderWindow.Draw(healthManaSprite, 0, new Point(28, 590));
            renderWindow.Draw(gameGlobeOverlapSprite, 0, new Point(28, 595));

            // Render the mana bar
            renderWindow.Draw(healthManaSprite, 1, new Point(691, 590));
            renderWindow.Draw(gameGlobeOverlapSprite, 1, new Point(693, 591));

        }

        public void Update(long ms)
        {
            
        }

        public void Dispose()
        {

        }

        private void RedrawMap()
        {
            gameState.MapDirty = false;

            var x = 0;
            var y = 0;
            testSprite = renderWindow.GenerateMapCell(gameState.MapData, 0, 0, eRenderCellType.Floor, gameState.CurrentPalette);
            testSprite.Location = new Point(((x * 80) - (y * 80)) + 100, ((x * 40) + (y * 40)) + 100);
        }
    }
}
