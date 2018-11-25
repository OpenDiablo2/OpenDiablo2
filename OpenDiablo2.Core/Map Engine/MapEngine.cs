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

        }

        public void Render()
        {
            // Lower Walls, Floors, and Shadows

            // Shadows of objects

            // Objects with OrderFlag = 1

            // Upper Walls and objects with ORderFlag = 0 or 2

            // Roofs


            for (int y = 0; y < gameState.MapData.Width; y++)
                for (int x = 0; x < gameState.MapData.Height; x++)
                {

                    var visualX = ((x - y) * (cellSizeX / 2)) - cOffX;
                    var visualY = ((x + y) * (cellSizeY / 2)) - cOffY;

                    if (visualX < -160 || visualX > 800 || visualY < -120 || visualY > 650)
                        continue;

                    // Render the floor
                    foreach (var floorLayer in gameState.MapData.FloorLayers)
                    {
                        var idx = x + (y * gameState.MapData.Width);
                        if (idx >= floorLayer.Props.Length)
                            break;
                        var floor = floorLayer.Props[idx];

                        if (floor.Prop1 == 0)
                            continue;

                        var sub_index = floor.Prop2;
                        var main_index = (floor.Prop3 >> 4) + ((floor.Prop4 & 0x03) << 4);


                        renderWindow.DrawMapCell(x, y, ((x - y) * 80) - cOffX, ((x + y) * 40) - cOffY, gameState.MapData, main_index, sub_index, gameState.CurrentPalette, null);
                    }

                }
            /*

            // Render the walls
            foreach (var wallLayer in gameState.MapData.WallLayers)
            {

                for (int y = 0; y < gameState.MapData.Width; y++)
                    for (int x = 0; x < gameState.MapData.Height; x++)
                    {

                        var visualX = ((x - y) * (cellSizeX / 2)) - cOffX;
                        var visualY = ((x + y) * (cellSizeY / 2)) - cOffY;

                        if (visualX < -160 || visualX > 800 || visualY < -120 || visualY > 650)
                            continue;

                        var idx = x + (y * gameState.MapData.Width);
                        if (idx >= wallLayer.Props.Length)
                            continue;
                        var wall = wallLayer.Props[idx];

                        if (wall.Prop1 == 0)
                            continue;

                        var sub_index = wall.Prop2;
                        var main_index = (wall.Prop3 >> 4) + ((wall.Prop4 & 0x03) << 4);

                        var orientation = wallLayer.Orientations[x + (y * gameState.MapData.Width)];
                        renderWindow.DrawMapCell(x, y, ((x - y) * 80) - cOffX, ((x + y) * 40) - cOffY + 80, gameState.MapData, main_index, sub_index, gameState.CurrentPalette, orientation);

                    }

            }
            */

        }

        public void Update(long ms)
        {

        }

        private void PurgeAllMapData()
        {

        }
    }
}
