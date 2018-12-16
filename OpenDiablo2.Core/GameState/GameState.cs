/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Services;
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
        private readonly ISoundProvider soundProvider;
        private readonly IMPQProvider mpqProvider;
        private readonly Func<IMapRenderer> getMapEngine;
        private readonly Func<eSessionType, ISessionManager> getSessionManager;
        private readonly Func<string, IRandomizedMapGenerator> getRandomizedMapGenerator;

        readonly private IMouseCursor originalMouseCursor;


        private float animationTime;
        private List<IMapInfo> mapInfo;
        private readonly List<MapCellInfo> mapDataLookup;
        private ISessionManager sessionManager;

        public int Act { get; private set; }
        public string MapName { get; private set; }
        public Palette CurrentPalette => paletteProvider.PaletteTable[$"ACT{Act}"];
        public List<PlayerInfo> PlayerInfos { get; private set; }
        public eDifficulty Difficulty { get; private set; }

        public int Seed { get; internal set; }
        public ItemInstance SelectedItem { get; internal set; }
        public object ThreadLocker { get; } = new object();
        public int CameraOffset { get; set; } = 0;

        IEnumerable<PlayerInfo> IGameState.PlayerInfos => PlayerInfos;

        const double Deg2Rad = Math.PI / 180.0;

        public GameState(
            ISceneManager sceneManager,
            IResourceManager resourceManager,
            IPaletteProvider paletteProvider,
            IEngineDataManager engineDataManager,
            IRenderWindow renderWindow,
            ISoundProvider soundProvider,
            IMPQProvider mpqProvider,
            Func<IMapRenderer> getMapEngine,
            Func<eSessionType, ISessionManager> getSessionManager,
            Func<string, IRandomizedMapGenerator> getRandomizedMapGenerator
            )
        {
            this.sceneManager = sceneManager;
            this.resourceManager = resourceManager;
            this.paletteProvider = paletteProvider;
            this.getMapEngine = getMapEngine;
            this.getSessionManager = getSessionManager;
            this.engineDataManager = engineDataManager;
            this.renderWindow = renderWindow;
            this.soundProvider = soundProvider;
            this.mpqProvider = mpqProvider;
            this.getRandomizedMapGenerator = getRandomizedMapGenerator;

            originalMouseCursor = renderWindow.MouseCursor;
            PlayerInfos = new List<PlayerInfo>();
            mapDataLookup = new List<MapCellInfo>();
        }

        public void Initialize(string characterName, eHero hero, eSessionType sessionType, eDifficulty difficulty)
        {
            sessionManager = getSessionManager(sessionType);
            sessionManager.Initialize();

            sessionManager.OnSetSeed += OnSetSeedEvent;
            sessionManager.OnLocatePlayers += OnLocatePlayers;
            sessionManager.OnPlayerInfo += OnPlayerInfo;
            sessionManager.OnFocusOnPlayer += OnFocusOnPlayer;

            Difficulty = difficulty;

            mapInfo = new List<IMapInfo>();

            sceneManager.ChangeScene(eSceneType.Game);
            sessionManager.JoinGame(characterName, hero);
        }

        private void OnFocusOnPlayer(int clientHash, Guid playerId)
            => getMapEngine().FocusedPlayerId = playerId;

        private void OnPlayerInfo(int clientHash, IEnumerable<PlayerInfo> playerInfo)
            => PlayerInfos = playerInfo.ToList();

        private void OnLocatePlayers(int clientHash, IEnumerable<LocationDetails> playerLocationDetails)
        {
            foreach (var player in PlayerInfos)
            {
                
                var details = playerLocationDetails.FirstOrDefault(x => x.UID == player.UID);

                if (details == null)
                    continue;

                details.CopyMobLocationDetailsTo(player);
            }
        }

        private void OnSetSeedEvent(int clientHash, int seed)
        {
            log.Info($"Setting seed to {seed}");
            this.Seed = seed;
            new MapGenerator(this).Generate();
        }

        public int HasMap(int cellX, int cellY)
            => mapInfo.Count(z => (cellX >= z.TileLocation.Left) && (cellX < z.TileLocation.Right)
            && (cellY >= z.TileLocation.Top) && (cellY < z.TileLocation.Bottom));

        public IEnumerable<Size> GetMapSizes(int cellX, int cellY)
            => mapInfo
                .Where(z => (cellX >= z.TileLocation.Left) && (cellX < z.TileLocation.Right) && (cellY >= z.TileLocation.Top) && (cellY < z.TileLocation.Bottom))
                .Select(x => x.TileLocation.Size);

        public void RemoveEverythingAt(int cellX, int cellY)
            => mapInfo.RemoveAll(z => (cellX >= z.TileLocation.Left) && (cellX < z.TileLocation.Right) && (cellY >= z.TileLocation.Top) && (cellY < z.TileLocation.Bottom));

        public IMapInfo GetSubMapInfo(int levelPresetId, int levelTypeId, IMapInfo primaryMap, Point origin, int subTile = -1)
        {
            var levelPreset = engineDataManager.LevelPresets.First(x => x.Def == levelPresetId);
            var levelType = engineDataManager.LevelTypes.First(x => x.Id == levelTypeId);

            // Some maps have variations, so lets pick a random one
            var mapNames = new List<string>();
            if (levelPreset.File1 != "0") mapNames.Add(levelPreset.File1);
            if (levelPreset.File2 != "0") mapNames.Add(levelPreset.File2);
            if (levelPreset.File3 != "0") mapNames.Add(levelPreset.File3);
            if (levelPreset.File4 != "0") mapNames.Add(levelPreset.File4);
            if (levelPreset.File5 != "0") mapNames.Add(levelPreset.File5);
            if (levelPreset.File6 != "0") mapNames.Add(levelPreset.File6);


            var random = new Random(Seed + origin.X + origin.Y);
            var mapName = "data\\global\\tiles\\" + mapNames[subTile == -1 ? random.Next(mapNames.Count) : subTile].Replace("/", "\\");
            var fileData = resourceManager.GetMPQDS1(mapName, levelPreset, levelType);

            var result = new SubMapInfo
            {
                FileData = fileData,
                PrimaryMap = primaryMap,
                CellInfo = new Dictionary<eRenderCellType, MapCellInfo[]>(),
                TileLocation = new Rectangle(origin, new Size(fileData.Width - 1, fileData.Height - 1))
            };

            return result;
        }

        public void AddMap(IMapInfo map) => mapInfo.Add(map);

        public IMapInfo InsertSubMap(int levelPresetId, int levelTypeId, Point origin, IMapInfo primaryMap, int subTile = -1)
        {
            var result = GetSubMapInfo(levelPresetId, levelTypeId, primaryMap, origin, subTile);
            mapInfo.Add(result);

            return result;
        }

        public IMapInfo InsertMap(eLevelId levelId, IMapInfo parentMap = null) => InsertMap((int)levelId, new Point(0, 0), parentMap);

        public IMapInfo InsertMap(int levelId, Point origin, IMapInfo parentMap = null)
        {

            var levelDetails = engineDataManager.Levels.First(x => x.Id == levelId);

            if (levelDetails.LevelPreset == null)
            {
                // There is no preset level, so we must generate one
                var generator = getRandomizedMapGenerator(levelDetails.LevelName);

                if (generator == null)
                    throw new OpenDiablo2Exception($"Could not locate a map generator for '{levelDetails.LevelName}'.");

                generator.Generate(parentMap, origin);

                // There is no core map so we cannot return a value here. If anyone actually uses
                // this value on a generated map they are making a terrible mistake anyway...
                return null;
            }

            // Some maps have variations, so lets pick a random one
            var mapNames = new List<string>();
            if (levelDetails.LevelPreset.File1 != "0") mapNames.Add(levelDetails.LevelPreset.File1);
            if (levelDetails.LevelPreset.File2 != "0") mapNames.Add(levelDetails.LevelPreset.File2);
            if (levelDetails.LevelPreset.File3 != "0") mapNames.Add(levelDetails.LevelPreset.File3);
            if (levelDetails.LevelPreset.File4 != "0") mapNames.Add(levelDetails.LevelPreset.File4);
            if (levelDetails.LevelPreset.File5 != "0") mapNames.Add(levelDetails.LevelPreset.File5);
            if (levelDetails.LevelPreset.File6 != "0") mapNames.Add(levelDetails.LevelPreset.File6);


            var random = new Random(Seed);
            // -------------------------------------------------------------------------------------
            // var mapName = "data\\global\\tiles\\" + mapNames[random.Next(mapNames.Count)].Replace("/", "\\");
            // -------------------------------------------------------------------------------------
            // TODO: ***TEMP FOR TESTING
            var mapName = "data\\global\\tiles\\" + mapNames[0].Replace("/", "\\");
            // -------------------------------------------------------------------------------------
            MapName = levelDetails.LevelPreset.Name;
            Act = levelDetails.LevelType.Act;

            var fileData = resourceManager.GetMPQDS1(mapName, levelDetails.LevelPreset, levelDetails.LevelType);

            var result = new MapInfo
            {
                LevelId = levelId,
                FileData = fileData,
                CellInfo = new Dictionary<eRenderCellType, MapCellInfo[]>(),
                TileLocation = new Rectangle(origin, new Size(fileData.Width - 1, fileData.Height - 1))
            };

            mapInfo.Add(result);

            // Only change music if loading a 'core' map
            if (Enum.IsDefined(typeof(eLevelId), levelId))
            {
                soundProvider.StopSong();
                soundProvider.LoadSong(mpqProvider.GetStream(ResourcePaths.GetMusicPathForLevel((eLevelId)levelId)));
                soundProvider.PlaySong();
            }


            return result;
        }

        public IEnumerable<MapCellInfo> GetMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType)
        {
            var map = GetMap(ref cellX, ref cellY);

            if (map == null)
                return Enumerable.Empty<MapCellInfo>();

            if (cellY >= map.FileData.Height || cellX >= map.FileData.Width)
                return new List<MapCellInfo>(); // Temporary code

            var idx = cellX + (cellY * map.FileData.Width);

            switch (renderCellType)
            {
                case eRenderCellType.Floor:
                    return map.FileData.FloorLayers
                        .Select(floorLayer => GetMapCellInfo(map, cellX, cellY, floorLayer.Props[idx], eRenderCellType.Floor, 0))
                        .Where(x => x != null);
                case eRenderCellType.Shadow:
                    return map.FileData.ShadowLayers
                        .Select(shadowLayer => GetMapCellInfo(map, cellX, cellY, shadowLayer.Props[idx], eRenderCellType.Shadow, 0))
                        .Where(x => x != null);
                case eRenderCellType.WallNormal:
                case eRenderCellType.WallLower:
                case eRenderCellType.Roof:
                    return map.FileData.WallLayers
                        .Select(wallLayer => GetMapCellInfo(map, cellX, cellY, wallLayer.Props[idx], renderCellType, wallLayer.Orientations[idx].Orientation1))
                        .Where(x => x != null);

                default:
                    throw new OpenDiablo2Exception("Unknown render cell type!");
            }
        }

        private IMapInfo GetMap(ref int cellX, ref int cellY)
        {
            var p = new Point(cellX, cellY);

            IMapInfo mi = null;
            for (var i = mapInfo.Count - 1; i >= 0; i--)
            {
                if (mapInfo[i].TileLocation.Contains(p))
                {
                    mi = mapInfo[i];
                    break;
                }

            }
            if (mi == null)
                return null;

            cellX -= mi.TileLocation.X;
            cellY -= mi.TileLocation.Y;
            return mi;
        }

        public void UpdateMapCellInfo(int cellX, int cellY, eRenderCellType renderCellType, IEnumerable<MapCellInfo> mapCellInfo)
        {
            throw new NotImplementedException();
        }

        public void SelectItem(ItemInstance item)
        {
            if (item == null)
            {
                renderWindow.MouseCursor = this.originalMouseCursor;
            }
            else
            {
                var cursorSprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(item.Item.InvFile), Palettes.Units);
                renderWindow.MouseCursor = renderWindow.LoadCursor(cursorSprite, 0, new Point(cursorSprite.FrameSize.Width / 2, cursorSprite.FrameSize.Height / 2));
            }

            this.SelectedItem = item;
        }

        private MapCellInfo GetMapCellInfo(IMapInfo map, int cellX, int cellY, MPQDS1TileProps props, eRenderCellType cellType, byte orientation)
        {
            if (props.Prop1 == 0)
                return null;

            if (!map.CellInfo.ContainsKey(cellType))
            {
                map.CellInfo[cellType] = new MapCellInfo[map.FileData.Width * map.FileData.Height];
            }

            var cellInfo = map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)];

            if (cellInfo != null && (cellInfo.Ignore || !cellInfo.Tile.Animated))
                return cellInfo.Ignore ? null : cellInfo;

            var mainIndex = (props.Prop3 >> 4) + ((props.Prop4 & 0x03) << 4);
            var subIndex = props.Prop2;

            if (orientation == 0)
            {
                // Floor or Shadow
                if (cellType != eRenderCellType.Floor && cellType != eRenderCellType.Shadow)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }
            else if (orientation == 10 || orientation == 11)
            {
                if (cellType != eRenderCellType.WallNormal)
                {
                    // Special tile
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }
            else if (orientation == 14)
            {
                // Walls (objects?) with precedent shadows
                if (cellType != eRenderCellType.WallNormal)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }
            else if (orientation < 15)
            {
                // Upper walls
                if (cellType != eRenderCellType.WallNormal)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }
            else if (orientation == 15)
            {
                // Roof
                if (cellType != eRenderCellType.Roof)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }
            else
            {
                // Lower Walls
                if (cellType != eRenderCellType.WallLower)
                {
                    map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                    return null;
                }
            }

            IEnumerable<MPQDT1Tile> tiles = Enumerable.Empty<MPQDT1Tile>();
            tiles = map.FileData.LookupTable
                .Where(x => x.MainIndex == mainIndex && x.SubIndex == subIndex && x.Orientation == orientation)
                .Select(x => x.TileRef);

            if (tiles == null || !tiles.Any())
            {
                log.Error($"Could not find tile [{mainIndex}:{subIndex}:{orientation}]!");
                map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = new MapCellInfo { Ignore = true };
                return null;
            }

            MPQDT1Tile tile = null;
            if (tiles.First().Animated)
            {
#if DEBUG
                if (!tiles.All(x => x.Animated))
                    throw new OpenDiablo2Exception("Some tiles are animated and some aren't...");
#endif
                var frameIndex = (int)Math.Floor(tiles.Count() * animationTime);
                tile = tiles.ElementAt(frameIndex);
            }
            else
            {
                if (tiles.Any())
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
                        throw new OpenDiablo2Exception("Why are we randomly finding an animated tile? Something's wrong here.");
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


            var mapCellInfo = mapDataLookup.FirstOrDefault(x => x.Tile.Id == tile.Id);
            if (mapCellInfo == null)
            {
                mapCellInfo = renderWindow.CacheMapCell(tile, cellType);
                mapDataLookup.Add(mapCellInfo);
            }

            map.CellInfo[cellType][cellX + (cellY * map.FileData.Width)] = mapCellInfo;
            return mapCellInfo;
        }

        public void Update(long ms)
        {
            animationTime += ms / 1000f;
            animationTime -= (float)Math.Truncate(animationTime);
            var seconds = ms / 1000f;

            foreach (var player in PlayerInfos)
                UpdatePlayer(player, seconds);
        }

        private void UpdatePlayer(PlayerInfo player, float seconds)
        {
            (new MobMovementService(player)).CalculateMovement(seconds);

        }

        public void Dispose()
        {
            sessionManager?.Dispose();
        }
    }
}
