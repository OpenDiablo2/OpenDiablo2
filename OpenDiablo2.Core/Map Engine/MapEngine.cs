﻿using System;
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
                cOffX = (int)((cameraLocation.X - cameraLocation.Y) * (cellSizeX / 2));
                cOffY = (int)((cameraLocation.X + cameraLocation.Y) * (cellSizeY / 2));
            }
        }

        private ISprite loadingSprite;
        private int cOffX, cOffY;
        //private ISprite[] tempMapCell;

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
            CameraLocation = new PointF(gameState.MapData.Width / 2, gameState.MapData.Height / 2);
        }

        private void LoadNewMapData()
        {
            /*
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
            */
            //CameraLocation = new Point(((gameState.MapData.Width * cellSizeX) / 2) - 400, ((gameState.MapData.Height * cellSizeY) / 2) - 300);
        }

        public void Render()
        {
            for (int y = 0; y < gameState.MapData.Width; y++)
                for (int x = 0; x < gameState.MapData.Height; x++)
                {

                    var visualX = ((x - y) * (cellSizeX / 2)) - cOffX;
                    var visualY = ((x + y) * (cellSizeY / 2)) - cOffY;

                    if (visualX < -160 || visualX > 800 || visualY < -80 || visualY > 600)
                        continue;

                    RenderFloorCell(x, y, ((x - y) * 80) - cOffX, ((x + y) * 40) - cOffY);
                }
        }

        private void RenderFloorCell(int x, int y, int xp, int yp)
        {
            if (x < 0 || y < 0 || x >= gameState.MapData.Width || y >= gameState.MapData.Height)
                return;


            renderWindow.DrawMapCell(x, y, xp, yp, gameState.MapData);
        }

        public void Update(long ms)
        {

        }

        private void PurgeAllMapData()
        {

        }
    }
}
