using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Drawing;
using System.Linq;

namespace OpenDiablo2.MapGenerators
{
    // TODO: Different difficulties have different sizes. We need to read this from levels.txt
    [RandomizedMap("Blood Moor")]
    public sealed class BloodMoor : IRandomizedMapGenerator
    {
        private readonly IGameState gameState;

        private readonly LevelDetail levelDetail;
        private readonly eDifficulty difficulty;

        public BloodMoor(IGameState gameState, IEngineDataManager dataManager)
        {
            this.gameState = gameState;
            this.difficulty = gameState.Difficulty;

            levelDetail = dataManager.Levels.First(x => x.LevelName == "Blood Moor");
        }

        public void Generate(IMapInfo parentMap, Point location)
        {

            // Generate the area inside of the borders
            GenerateGeneralContents(location, parentMap);
        }

        private void GenerateGeneralContents(Point location, IMapInfo parentMap)
        {
            bool northExit = location.Y < 0;
            bool southExit = location.Y > 0;
            bool westExit = location.X < 0;
            bool eastExit = location.X > 0;


            if (northExit)
            {
                GenerateSouthernTownEntrance(location, parentMap);
                GenerateEasternRiver(location, parentMap);
                GenerateWesternFence(location, parentMap);
                GenerateNorthernFence(location, parentMap);
            }
            FillCenterArea(location, parentMap);
            GeneratePointsOfInterests(location, parentMap);

            #region old code
            //for (var y = 0; y < levelDetail.SizeY; y += 8)
            //{
            //    for (var x = 0; x < levelDetail.SizeX; x += 8)
            //    {
            //        var px = location.X + x;
            //        var py = location.Y + y;

            //        if ((x == 0) && (y == 0)) // North West
            //        {
            //            gameState.InsertSubMap((int)eWildBorder.NorthWest, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if ((x == levelDetail.SizeX - 8) && (y == 0)) // North East
            //        {
            //            gameState.InsertSubMap((int)eWildBorder.NorthEast, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if ((x == levelDetail.SizeX - 8) && (y == levelDetail.SizeY - 8)) // South East
            //        {
            //            if (northExit)
            //            {
            //                gameState.InsertSubMap((int)eWildBorder.RiverUpper, 2, new Point(location.X + x, location.Y + y), parentMap, 0);
            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.SouthEast, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if ((x == 0) && (y == levelDetail.SizeY - 8)) // South West
            //        {
            //            if (northExit)
            //            {
            //                gameState.InsertSubMap((int)eWildBorder.West, 2, new Point(px, py), parentMap, 0);
            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.SouthWest, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if ((x == 0) && ((y % 8) == 0)) // West
            //        {
            //            if (westExit)
            //            {
            //                gameState.InsertSubMap((int)eWildBorder.RiverUpper, 2, new Point(px, py), parentMap, 3);
            //            }
            //            else if (eastExit)
            //            {
            //                // TODO: Transition to town
            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.West, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if ((x == levelDetail.SizeX - 8) && ((y % 8) == 0)) // East
            //        {
            //            if (westExit)
            //            {
            //                // TODO: Transition to town
            //            }
            //            if (northExit || eastExit)
            //            {
            //                gameState.InsertSubMap((int)eWildBorder.RiverUpper, 2, new Point(px, py), parentMap, 3);
            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.East, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if (((x % 8) == 0) && (y == 0)) // North
            //        {
            //            if (southExit)
            //            {

            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.North, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else if (((x % 8) == 0) && (y == (levelDetail.SizeY - 8))) // South
            //        {
            //            if (northExit)
            //            {
            //                //var tileIdx = 31; // 8x8 fill
            //                //gameState.InsertSubMap(tileIdx, new Point(bloodMooreRect.Left + x, bloodMooreRect.Top + y), townMap);
            //            }
            //            else gameState.InsertSubMap((int)eWildBorder.South, 2, new Point(px, py), parentMap, 0);
            //        }
            //        else
            //        {

            //            //if (((x % 8) == 0) && ((y % 8)) == 0)
            //            //{
            //            //    var tileIdx = 31; // 8x8 fill
            //            //    gameState.InsertSubMap(tileIdx, 2, new Point(px, py), parentMap, 0);
            //            //}

            //        }
            //    }
            //}
            #endregion
        }

        private void GenerateNorthernFence(Point location, IMapInfo parentMap)
        {

            for (var x = location.X + 8; x < location.X + 64; x += 8)
            {
                gameState.InsertSubMap((int)eWildBorder.North, 2, new Point(x, location.Y), parentMap, 0);
            }
        }

        private void GenerateWesternFence(Point location, IMapInfo parentMap)
        {
            gameState.InsertSubMap((int)eWildBorder.NorthWest, 2, new Point(location.X, location.Y), parentMap, 0);

            for (var y = location.Y + 8; y < location.Y + 72; y += 8)
            {
                gameState.InsertSubMap((int)eWildBorder.West, 2, new Point(location.X, y), parentMap, 0);
            }
        }

