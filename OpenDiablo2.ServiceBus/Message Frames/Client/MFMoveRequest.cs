using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Drawing;
using System.IO;

namespace OpenDiablo2.ServiceBus.Message_Frames.Client
{
    [MessageFrame(eMessageFrameType.MoveRequest)]
    public sealed class MFMoveRequest : IMessageFrame
    {
        public PointF Target { get; set; }
        public eMovementType MovementType { get; set; } = eMovementType.Stopped;

        public MFMoveRequest() { }

        public MFMoveRequest(PointF targetCell, eMovementType movementType)
        {
            this.Target = targetCell;
            this.MovementType = movementType;
        }

        public void LoadFrom(BinaryReader br)
        {
            MovementType = (eMovementType)br.ReadByte();
            Target = new PointF
            {
                X = br.ReadSingle(),
                Y = br.ReadSingle()
            };
        }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write((byte)MovementType);
            bw.Write((Single)Target.X);
            bw.Write((Single)Target.Y);
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnMoveRequest(clientHash, Target, MovementType);
    }
}
