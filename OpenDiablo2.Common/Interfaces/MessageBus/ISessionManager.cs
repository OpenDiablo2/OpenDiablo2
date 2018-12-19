using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{

    public interface ISessionManager : ISessionEventProvider, IDisposable
    {
        void Initialize();
        void Stop();

        void JoinGame(string playerName, eHero heroType);
        void MoveRequest(PointF targetCell, eMovementType movementType);
        void UpdateEquipment(string slot, ItemInstance itemInstance);
    }
}