        private void GenerateEasternRiver(Point location, IMapInfo parentMap)
        {
            for (var y = location.Y + 8; y < location.Y + 72; y += 8)
            {
                gameState.InsertSubMap(26 /* River Upper */, 2, new Point(location.X + 62, y), parentMap, 3);
                gameState.InsertSubMap(27 /* River Lower */, 2, new Point(location.X + 70, y), parentMap, 3);
            }

            // River near town
            gameState.InsertSubMap(26 /* River Upper */, 2, new Point(location.X + 62, location.Y + 72), parentMap, 2);
            gameState.InsertSubMap(27 /* River Lower */, 2, new Point(location.X + 70, location.Y + 72), parentMap, 2);

            // River near the back
            gameState.InsertSubMap(26 /* River Upper */, 2, new Point(location.X + 62, location.Y), parentMap, 1);
            gameState.InsertSubMap(27 /* River Lower */, 2, new Point(location.X + 70, location.Y), parentMap, 1);
        }

        private void GenerateSouthernTownEntrance(Point location, IMapInfo parentMap)
        {
            var random = new Random(gameState.Seed + location.X + location.Y);
            for (var x = location.X; x < location.X + 64; x += 8)
            {
                if (x == location.X) // Bottom left of the map
                    gameState.InsertSubMap((int)eWildBorder.ClosedBoxBottomLeft, 2, new Point(x, location.Y + 72), parentMap);
                else if (x == location.X + 48) // Town exit
                    gameState.InsertSubMap((int)eWildBorder.South, 2, new Point(x, location.Y + 72), parentMap, 3);
                else // Everything else
                    gameState.InsertSubMap((int)eWildBorder.South, 2, new Point(x, location.Y + 72), parentMap, 0);
            }
        }

        private void GeneratePointsOfInterests(Point location, IMapInfo parentMap)
        {
            var random = new Random(gameState.Seed + location.X + location.Y);

            // Generate the cave
            while (true)
            {
                var rx = random.Next(8, levelDetail.Difficulties[difficulty].SizeX - 16);
                var ry = random.Next(8, levelDetail.Difficulties[difficulty].SizeY - 16);
                rx -= (rx % 8);
                ry -= (ry % 8);
                var caveX = rx + location.X;
                var caveY = ry + location.Y;
                /*
                // Don't generate a camp on something that's already generated
                var loc = gameState.GetMapSizes(caveX, caveY).First();

                if (loc.Width != 8 || loc.Height != 8)
                    continue;

                gameState.RemoveEverythingAt(caveX, caveY);
                */
                gameState.InsertSubMap(52/*cave entrance*/, 2, new Point(caveX, caveY), parentMap);
                break;
            }

            // Generate camps
            var campsToGenerate = 3;
            while (campsToGenerate > 0)
            {
                var rx = random.Next(8, levelDetail.Difficulties[difficulty].SizeX - 16);
                var ry = random.Next(8, levelDetail.Difficulties[difficulty].SizeY - 16);
                rx -= (rx % 8);
                ry -= (ry % 8);
                var campX = rx + location.X;
                var campY = ry + location.Y;
                /*
                // Don't generate a camp on something that's already generated
                var loc = gameState.GetMapSizes(campX, campY).First();

                if (loc.Width != 8 || loc.Height != 8)
                    continue;

                gameState.RemoveEverythingAt(campX, campY);
                */

                if (gameState.HasMap(campX, campY) > 1)
                    continue;

                gameState.InsertSubMap(random.Next(42, 43 + 1), 2, new Point(campX, campY), parentMap);
                campsToGenerate--;
            }

        }

        static int[] tileFillIds = new int[] { 29, 30, 31, 38, 39, 40, 41 };
        private void FillCenterArea(Point location, IMapInfo parentMap)
        {
            var rightEdge = levelDetail.Difficulties[difficulty].SizeX - 8;
            var bottomEdge = levelDetail.Difficulties[difficulty].SizeY - 8;

            for (var y = 8; y < bottomEdge; y += 8)
            {
                for (var x = 8; x < rightEdge; x += 8)
                {
                    var px = location.X + x;
                    var py = location.Y + y;
                  
                    // If this space is already filled, move on
                    if (gameState.HasMap(px, py) > 0)
                        continue;

                    // Generate filler
                    var random = new Random(gameState.Seed + location.X + location.Y + x + y);

                    while (true)
                    {
                        var tileIdx = tileFillIds[random.Next(0, tileFillIds.Count())];
                        var info = gameState.GetSubMapInfo(tileIdx, 2, parentMap, new Point(location.X + x, location.Y + y));

                        // Make sure this tile actually fits here
                        if (info.TileLocation.Right > (location.X + rightEdge) || info.TileLocation.Bottom > (location.Y + bottomEdge))
                            continue;

                        if (info.TileLocation.Width == 16 && gameState.HasMap(px + 8, py) > 0)
                            continue;

                        if (info.TileLocation.Height == 16 && gameState.HasMap(px, py + 8) > 0)
                            continue;

                        // We found a tile that fits, so add it
                        gameState.AddMap(info);
                        break;
                    }
                    

                }
            }
        }
    }
}
