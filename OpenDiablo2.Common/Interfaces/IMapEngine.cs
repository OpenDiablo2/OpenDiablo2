﻿using System;
using System.Drawing;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMapEngine
    {
        PointF CameraLocation { get; set; } 
        void Update(long ms);
        void Render();
        void NotifyMapChanged();
        MapCellInfo GetMapCellInfo(Guid mapId, Guid tileId);
        void SetMapCellInfo(Guid mapId, MapCellInfo cellInfo);
    }
}
