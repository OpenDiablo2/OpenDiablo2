using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMessageFrame
    {
        byte[] Data { get; set; }
        void Process(object sender, ISessionEventProvider sessionEventProvider);
    }
}
