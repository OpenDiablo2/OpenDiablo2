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
        private readonly IGameState _gameState;
        private readonly IRenderWindow _renderWindow;

        private readonly List<ICharacterRenderer> _characterRenderers = new List<ICharacterRenderer>();

        public int FocusedPlayerId { get; set; } = 0;

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
            CellSizeXHalf = 80,
            CellSizeYHalf = 40;

        public MapEngine(
            IGameState gameState,
            IRenderWindow renderWindow,
            ISessionManager sessionManager
            )
        {
            _gameState = gameState;
            _renderWindow = renderWindow;

            sessionManager.OnPlayerInfo += OnPlayerInfo;
            sessionManager.OnLocatePlayers += OnLocatePlayers;
        }

        private void OnLocatePlayers(int clientHash, IEnumerable<PlayerLocationDetails> playerLocationDetails)
        {
            foreach (var loc in playerLocationDetails)
            {
                var cr = _characterRenderers.FirstOrDefault(x => x.LocationDetails.PlayerId == loc.PlayerId);
                if (cr == null)
                {
                    // TODO: Should we log this?
                    continue;
                }
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
            _characterRenderers.RemoveAll(x => playerInfo.Any(z => z.UID == x.UID));

            // Update existing character renderers
            foreach (var cr in _characterRenderers)
            {
                var info = playerInfo.FirstOrDefault(x => x.UID == cr.UID);
                if (info == null)
                    continue;

                // TODO: This shouldn't be necessary...
                cr.LocationDetails = info.LocationDetails;
                cr.MobMode = info.MobMode;
                cr.Hero = info.Hero;
                cr.Equipment = info.Equipment;

            }

            // Add character renderers for characters that now exist
            foreach (var info in playerInfo.Where(x => _characterRenderers.All(z => x.UID != z.UID)).ToArray())
            {
                var cr = _renderWindow.CreateCharacterRenderer();
                cr.UID = info.UID;
                cr.LocationDetails = info.LocationDetails;
                cr.MobMode = info.MobMode;
                cr.Equipment = info.Equipment;
                cr.Hero = info.Hero;
                _characterRenderers.Add(cr);
                cr.ResetAnimationData();
            }
        }


        private const int SkewX = 400;
        private const int SkewY = 300;

        public void Render()
        {
            var xOffset = _gameState.CameraOffset;

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


                    foreach (var cellInfo in _gameState.GetMapCellInfo(ax, ay, eRenderCellType.WallLower))
                        _renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy + 80);

                    foreach (var cellInfo in _gameState.GetMapCellInfo(ax, ay, eRenderCellType.Floor))
                        _renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy);


                    foreach (var cellInfo in _gameState.GetMapCellInfo(ax, ay, eRenderCellType.WallNormal))
                        _renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy + 80);

                    foreach (var character in _characterRenderers.Where(x =>
                        (int)Math.Truncate(x.LocationDetails.PlayerX) == ax &&
                        (int)Math.Truncate(x.LocationDetails.PlayerY) == ay)
                    )
                    {
                        var ptx = character.LocationDetails.PlayerX - Math.Truncate(_cameraLocation.X);
                        var pty = character.LocationDetails.PlayerY - Math.Truncate(_cameraLocation.Y);

                        var ppx = (int)((ptx - pty) * CellSizeXHalf);
                        var ppy = (int)((ptx + pty) * CellSizeYHalf);

                        character.Render(SkewX + (int)ppx + (int)ox + xOffset, SkewY + (int)ppy + (int)oy);
                    }

                    foreach (var cellInfo in _gameState.GetMapCellInfo(ax, ay, eRenderCellType.Roof))
                        _renderWindow.DrawMapCell(cellInfo, SkewX + px + (int)ox + xOffset, SkewY + py + (int)oy - 80);
                }
            }




        }


        public void Update(long ms)
        {
            foreach (var character in _characterRenderers)
                character.Update(ms);

            if (FocusedPlayerId == 0)
                return;

            var player = _gameState.PlayerInfos.FirstOrDefault(x => x.LocationDetails.PlayerId == FocusedPlayerId);
            if (player == null)
                return;

            CameraLocation = new PointF(player.LocationDetails.PlayerX, player.LocationDetails.PlayerY);

        }


    }
}
