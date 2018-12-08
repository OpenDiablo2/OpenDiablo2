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
        public UInt32 GetBit()
        {
            var result = (UInt32)(data[Offset / 8] >> (Offset % 8)) & 0x01;
            Offset++;
            BitsRead++;
            return result;
        }

        public void SkipBits(int bits) => Offset += bits;

        public byte GetByte() => (byte)GetBits(8);
        public Int32 GetInt32() => MakeSigned(GetBits(32), 32);
        public UInt32 GetUInt32() => GetBits(32);

        public UInt32 GetBits(int bits)
        {
            if (bits == 0)
                return 0;

            var result = 0U;
            for (var i = 0; i < bits; i++)
                result |= (GetBit() << i);

            return result;
        }

        public int GetSignedBits(int bits) => MakeSigned(GetBits(bits), bits);


        private Int32 MakeSigned(UInt32 value, int bits)
        {
            if (bits == 0)
                return 0;
            // If its a single bit, a value of 1 is -1 automagically
            if (bits == 1)
                return -(Int32)value;

            // If there is no sign bit, return the value as is
            if ((value & (1 << (bits - 1))) == 0)
                return (Int32)value;

            // We need to extend the signed bit out so that the negative value
            // representation still works with the 2s compliment rule.
            var result = UInt32.MaxValue;
            for (byte i = 0; i < bits; i++)
            {
                if (((value >> i) & 1) == 0)
                    result -= (UInt32)(1 << i);
            }

            return (Int32)result; // Force casting to a signed value
        }

    }
}
