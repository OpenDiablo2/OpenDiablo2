using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
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
                var idx = x + (y * gameState.MapData.Width);

                var cellInfo = GetMapCellInfo(
                    gameState.MapData.Id, cx, cy, wallLayer.Props[idx],
                    eRenderCellType.Roof,
                    wallLayer.Orientations[idx]);

                if (cellInfo == null)
                    return;

                renderWindow.DrawMapCell(cellInfo, cx, cy);
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

                var cellInfo = GetMapCellInfo(gameState.MapData.Id, cx, cy, floorLayer.Props[idx], eRenderCellType.Floor);

                if (cellInfo == null)
                    return;

                renderWindow.DrawMapCell(cellInfo, cx, cy);
            }
        }

        private void DrawWall(int x, int y, int visualX, int visualY, bool upper)
        {
            var cx = ((x - y) * 80) - cOffX;
            var cy = ((x + y) * 40) - cOffY;


            foreach (var wallLayer in gameState.MapData.WallLayers)
            {
                var idx = x + (y * gameState.MapData.Width);

                var cellInfo = GetMapCellInfo(
                    gameState.MapData.Id, cx, cy, wallLayer.Props[idx], 
                    upper ? eRenderCellType.WallUpper : eRenderCellType.WallLower,
                    wallLayer.Orientations[idx]);

                if (cellInfo == null)
                    return;

                renderWindow.DrawMapCell(cellInfo, cx, cy);
            }
        }

        public void Update(long ms)
        {

        }

        private void PurgeAllMapData()
        {

        }

        private MapCellInfo GetMapCellInfo(
            Guid mapId,
            int cellX,
            int cellY,
            MPQDS1TileProps props,
            eRenderCellType cellType,
            MPQDS1WallOrientationTileProps wallOrientations = null
            )
        {
            var sub_index = props.Prop2;
            var main_index = (props.Prop3 >> 4) + ((props.Prop4 & 0x03) << 4);
            var orientation = 0;

            if (cellType == eRenderCellType.Floor)
            {
                // Floors can't have rotations, should we blow up here?
                if (props.Prop1 == 0)
                    return null;
            }

            if (cellType == eRenderCellType.Roof)
            {
                if (orientation != 15) // Only 15 (roof)
                    return null;

                if (props.Prop1 == 0)
                    return null;

                if ((props.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                        return null;
                }
            }
            if (cellType == eRenderCellType.WallUpper || cellType == eRenderCellType.WallLower)
            {
                orientation = wallOrientations.Orientation1;


                if (props.Prop1 == 0)
                    return null;

                // < 15 shouldn't happen for upper wall types, should we even check for this?
                if (cellType == eRenderCellType.WallUpper && orientation <= 15)
                    return null;

                // TODO: Support special walls
                if (orientation == 10 || orientation == 11)
                    return null;

                // This is also a thing apparently
                if ((props.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                        return null;
                }

            }

            int frame = 0;
            var tiles = gameState.MapData.LookupTable
                .Where(x => x.MainIndex == main_index && x.SubIndex == sub_index && x.Orientation == orientation)
                .Select(x => x.TileRef);

            if (!tiles.Any())
                throw new ApplicationException("Invalid tile id found!");

            MPQDT1Tile tile = null;
            if (tiles.First().Animated)
            {
#if DEBUG
                if (!tiles.All(x => x.Animated))
                    throw new ApplicationException("Some tiles are animated and some aren't...");

                // TODO: Animated tiles
#endif
            }
            else
            {
                if (tiles.Count() > 0)
                {
                    var totalRarity = tiles.Sum(q => q.RarityOrFrameIndex);
                    var random = new Random(gameState.Seed + cellX + (gameState.MapData.Width * cellY));
                    var x = random.Next(totalRarity);
                    var z = 0;
                    foreach (var t in tiles)
                    {
                        z += t.RarityOrFrameIndex;
                        if (x <= z)
                        {
                            tile = t;
                            break;
                        }
                    }

                    if (tile.Animated)
                        throw new ApplicationException("Why are we randomly finding an animated tile? Something's wrong here.");
                }
                else tile = tiles.First();
            }

            // This WILL happen to you
            if (tile.Width == 0 || tile.Height == 0)
                return null;

            if (!mapDataLookup.ContainsKey(mapId))
                mapDataLookup[mapId] = new List<MapCellInfo>();

            var result = mapDataLookup[mapId].FirstOrDefault(x => x.TileId == tile.Id && x.AnimationId == frame);
            if (result != null)
                return result;

            var mapCellInfo = renderWindow.CacheMapCell(tile);
            mapDataLookup[mapId].Add(mapCellInfo);

            return mapCellInfo;
        }


    }
}
