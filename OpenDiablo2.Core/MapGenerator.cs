using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
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
            var wildBorder = 5; // (4-15)
            // TODO: Is there no data file that explicitly defines this??
            var townMap = gameState.LoadMap(eLevelId.Act1_Town1, new Point(0, 0));
            if (townMap.FileData.MapFile.Contains("S1"))
            {
                var defId = 3; // Act 1 - Town 1 Transition S
                var borderMap = gameState.LoadSubMap(defId, new Point(0, townMap.FileData.Height));
                borderMap.PrimaryMap = townMap;

                var wilderness = gameState.LoadSubMap(wildBorder, new Point(26, townMap.FileData.Height + borderMap.FileData.Height));
                wilderness.PrimaryMap = townMap;
            }
            else if (townMap.FileData.MapFile.Contains("E1"))
            {
                var defId = 2; // Act 1 - Town 1 Transition E
                var borderMap = gameState.LoadSubMap(defId, new Point(townMap.FileData.Width, 0));
                borderMap.PrimaryMap = townMap;

                for (int i = 4; i <= 15; i++)
                {
                    var wilderness = gameState.LoadSubMap(i, new Point(townMap.FileData.Width + borderMap.FileData.Width + ((i-4) * 10), 26));
                    wilderness.PrimaryMap = townMap;
                }
            }

        }

      
    }
}
