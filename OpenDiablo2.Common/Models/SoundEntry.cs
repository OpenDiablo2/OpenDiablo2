using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public struct SoundEntry
    {
        public string Handle { get; set; }
        public int Index { get; set; }
        public string FileName { get; set; }
        public byte Volume { get; set; }
        public int GroupSize { get; set; }
        public bool Loop { get; set; }
        public int FadeIn { get; set; }
        public int FadeOut { get; set; }
        public int DeferInst { get; set; }
        public int StopInst { get; set; }
        public int Duration { get; set; }
        public int Compound { get; set; }
        public bool Reverb { get; set; }
        public int Falloff { get; set; }
        public int Cache { get; set; }
        public bool AsyncOnly { get; set; }
        public int Priority { get; set; }
        public int Stream { get; set; }
        public int Stereo { get; set; }
        public int Tracking { get; set; }
        public int Solo { get; set; }
        public int MusicVol { get; set; }
        public int Block1 { get; set; }
        public int Block2 { get; set; }
        public int Block3 { get; set; }
    }

    public static class SoundEntryHelper
    {
        public static SoundEntry ToSoundEntry(this string source)
        {
            var props = source.Split('\t');
            return new SoundEntry
            {
                Handle = props[0],
                Index = Convert.ToInt32(props[1]),
                FileName = props[2],
                Volume = Convert.ToByte(props[3]),
                GroupSize = Convert.ToInt32(props[4]),
                Loop = Convert.ToInt32(props[5]) == 1,
                FadeIn = Convert.ToInt32(props[6]),
                FadeOut = Convert.ToInt32(props[7]),
                DeferInst = Convert.ToInt32(props[8]),
                StopInst = Convert.ToInt32(props[9]),
                Duration = Convert.ToInt32(props[10]),
                Compound = Convert.ToInt32(props[11]),
                Reverb = Convert.ToInt32(props[12]) == 1,
                Falloff = Convert.ToInt32(props[13]),
                Cache = Convert.ToInt32(props[14]),
                AsyncOnly = Convert.ToInt32(props[15]) == 1,
                Priority = Convert.ToInt32(props[16]),
                Stream = Convert.ToInt32(props[17]),
                Stereo = Convert.ToInt32(props[18]),
                Tracking = Convert.ToInt32(props[19]),
                Solo = Convert.ToInt32(props[20]),
                MusicVol = Convert.ToInt32(props[21]),
                Block1 = Convert.ToInt32(props[22]),
                Block2 = Convert.ToInt32(props[23]),
                Block3 = Convert.ToInt32(props[24]),
            };
        }
    }
}
