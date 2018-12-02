using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.MessageBus;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Core.Map_Engine;

namespace OpenDiablo2.Core.GameState_
{
    public sealed class GameState : IGameState
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly ISceneManager sceneManager;
        private readonly IResourceManager resourceManager;
        private readonly IPaletteProvider paletteProvider;
        private readonly IEngineDataManager engineDataManager;
        private readonly IRenderWindow renderWindow;
        private readonly Func<IMapEngine> getMapEngine;
        private readonly Func<eSessionType, ISessionManager> getSessionManager;

        private float animationTime = 0f;
        private List<MapInfo> mapInfo;
        private List<MapCellInfo> mapDataLookup = new List<MapCellInfo>();
        private ISessionManager sessionManager;

        public int Act { get; private set; }
        public string MapName { get; private set; }
        public Palette CurrentPalette => paletteProvider.PaletteTable[$"ACT{Act}"];
        public IEnumerable<PlayerLocationDetails> PlayerLocationDetails { get; private set; } = new List<PlayerLocationDetails>();
        public IEnumerable<PlayerInfo> PlayerInfos { get; private set; } = new List<PlayerInfo>();

        public bool ShowInventoryPanel { get; set; } = false;
        public bool ShowCharacterPanel { get; set; } = false;

        readonly private IMouseCursor originalMouseCursor;

        public int Seed { get; internal set; }

        public Item SelectedItem { get; internal set; }
        public object ThreadLocker { get; } = new object();

        public GameState(
            ISceneManager sceneManager,
            IResourceManager resourceManager,
            IPaletteProvider paletteProvider,
            IEngineDataManager engineDataManager,
            IRenderWindow renderWindow,
            Func<IMapEngine> getMapEngine,
            Func<eSessionType, ISessionManager> getSessionManager
            )
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.getMapEngine = getMapEngine;
            this.getSessionManager = getSessionManager;
            this.engineDataManager = engineDataManager;
            this.renderWindow = renderWindow;

