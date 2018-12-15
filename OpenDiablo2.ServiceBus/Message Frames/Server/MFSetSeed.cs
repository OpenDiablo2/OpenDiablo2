using System;
using System.IO;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.SetSeed)]
    public sealed class MFSetSeed : IMessageFrame
    {

        public byte[] Data
        {
            get => BitConverter.GetBytes(Seed);
            set => Seed = BitConverter.ToInt32(value, 0);
        }

        public void LoadFrom(BinaryReader br)
        {
            Seed = br.ReadInt32();
        }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write((Int32)Seed);
        }

        public Int32 Seed { get; private set; }

        public MFSetSeed()
        {
        }

        public MFSetSeed(int seed)
        {
            Seed = seed;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnSetSeed?.Invoke(clientHash, Seed);

    }
}
