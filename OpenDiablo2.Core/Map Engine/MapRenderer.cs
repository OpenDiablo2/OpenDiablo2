using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Core.Map_Engine
{
    public sealed class MapRenderer : IMapRenderer
    {
        private readonly IGameState gameState;
        private readonly IRenderWindow renderWindow;
        private readonly Func<IGameHUD> getGameHud;

        private readonly List<ICharacterRenderer> _characterRenderers = new List<ICharacterRenderer>();

        public Guid FocusedPlayerId { get; set; } = Guid.Empty;
        public int CameraOffset { get; set; }

        private PointF _cameraLocation = new PointF();
        public PointF CameraLocation
        {
            get => _cameraLocation;
            set
            {
                // ReSharper disable once RedundantCheckBeforeAssignment (This is a false positive)
                if (_cameraLocation == value)
                    return;

                _cameraLocation = value;
            }
        }

        private const int
            CellSizeX = 160,
            CellSizeY = 80,
            CellSizeXHalf = CellSizeX / 2,
            CellSizeYHalf = CellSizeY / 2;

        public MapRenderer(
            IGameState gameState,
            IRenderWindow renderWindow,
            ISessionManager sessionManager,
            Func<IGameHUD> getGameHud
            )
        {
            this.gameState = gameState;
            this.renderWindow = renderWindow;
            this.getGameHud = getGameHud;

            sessionManager.OnPlayerInfo += OnPlayerInfo;
            sessionManager.OnLocatePlayers += OnLocatePlayers;
        }

        private void OnLocatePlayers(int clientHash, IEnumerable<LocationDetails> locationDetails)
        {
            foreach (var locationDetail in locationDetails)
            {
                var characterRenderer = _characterRenderers.FirstOrDefault(x => x.UID == locationDetail.UID);
                var player = gameState.PlayerInfos.FirstOrDefault(x => x.UID == locationDetail.UID);
                if (characterRenderer == null)
                {
                    // TODO: Should we log this?
                    continue;
                }
                var newDirection = locationDetail.MovementDirection != player.MovementDirection;
                var stanceChanged = locationDetail.MovementType != player.MovementType;

                locationDetail.CopyMobLocationDetailsTo(player);

                if (newDirection || stanceChanged)
                    characterRenderer.ResetAnimationData();
            }
        }

        private void OnPlayerInfo(int clientHash, IEnumerable<PlayerInfo> playerInfo)
        {
            // Remove character renderers for players that no longer exist...
            _characterRenderers.RemoveAll(x => playerInfo.Any(z => z.UID == x.UID));

            // Add character renderers for characters that now exist
            foreach (var info in playerInfo.Where(x => _characterRenderers.All(z => x.UID != z.UID)).ToArray())
            {
                var cr = renderWindow.CreateCharacterRenderer();
                cr.UID = info.UID;
                _characterRenderers.Add(cr);
                cr.ResetAnimationData();
            }
        }

        private const int SkewX = 400;
        private const int SkewY = 300;

        public void Render()
        {
            var xOffset = (getGameHud().IsRightPanelVisible ? -200 : 0) + (getGameHud().IsLeftPanelVisible ? 200 : 0);

            var cx = -(_cameraLocation.X - Math.Truncate(_cameraLocation.X));
            var cy = -(_cameraLocation.Y - Math.Truncate(_cameraLocation.Y));

            for (var ty = -7; ty <= 9; ty++)
            {
                for (var tx = -8; tx <= 8; tx++)
                {
                    var ax = (int)(tx + Math.Truncate(_cameraLocation.X));
                    var ay = (int)(ty + Math.Truncate(_cameraLocation.Y));

                    var px = (tx - ty) * CellSizeXHalf;
                    var py = (tx + ty) * CellSizeYHalf;

                    var ox = (cx - cy) * CellSizeXHalf;
                    var oy = (cx + cy) * CellSizeYHalf;


                    foreach (var cellInfo in gameState.GetMapCellInfo(ax, ay, eRenderCellType.WallLower))
                        renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy + 80);

                    foreach (var cellInfo in gameState.GetMapCellInfo(ax, ay, eRenderCellType.Floor))
                        renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy);


                    foreach (var cellInfo in gameState.GetMapCellInfo(ax, ay, eRenderCellType.WallNormal))
                        renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy + 80);

                    foreach (var player in gameState.PlayerInfos.Where(x =>
                        (int)Math.Truncate(x.X) == ax &&
                        (int)Math.Truncate(x.Y) == ay)
                    )
                    {
                        var ptx = player.X - Math.Truncate(_cameraLocation.X);
                        var pty = player.Y - Math.Truncate(_cameraLocation.Y);

                        var ppx = (int)((ptx - pty) * CellSizeXHalf);
                        var ppy = (int)((ptx + pty) * CellSizeYHalf);

                        _characterRenderers.FirstOrDefault(x => x.UID == player.UID)
                            .Render(SkewX + (int)ppx + (int)ox + xOffset, SkewY + (int)ppy + (int)oy);
                    }

                    foreach (var cellInfo in gameState.GetMapCellInfo(ax, ay, eRenderCellType.Roof))
                        renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy - 80);
                }
            }




        }


        public void Update(long ms)
        {
            foreach (var character in _characterRenderers)
                character.Update(ms);

            if (FocusedPlayerId == Guid.Empty)
                return;

            var player = gameState.PlayerInfos.FirstOrDefault(x => x.UID == FocusedPlayerId);
            if (player == null)
                return;

            CameraLocation = new PointF(player.X, player.Y);

        }

        public PointF GetCellPositionAt(int x, int y)
        {
            var xOffset = (getGameHud().IsRightPanelVisible ? -200 : 0) + (getGameHud().IsLeftPanelVisible ? 200 : 0);
            var mx = x - 400 - xOffset;
            var my = y - 300;
            return new PointF
            {
                X = (float)Math.Round((mx / (float)CellSizeXHalf) + (my / (float)CellSizeYHalf) / 2f, 1) + _cameraLocation.X,
                Y = (float)Math.Round((my / (float)CellSizeYHalf) - (mx / (float)CellSizeXHalf) / 2f, 1) + _cameraLocation.Y
            };
        }
        /*
        var mx = x - 400 - gameState.CameraOffset;
        var my = y - 300;

        var tx = (mx / 60f + my / 40f) / 2f;
        var ty = (my / 40f - (mx / 60f)) / 2f;
        */



    }
}
