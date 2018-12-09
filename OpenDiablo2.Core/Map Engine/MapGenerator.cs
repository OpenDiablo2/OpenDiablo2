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
            //var wildBorder = 5; // (4-15)
            // TODO: Is there no data file that explicitly defines this??
            var townMap = gameState.LoadMap(eLevelId.Act1_Town1, new Point(0, 0));

            /*
            Rectangle bloodMooreRect;

            // 32-37 is grassy field?

            if (townMap.FileData.MapFile.Contains("S1"))
            {
                var defId = 3; // Act 1 - Town 1 Transition S
                var borderMap = gameState.LoadSubMap(defId, new Point(0, townMap.FileData.Height - 2));
                borderMap.PrimaryMap = townMap;

                var wilderness = gameState.LoadSubMap(wildBorder, new Point(26, townMap.FileData.Height + borderMap.FileData.Height - 2));
                wilderness.PrimaryMap = townMap;

                bloodMooreRect = new Rectangle(-40, townMap.FileData.Height + borderMap.FileData.Height, 120, 80);
            }
            else if (townMap.FileData.MapFile.Contains("E1"))
            {
                var defId = 2; // Act 1 - Town 1 Transition E
                var borderMap = gameState.LoadSubMap(defId, new Point(townMap.FileData.Width - 2, 0));
                borderMap.PrimaryMap = townMap;

                bloodMooreRect = new Rectangle(townMap.FileData.Width + borderMap.FileData.Width, -40, 80, 120);
            }
            else if (townMap.FileData.MapFile.Contains("W1"))
            {
                bloodMooreRect = new Rectangle(-120, 0, 120, townMap.FileData.Height);
            } else // North
            {
                bloodMooreRect = new Rectangle(0, -120, townMap.FileData.Width, 120);
            }

            // Generate the Blood Moore?
            for (var y = 0; y < bloodMooreRect.Height; y+= 8)
            {
                
                for (var x = 0; x < bloodMooreRect.Width; x += 8)
                {
                    var tileIdx = 35;
                    var mapTile = gameState.LoadSubMap(tileIdx, new Point(bloodMooreRect.Left + x, bloodMooreRect.Top + y));
                    mapTile.PrimaryMap = townMap;
                }
            }
            */

        }

      
    }
}
