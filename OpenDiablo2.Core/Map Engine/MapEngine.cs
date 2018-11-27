using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapEngine : IMapEngine
    {
        private readonly IGameState gameState;
        private readonly IRenderWindow renderWindow;
        private readonly IResourceManager resourceManager;

        // TODO: Break this out further so we can support multiple maps
        private Dictionary<Guid, List<MapCellInfo>> mapDataLookup = new Dictionary<Guid, List<MapCellInfo>>();

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
            for (int y = 0; y < gameState.MapData.Height; y++)
            {
                for (int x = 0; x < gameState.MapData.Width; x++)
                {
                    var visualX = ((x - y) * (cellSizeX / 2)) - cOffX;
                    var visualY = ((x + y) * (cellSizeY / 2)) - cOffY;

                    
                    DrawFloor(x, y, visualX, visualY);
                    DrawWall(x, y, visualX, visualY, false);
                    DrawWall(x, y, visualX, visualY, true);
                    DrawRoof(x, y, visualX, visualY);
                   // //DrawShadow(x, y, visualX, visualY);
                }
            }

        }

        private void DrawRoof(int x, int y, int visualX, int visualY)
        {
            var cx = ((x - y) * 80) - cOffX;
            var cy = ((x + y) * 40) - cOffY;


            foreach (var wallLayer in gameState.MapData.WallLayers)
            {
                var wall = wallLayer.Props[x + (y * gameState.MapData.Width)];
                var orientation = wallLayer.Orientations[x + (y * gameState.MapData.Width)].Orientation1;

                if (orientation != 15) // Only 15 (roof)
                    return;

                if (wall.Prop1 == 0)
                    continue;

                if ((wall.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                        return;
                }

                var sub_index = wall.Prop2;
                var main_index = (wall.Prop3 >> 4) + ((wall.Prop4 & 0x03) << 4);

                var lt = gameState.MapData.LookupTable.First(z => z.MainIndex == main_index && z.SubIndex == sub_index && z.Orientation == orientation);
                renderWindow.DrawMapCell(x, y, cx, cy - lt.TileRef.RoofHeight, gameState.MapData, main_index, sub_index, gameState.CurrentPalette, orientation);
            }
        }

        private void DrawShadow(int x, int y, int visualX, int visualY)
        {
            
        }

        private void DrawFloor(int x, int y, int visualX, int visualY)
        {
            if (visualX < -160 || visualX > 800 || visualY < -120 || visualY > 650)
                return;

            var cx = ((x - y) * 80) - cOffX;
            var cy = ((x + y) * 40) - cOffY;

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


                renderWindow.DrawMapCell(x, y, cx, cy, gameState.MapData, main_index, sub_index, gameState.CurrentPalette);
            }
        }

        private void DrawWall(int x, int y, int visualX, int visualY, bool upper)
        {
            var cx = ((x - y) * 80) - cOffX;
            var cy = ((x + y) * 40) - cOffY;


            foreach (var wallLayer in gameState.MapData.WallLayers)
            {
                var wall = wallLayer.Props[x + (y * gameState.MapData.Width)];
                var orientation = wallLayer.Orientations[x + (y * gameState.MapData.Width)].Orientation1;


                if (wall.Prop1 == 0)
                    continue;

                if (upper && orientation <= 15)
                    return;


                if (orientation == 10 || orientation == 11)
                    return; // TODO: Support special walls
                
                if ((wall.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                        return;
                }

                var sub_index = wall.Prop2;
                var main_index = (wall.Prop3 >> 4) + ((wall.Prop4 & 0x03) << 4);

                var lt = gameState.MapData.LookupTable.First(z => z.MainIndex == main_index && z.SubIndex == sub_index && z.Orientation == orientation);
                renderWindow.DrawMapCell(x, y, cx, cy + 80, gameState.MapData, main_index, sub_index, gameState.CurrentPalette, orientation);
            }
        }

        public void Update(long ms)
        {

        }

        private void PurgeAllMapData()
        {

        }

        public MapCellInfo GetMapCellInfo(Guid mapId, Guid tileId)
        {
            if (!mapDataLookup.ContainsKey(mapId))
                return null;

            return  mapDataLookup[mapId].FirstOrDefault(x => x.TileId == tileId);
        }

        public void SetMapCellInfo(Guid mapId, MapCellInfo cellInfo)
        {
            if (!mapDataLookup.ContainsKey(mapId))
                mapDataLookup[mapId] = new List<MapCellInfo>();

            mapDataLookup[mapId].Add(cellInfo);
        }
    }
}
