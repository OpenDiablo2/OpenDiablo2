using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class AnimationData
    {
        public string COFName { get; set; }
        public int FramesPerDirection { get; set; }
        public int AnimationSpeed { get; set; }
        public byte[] Flags { get; set; }

        public static Dictionary<string, List<AnimationData>> LoadFromStream(Stream stream)
        {
            var result = new Dictionary<string, List<AnimationData>>();
            var br = new BinaryReader(stream);

            while(stream.Length > stream.Position)
            {
                var dataCount = br.ReadInt32();
                
                for (int i = 0; i < dataCount; ++i)
                {
                    var data = new AnimationData
                    {
                        COFName = Encoding.ASCII.GetString(br.ReadBytes(8)).Trim('\0'),
                        FramesPerDirection = br.ReadInt32(),
                        AnimationSpeed = br.ReadInt32(),
                        Flags = br.ReadBytes(144),
                    };

                    if (!result.ContainsKey(data.COFName.ToUpper()))
                        result[data.COFName.ToUpper()] = new List<AnimationData>();

                    result[data.COFName.ToUpper()].Add(data);
                }
            }

            br.Dispose();
            return result;
        }
    }
}
