using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapEngine : IMapEngine
    {
        private readonly IGameState gameState;
        private readonly IRenderWindow renderWindow;

        private readonly List<ICharacterRenderer> characterRenderers = new List<ICharacterRenderer>();

        public int FocusedPlayerId { get; set; } = 0;

        private PointF cameraLocation = new PointF();
        public PointF CameraLocation
        {
            get => cameraLocation;
            set
            {
                if (cameraLocation == value)
                    return;

                cameraLocation = value;
            }
        }

        private const int
            cellSizeX = 160,
            cellSizeY = 80;

        public MapEngine(
            IGameState gameState,
            IRenderWindow renderWindow,
            ISessionManager sessionManager
            )
        {
            this.gameState = gameState;
            this.renderWindow = renderWindow;

            sessionManager.OnPlayerInfo += OnPlayerInfo;
            sessionManager.OnLocatePlayers += OnLocatePlayers;
        }

        private void OnLocatePlayers(int clientHash, IEnumerable<PlayerLocationDetails> playerLocationDetails)
        {
            foreach (var loc in playerLocationDetails)
            {
                var cr = characterRenderers.FirstOrDefault(x => x.LocationDetails.PlayerId == loc.PlayerId);
                var newDirection = loc.MovementDirection != cr.LocationDetails.MovementDirection;
                var stanceChanged = loc.MovementType != cr.LocationDetails.MovementType;
                cr.LocationDetails = loc;
                if (newDirection || stanceChanged)
                    cr.ResetAnimationData();
            }
        }

        private void OnPlayerInfo(int clientHash, IEnumerable<PlayerInfo> playerInfo)
        {
            // Remove character renderers for players that no longer exist...
            characterRenderers.RemoveAll(x => playerInfo.Any(z => z.UID == x.UID));

            // Update existing character renderers
            foreach (var cr in characterRenderers)
            {
                var info = playerInfo.FirstOrDefault(x => x.UID == cr.UID);
                if (info == null)
                    continue;

                // TODO: This shouldn't be necessary...
                cr.LocationDetails = info.LocationDetails;
                cr.MobMode = info.MobMode;
                cr.WeaponClass = info.WeaponClass;
                cr.Hero = info.Hero;
                cr.ArmorType = info.ArmorType;

            }

            // Add character renderers for characters that now exist
            foreach (var info in playerInfo.Where(x => !characterRenderers.Any(z => x.UID == z.UID)))
            {
                var cr = renderWindow.CreateCharacterRenderer();
                cr.UID = info.UID;
                cr.LocationDetails = info.LocationDetails;
                cr.MobMode = info.MobMode;
                cr.WeaponClass = info.WeaponClass;
                cr.Hero = info.Hero;
                cr.ArmorType = info.ArmorType;
                cr.ShieldCode = info.ShieldCode;
                cr.WeaponCode = info.WeaponCode;
                characterRenderers.Add(cr);
                cr.ResetAnimationData();
            }
        }


        const int skewX = 400;
        const int skewY = 300;

        public void Render()
        {
            var xOffset = gameState.CameraOffset;

            var cx = -(cameraLocation.X - Math.Truncate(cameraLocation.X));
            var cy = -(cameraLocation.Y - Math.Truncate(cameraLocation.Y));

            for (int ty = -7; ty <= 9; ty++)
            {
                for (int tx = -8; tx <= 8; tx++)
                {
                    var ax = tx + Math.Truncate(cameraLocation.X);
                    var ay = ty + Math.Truncate(cameraLocation.Y);

                    var px = (tx - ty) * (cellSizeX / 2);
                    var py = (tx + ty) * (cellSizeY / 2);

                    var ox = (cx - cy) * (cellSizeX / 2);
                    var oy = (cx + cy) * (cellSizeY / 2);


                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.WallLower))
                        renderWindow.DrawMapCell(cellInfo, skewX + px + (int)ox + xOffset, skewY + py + (int)oy + 80);

                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.Floor))
                        renderWindow.DrawMapCell(cellInfo, skewX + px + (int)ox + xOffset, skewY + py + (int)oy);


                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.WallNormal))
                        renderWindow.DrawMapCell(cellInfo, skewX + px + (int)ox + xOffset, skewY + py + (int)oy + 80);

                    foreach (var character in characterRenderers.Where(x => Math.Truncate(x.LocationDetails.PlayerX) == ax && Math.Truncate(x.LocationDetails.PlayerY) == ay))
                    {
                        var ptx = character.LocationDetails.PlayerX - Math.Truncate(cameraLocation.X);
                        var pty = character.LocationDetails.PlayerY - Math.Truncate(cameraLocation.Y);

                        var ppx = (ptx - pty) * (cellSizeX / 2);
                        var ppy = (ptx + pty) * (cellSizeY / 2);

                        character.Render(skewX + (int)ppx + (int)ox + xOffset, skewY + (int)ppy + (int)oy);
                    }

                    foreach (var cellInfo in gameState.GetMapCellInfo((int)ax, (int)ay, eRenderCellType.Roof))
                        renderWindow.DrawMapCell(cellInfo, skewX + px + (int)ox + xOffset, skewY + py + (int)oy - 80);
                }
            }




        }


        public void Update(long ms)
        {
            foreach (var character in characterRenderers)
                character.Update(ms);

            if (FocusedPlayerId != 0)
            {
                var player = gameState.PlayerInfos.FirstOrDefault(x => x.LocationDetails.PlayerId == FocusedPlayerId);
                if (player != null)
                {
                    // TODO: Maybe smooth movement? Maybe not?
                    CameraLocation = new PointF(player.LocationDetails.PlayerX, player.LocationDetails.PlayerY);
                }
            }
        }


    }
}
