using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class ObjectTypeInfo
    {
        public string Name { get; internal set; }
        public string Token { get; internal set; }
        public bool Beta { get; internal set; }
    }

    public static class ObjectTypeInfoHelper
    {
        public static ObjectTypeInfo ToObjectTypeInfo(this string[] row)
            => new ObjectTypeInfo { Name = row[0], Token = row[1], Beta = Convert.ToInt32(row[2]) == 1 };
    }
}
