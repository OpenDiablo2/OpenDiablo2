using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class LevelType
    {
        public string Name { get; set; }
        public int Id { get; set; }
        public string[] File { get; set; } = new string[32];
        public bool Beta { get; set; }
        public int Act { get; set; }
    }

    public static class LevelTypeHelper
    {
        public static LevelType ToLevelType(this string[] row)
        {
            var result = new LevelType
            {
                Name = row[0],
                Id = Convert.ToInt32(row[1]),
                Beta = Convert.ToInt32(row[34]) == 1,
                Act = Convert.ToInt32(row[35])
            };


            for (int i = 0; i < 32; i++)
            {
                result.File[i] = row[i + 2];
            }

            return result;
        }
    }
}
