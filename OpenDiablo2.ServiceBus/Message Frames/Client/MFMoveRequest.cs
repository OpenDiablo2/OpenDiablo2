using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.MoveRequest)]
    public sealed class MFMoveRequest : IMessageFrame
    {
        public int Direction { get; set; } = 0;
        public eMovementType MovementType { get; set; } = eMovementType.Stopped;

        public byte[] Data
        {
            get => new byte[] { (byte)Direction, (byte)MovementType };
            set
            {
                Direction = value[0];
                MovementType = (eMovementType)value[1];
            }
        }

        public MFMoveRequest() { }

        public MFMoveRequest(int direction, eMovementType movementType)
        {
            this.Direction = direction;
            this.MovementType = movementType;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnMoveRequest(clientHash, Direction, MovementType);
    }
}
