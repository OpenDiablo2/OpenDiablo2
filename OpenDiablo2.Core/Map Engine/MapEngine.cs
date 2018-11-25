using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapEngine : IMapEngine
    {
        private readonly IGameState gameState;
        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;

        private PointF cameraLocation = new PointF();
        public PointF CameraLocation
        {
            get => cameraLocation;
            set
            {
                if (cameraLocation == value)
                    return;

                cameraLocation = value;
                /*
                cellOffsetX = CameraLocation.X / cellSizeX;
                cellOffsetY = CameraLocation.Y / cellSizeY;
                pixelOffsetX = CameraLocation.X % cellSizeX;
                pixelOffsetY = CameraLocation.Y % cellSizeY;
                */
            }
        }

        private ISprite loadingSprite;
        private ISprite[] tempMapCell;

        private int cellOffsetX, cellOffsetY, pixelOffsetX, pixelOffsetY;

        private const int
            cellSizeX = 160,
            cellSizeY = 80,
            renderCellsX = (800 / cellSizeX) + 1,
            renderCellsY = (600 / cellSizeY) + 1;

        public MapEngine(
            IGameState gameState,
            IRenderWindow renderWindow,
            IResourceManager resourceManager
            )
        {
            this.gameState = gameState;
            this.renderWindow = renderWindow;
            this.resourceManager = resourceManager;

            loadingSprite = renderWindow.LoadSprite(ResourcePaths.LoadingScreen, Palettes.Loading, new Point(300, 400));
        }

        public void NotifyMapChanged()
        {
            PurgeAllMapData();
            LoadNewMapData();
        }

        private void LoadNewMapData()
        {
            var cellsToLoad = gameState.MapData.Width * gameState.MapData.Height;
            tempMapCell = new ISprite[cellsToLoad];

            for (var cell = 0; cell < cellsToLoad; cell++)
            {
                renderWindow.Clear();
                loadingSprite.Frame = (int)(loadingSprite.TotalFrames * ((float)cell / (float)cellsToLoad));
                renderWindow.Draw(loadingSprite);
                renderWindow.Sync();

                tempMapCell[cell] = renderWindow.GenerateMapCell(
                    gameState.MapData, 
                    cell % gameState.MapData.Width, 
                    cell / gameState.MapData.Width, 
                    Common.Enums.eRenderCellType.Floor, 
                    gameState.CurrentPalette
                );
            }
            
            //CameraLocation = new Point(((gameState.MapData.Width * cellSizeX) / 2) - 400, ((gameState.MapData.Height * cellSizeY) / 2) - 300);
        }

        public void Render()
        {
            var cOffX = (int)((cameraLocation.X - cameraLocation.Y) * (cellSizeX / 2));
            var cOffY = (int)((cameraLocation.X + cameraLocation.Y) * (cellSizeY / 2));

            for (int y = 0; y < gameState.MapData.Width; y++)
                for (int x = 0; x < gameState.MapData.Height; x++)
                {
                    RenderFloorCell(
                        (x + cellOffsetX),
                        (y + cellOffsetY),
                        ((x - y) * 80) - cOffX,
                        ((x + y) * 40) - cOffY
                    );
                }
        }

        public void RenderFloorCell(int x, int y, int xp, int yp)
        {
            if (x < 0 || y < 0 || x >= gameState.MapData.Width || y >= gameState.MapData.Height)
                return;

            renderWindow.Draw(tempMapCell[x + (y * gameState.MapData.Width)], new Point(xp, yp));
        }

        public void Update(long ms)
        {

        }

        private void PurgeAllMapData()
        {

        }
    }
}
