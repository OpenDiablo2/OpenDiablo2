using System;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapGenerator
    {
        private readonly IGameState gameState;

        public MapGenerator(IGameState gameState)
        {
            this.gameState = gameState;
        }

        public void Generate()
        {
            GenerateAct1Town();
        }

        private void GenerateAct1Town()
        {
            var townMap = gameState.LoadMap(eLevelId.Act1_Town1, new Point(0, 0));
            //var wildBorder = 5; // (4-15)

            Rectangle bloodMooreRect;

            // 32-37 is grassy field?
            bool westExit = false;
            bool eastExit = false;
            bool southExit = false;
            bool northExit = false;

            if (townMap.FileData.MapFile.Contains("S1"))
            {
                var defId = 3; // Act 1 - Town 1 Transition S
                var borderMap = gameState.LoadSubMap(defId, new Point(0, townMap.FileData.Height - 2), townMap);
                borderMap.PrimaryMap = townMap;

                var wilderness = gameState.LoadSubMap(defId, new Point(26, townMap.FileData.Height + borderMap.FileData.Height - 2), townMap);
                wilderness.PrimaryMap = townMap;

                bloodMooreRect = new Rectangle(-40, townMap.FileData.Height + borderMap.FileData.Height + 4, 120, 80);
                southExit = true;
            }
            else if (townMap.FileData.MapFile.Contains("E1"))
            {
                var defId = 2; // Act 1 - Town 1 Transition E
                var borderMap = gameState.LoadSubMap(defId, new Point(townMap.FileData.Width - 2, 0), townMap);
                borderMap.PrimaryMap = townMap;

                bloodMooreRect = new Rectangle(townMap.FileData.Width + borderMap.FileData.Width - 4, -40, 80, 120);
                eastExit = true;
            }
            else if (townMap.FileData.MapFile.Contains("W1"))
            {
                // West
                bloodMooreRect = new Rectangle(-120, 0, 120, townMap.FileData.Height);
                westExit = true;
            }
            else // North
            {
                bloodMooreRect = new Rectangle(0, -120, townMap.FileData.Width - 8, 120);
                northExit = true;
            }

            // Generate the Blood Moore?
            for (var y = 0; y < bloodMooreRect.Height; y += 8)
            {
                for (var x = 0; x < bloodMooreRect.Width; x += 8)
                {
                    var px = bloodMooreRect.Left + x;
                    var py = bloodMooreRect.Top + y;

                    if ((x == 0) && (y == 0)) // North West
                    {
                        gameState.LoadSubMap((int)eWildBorder.NorthWest, new Point(px, py), townMap, 0);
                    }
                    else if ((x == bloodMooreRect.Width - 9) && (y == 0)) // North East
                    {
                        gameState.LoadSubMap((int)eWildBorder.NorthEast, new Point(px, py), townMap, 0);
                    }
                    else if ((x == bloodMooreRect.Width - 9) && (y == bloodMooreRect.Height - 9)) // South East
                    {
                        if (northExit)
                        {
                            gameState.LoadSubMap((int)eWildBorder.RiverUpper, new Point(bloodMooreRect.Left + x, bloodMooreRect.Top + y), townMap, 0);
                        }
                        else gameState.LoadSubMap((int)eWildBorder.SouthEast, new Point(px, py), townMap, 0);
                    }
                    else if ((x == 0) && (y == bloodMooreRect.Height - 9)) // South West
                    {
                        if (northExit)
                        {
                            gameState.LoadSubMap((int)eWildBorder.West, new Point(px, py), townMap, 0);
                        }
                        else gameState.LoadSubMap((int)eWildBorder.SouthWest, new Point(px, py), townMap, 0);
                    }
                    else if ((x == 0) && ((y % 8) == 0)) // West
                    {
                        if (westExit)
                        {
                            gameState.LoadSubMap((int)eWildBorder.RiverUpper, new Point(px, py), townMap, 3);
                        }
                        else if (eastExit)
                        {
                            // TODO: Transition to town
                        }
                        else gameState.LoadSubMap((int)eWildBorder.West, new Point(px, py), townMap, 0);
                    }
                    else if ((x == bloodMooreRect.Width - 9) && ((y % 8) == 0)) // East
                    {
                        if (westExit)
                        {
                            // TODO: Transition to town
                        }
                        if (northExit || eastExit)
                        {
                            gameState.LoadSubMap((int)eWildBorder.RiverUpper, new Point(px, py), townMap, 3);
                        }
                        else gameState.LoadSubMap((int)eWildBorder.East, new Point(px, py), townMap, 0);
                    }
                    else if (((x % 8) == 0) && (y == 0)) // North
                    {
                        if (southExit)
                        {

                        }
                        else gameState.LoadSubMap((int)eWildBorder.North, new Point(px, py), townMap, 0);
                    }
                    else if (((x % 8) == 0) && (y == (bloodMooreRect.Height - 9))) // South
                    {
                        if (northExit)
                        {
                            //var tileIdx = 31; // 8x8 fill
                            //gameState.LoadSubMap(tileIdx, new Point(bloodMooreRect.Left + x, bloodMooreRect.Top + y), townMap);
                        }
                        else gameState.LoadSubMap((int)eWildBorder.South, new Point(px, py), townMap, 0);
                    }
                    else
                    {
                        if (((x % 8) == 0) && ((y % 8)) == 0)
                        {
                            var tileIdx = 31; // 8x8 fill
                            gameState.LoadSubMap(tileIdx, new Point(bloodMooreRect.Left + x, bloodMooreRect.Top + y), townMap, 0);
                        }

                    }
                }
            }



        }

    }
}
