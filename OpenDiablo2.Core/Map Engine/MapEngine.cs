using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapEngine : IMapEngine
    {
        private readonly IGameState gameState;
        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;

        public int FocusedMobId { get; set; } = -1;

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
            renderCellsX = (800 / cellSizeX),
            renderCellsY = (600 / cellSizeY);

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

        public void Render()
        {
            var cx = -(cameraLocation.X - Math.Truncate(cameraLocation.X));
            var cy = -(cameraLocation.Y - Math.Truncate(cameraLocation.Y));

            for (int ty = -5; ty <= 9; ty++)
            {
                for (int tx = -5; tx <= 9; tx++)
                {
                    var ax = tx + Math.Truncate(cameraLocation.X);
                    var ay = ty + Math.Truncate(cameraLocation.Y);

                    var px = (tx - ty) * (cellSizeX / 2);
                    var py = (tx + ty) * (cellSizeY / 2);

                    var ox = (cx - cy) * (cellSizeX / 2);
                    var oy = (cx + cy) * (cellSizeY / 2);


                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.Floor))
                        renderWindow.DrawMapCell(cellInfo, 320 + (int)px + (int)ox, 210 + (int)py + (int)oy);

                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.WallLower))
                        renderWindow.DrawMapCell(cellInfo, 320 + (int)px + (int)ox, 210 + (int)py + (int)oy);

                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.WallUpper))
                        renderWindow.DrawMapCell(cellInfo, 320 + (int)px + (int)ox, 210 + (int)py + (int)oy);

                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.Roof))
                        renderWindow.DrawMapCell(cellInfo, 320 + (int)px + (int)ox, 210 + (int)py + (int)oy);

                }
            }

        }


        public void Update(long ms)
        {
        }


    }
}
