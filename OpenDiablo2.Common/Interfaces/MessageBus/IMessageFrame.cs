using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMessageFrame
    {
        void LoadFrom(BinaryReader br);
        void WriteTo(BinaryWriter bw);
        void Process(int clientHash, ISessionEventProvider sessionEventProvider);
    }
}
