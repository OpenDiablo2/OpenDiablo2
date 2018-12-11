using System;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

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
            var townMap = gameState.InsertMap((int)eLevelId.Act1_Town1, new Point(0, 0));
            Point bloodMoorOrigin;

            if (townMap.FileData.MapFile.Contains("S1"))
            {
                var defId = 3; // Act 1 - Town 1 Transition S
                var borderMap = gameState.InsertSubMap(defId, 1, new Point(0, townMap.FileData.Height - 2), townMap);

                gameState.InsertSubMap(defId, 1, new Point(26, townMap.FileData.Height + borderMap.FileData.Height - 2), townMap);

                bloodMoorOrigin = new Point(0, townMap.FileData.Height + borderMap.FileData.Height + 4);
            }
            else if (townMap.FileData.MapFile.Contains("E1"))
            {
                var defId = 2; // Act 1 - Town 1 Transition E
                var borderMap = gameState.InsertSubMap(defId, 1, new Point(townMap.FileData.Width - 2, 0), townMap);

                bloodMoorOrigin = new Point(townMap.FileData.Width + borderMap.FileData.Width - 4, 0);
            }
            else if (townMap.FileData.MapFile.Contains("W1"))
            {
                // West
                bloodMoorOrigin = new Point(-80, 0);
            }
            else // North
            {
                bloodMoorOrigin = new Point(-22, -80); // Align along the eastern edge
            }

            gameState.InsertMap(2 /*Wilderness 1*/, bloodMoorOrigin, townMap);
        }

    }
}
