using System;
using System.IO;

namespace OpenDiablo2.Common.Models
{
    public sealed class BitMuncher
    {
        private readonly byte[] data;
        
        public int Offset { get; private set; }
        public int BitsRead { get; private set; }

        public BitMuncher(byte[] data, int offset = 0)
        {
            this.data = data;
            this.Offset = offset;
        }

        public BitMuncher(BitMuncher source)
        {
            this.data = source.data;
            this.Offset = source.Offset;
        }
        public int GetBit()
        {
            var result = (data[Offset / 8] >> (Offset % 8)) & 0x01;
            Offset++;
            BitsRead++;
            return result;
        }

        public void SkipBits(int bits) => Offset += bits;

        public byte GetByte() => (byte)GetBits(8);
        public Int32 GetInt32() => MakeSigned(GetBits(32), 32);

        public int GetBits(int bits)
        {
            var result = 0;
            for (var i = 0; i < bits; i++)
                result |= (GetBit() << i);

            return result;
        }

        public int GetSignedBits(int bits) => MakeSigned(GetBits(bits), bits);

        private int MakeSigned(int value, int bits)
        {
            if (bits == 1)
                return -value;

            if ((value & (1 << (bits - 1))) == 0)
                return value;

            var result = UInt32.MaxValue;
            for (byte i = 0; i < bits; i++)
            {
                if (((value >> i) & 1) == 0)
                    result -= (UInt32)(1 << i);
            }

            var newResult = (int)result;
            return newResult;
        }

    }
}
