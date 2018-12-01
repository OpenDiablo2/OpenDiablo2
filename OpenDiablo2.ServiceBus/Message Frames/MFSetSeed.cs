using System;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames
{
    [MessageFrame(eMessageFrameType.SetSeed)]
    public sealed class MFSetSeed : IMessageFrame
    {

        public byte[] Data
        {
            get => BitConverter.GetBytes(Seed);
            set => BitConverter.ToInt32(value, 0);
        }

        public Int32 Seed { get; private set; }

        public MFSetSeed()
        {
            Seed = (new Random()).Next();
        }

        public MFSetSeed(int seed)
        {
            Seed = seed;
        }

        public void Process(object sender, ISessionEventProvider sessionEventProvider)
        {
            sessionEventProvider.OnSetSeed?.Invoke(sender, Seed);
        }
    }
}