            this.originalMouseCursor = renderWindow.MouseCursor;

        }

        public void Initialize(string characterName, eHero hero, eSessionType sessionType)
        {
            sessionManager = getSessionManager(sessionType);
            sessionManager.Initialize();

            sessionManager.OnSetSeed += OnSetSeedEvent;
            sessionManager.OnLocatePlayers += OnLocatePlayers;
            sessionManager.OnPlayerInfo += OnPlayerInfo;
            sessionManager.OnFocusOnPlayer += OnFocusOnPlayer;

            mapInfo = new List<MapInfo>();
            sceneManager.ChangeScene("Game");

            sessionManager.JoinGame(characterName, hero);
        }

        private void OnFocusOnPlayer(int clientHash, int playerId)
            => getMapEngine().FocusedPlayerId = playerId;

        private void OnPlayerInfo(int clientHash, IEnumerable<PlayerInfo> playerInfo)
            => this.PlayerInfos = playerInfo;

        private void OnLocatePlayers(int clientHash, IEnumerable<PlayerLocationDetails> playerLocationDetails)
        {
            PlayerLocationDetails = playerLocationDetails;
        }

        private void OnSetSeedEvent(int clientHash, int seed)
        {
            log.Info($"Setting seed to {seed}");
            this.Seed = seed;
            (new MapGenerator(this)).Generate();
        }

        public MapInfo LoadSubMap(int levelDefId, Point origin)
        {
            var level = engineDataManager.LevelPresets.First(x => x.Def == levelDefId);
            var levelDetails = engineDataManager.LevelDetails.First(x => x.Id == level.LevelId);
            var levelType = engineDataManager.LevelTypes.First(x => x.Id == levelDetails.LevelType);

            // Some maps have variations, so lets pick a random one
            var mapNames = new List<string>();
            if (level.File1 != "0") mapNames.Add(level.File1);
            if (level.File2 != "0") mapNames.Add(level.File2);
            if (level.File3 != "0") mapNames.Add(level.File3);
            if (level.File4 != "0") mapNames.Add(level.File4);
            if (level.File5 != "0") mapNames.Add(level.File5);
            if (level.File6 != "0") mapNames.Add(level.File6);


            var random = new Random(Seed);
            var mapName = "data\\global\\tiles\\" + mapNames[random.Next(mapNames.Count())].Replace("/", "\\");
            var fileData = resourceManager.GetMPQDS1(mapName, level, levelDetails, levelType);

            var result = new MapInfo
            {
                LevelId = eLevelId.None,
                LevelPreset = level,
                LevelDetail = levelDetails,
                LevelType = levelType,
                FileData = fileData,
                CellInfo = new Dictionary<eRenderCellType, MapCellInfo[]>(),
                TileLocation = new Rectangle(origin, new Size(fileData.Width - 1, fileData.Height - 1))
            };

            mapInfo.Add(result);

            return result;
        }

        public MapInfo LoadMap(eLevelId levelId, Point origin)
        {
            var level = engineDataManager.LevelPresets.First(x => x.LevelId == (int)levelId);
            var levelDetails = engineDataManager.LevelDetails.First(x => x.Id == level.LevelId);
            var levelType = engineDataManager.LevelTypes.First(x => x.Id == levelDetails.LevelType);

            // Some maps have variations, so lets pick a random one
            var mapNames = new List<string>();
            if (level.File1 != "0") mapNames.Add(level.File1);
            if (level.File2 != "0") mapNames.Add(level.File2);
            if (level.File3 != "0") mapNames.Add(level.File3);
            if (level.File4 != "0") mapNames.Add(level.File4);
            if (level.File5 != "0") mapNames.Add(level.File5);
            if (level.File6 != "0") mapNames.Add(level.File6);


            var random = new Random(Seed);
            var mapName = "data\\global\\tiles\\" + mapNames[random.Next(mapNames.Count())].Replace("/", "\\");
            MapName = level.Name;
            Act = levelType.Act;

            var fileData = resourceManager.GetMPQDS1(mapName, level, levelDetails, levelType);

            var result = new MapInfo
            {
                LevelId = levelId,
                LevelPreset = level,
                LevelDetail = levelDetails,
                LevelType = levelType,
                FileData = fileData,
                CellInfo = new Dictionary<eRenderCellType, MapCellInfo[]>(),
                TileLocation = new Rectangle(origin, new Size(fileData.Width - 1, fileData.Height - 1))
            };

            mapInfo.Add(result);

            return result;
        }

        public IEnumerable<MapCellInfo> GetMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType)
        {
            var map = GetMap(ref cellX, ref cellY);

            if (map == null)
                return new List<MapCellInfo>();

            if (cellY >= map.FileData.Height || cellX >= map.FileData.Width)
                return new List<MapCellInfo>(); // Temporary code

            var idx = cellX + (cellY * map.FileData.Width);

            switch (renderCellType)
            {
                case eRenderCellType.Floor:
                    return map.FileData.FloorLayers
                        .Select(floorLayer => GetMapCellInfo(map, cellX, cellY, floorLayer.Props[idx], eRenderCellType.Floor))
                        .Where(x => x != null);

                case eRenderCellType.WallUpper:
                case eRenderCellType.WallLower:
                case eRenderCellType.Roof:
                    return map.FileData.WallLayers
                        .Select(wallLayer => GetMapCellInfo(map, cellX, cellY, wallLayer.Props[idx], renderCellType, wallLayer.Orientations[idx]))
                        .Where(x => x != null);

                default:
                    throw new ApplicationException("Unknown render cell type!");
            }
        }

        private MapInfo GetMap(ref int cellX, ref int cellY)
        {
            var x = cellX;
            var y = cellY;
            var map = mapInfo.FirstOrDefault(z => (x >= z.TileLocation.X) && (y >= z.TileLocation.Y)
                && (x < z.TileLocation.Right) && (y < z.TileLocation.Bottom));
            if (map == null)
            {
                return null;
            }

            cellX -= map.TileLocation.X;
            cellY -= map.TileLocation.Y;
            return map;
        }

        public void UpdateMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType, IEnumerable<MapCellInfo> mapCellInfo)
        {

        }

        public bool ToggleShowInventoryPanel()
        {
            ShowInventoryPanel = !ShowInventoryPanel;

            return ShowInventoryPanel;
        }

        public bool ToggleShowCharacterPanel()
        {
            ShowCharacterPanel = !ShowCharacterPanel;

            return ShowCharacterPanel;
        }

        public void SelectItem(Item item)
        {
            if(item == null)
            {
                renderWindow.MouseCursor = this.originalMouseCursor;
            } else {
                var cursorsprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(item.InvFile), Palettes.Units);

                renderWindow.MouseCursor = renderWindow.LoadCursor(cursorsprite, 0, new Point(cursorsprite.FrameSize.Width / 2, cursorsprite.FrameSize.Height / 2));
            }

            this.SelectedItem = item;
        }

        private MapCellInfo GetMapCellInfo(MapInfo map, int cellX, int cellY, MPQDS1TileProps props, eRenderCellType cellType, MPQDS1WallOrientationTileProps wallOrientations = null)
        {
            if (!map.CellInfo.ContainsKey(cellType))
            {
                map.CellInfo[cellType] = new MapCellInfo[map.FileData.Width * map.FileData.Height];
            }

            var cellInfo = map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)];
            if (cellInfo != null)
                return cellInfo.Ignore ? null : cellInfo;


            var sub_index = props.Prop2;
            var main_index = (props.Prop3 >> 4) + ((props.Prop4 & 0x03) << 4);
            var orientation = 0;

            if (cellType == eRenderCellType.Floor)
            {
                // Floors can't have rotations, should we blow up here?
                if (props.Prop1 == 0)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }

            if (cellType == eRenderCellType.Roof)
            {
                if (orientation != 15) // Only 15 (roof)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }

                if (props.Prop1 == 0)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }

                if ((props.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                    {
                        map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                        return null;
                    }
                }
            }
            if (cellType == eRenderCellType.WallUpper || cellType == eRenderCellType.WallLower)
            {
                orientation = wallOrientations.Orientation1;


                if (props.Prop1 == 0)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }

                // < 15 shouldn't happen for upper wall types, should we even check for this?
                if (cellType == eRenderCellType.WallUpper && orientation <= 15)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }

                // TODO: Support special walls
                if (orientation == 10 || orientation == 11)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }

                // This is also a thing apparently
                if ((props.Prop4 & 0x80) > 0)
                {
                    if (orientation != 10 && orientation != 11)
                    {
                        map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                        return null;
                    }
                }

            }

            int frame = 0;
            var tiles = (map.PrimaryMap ?? map).FileData.LookupTable
                .Where(x => x.MainIndex == main_index && x.SubIndex == sub_index && x.Orientation == orientation)
                .Select(x => x.TileRef);

            if (!tiles.Any())
            {
                map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                return null;
            }

            //throw new ApplicationException("Invalid tile id found!");

            MPQDT1Tile tile = null;
            if (tiles.First().Animated)
            {
#if DEBUG
                if (!tiles.All(x => x.Animated))
                    throw new ApplicationException("Some tiles are animated and some aren't...");
#endif
                var frameIndex = (int)Math.Floor(tiles.Count() * animationTime);
                tile = tiles.ElementAt(frameIndex);
            }
            else
            {
                if (tiles.Count() > 0)
                {
                    var totalRarity = tiles.Sum(q => q.RarityOrFrameIndex);
                    var random = new Random(Seed + cellX + (map.FileData.Width * cellY));
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
            {
                map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                return null;
            }


            if (tile.BlockDataLength == 0)
            {
                map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                return null; // Why is this a thing?
            }


            var mapCellInfo = mapDataLookup.FirstOrDefault(x => x.Tile.Id == tile.Id && x.AnimationId == frame);
            if (mapCellInfo == null)
            {
                mapCellInfo = renderWindow.CacheMapCell(tile);
                mapDataLookup.Add(mapCellInfo);

                switch (cellType)
                {
                    case eRenderCellType.WallUpper:
                    case eRenderCellType.WallLower:
                        mapCellInfo.OffY -= 80;
                        break;
                }

            }

            map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = mapCellInfo;
            return mapCellInfo;
        }

        public void Update(long ms)
        {
            animationTime += ((float)ms / 1000f);
            animationTime -= (float)Math.Truncate(animationTime);
        }

        public void Dispose()
        {
            sessionManager?.Dispose();
        }
    }
}
