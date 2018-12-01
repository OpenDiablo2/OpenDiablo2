using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eMessageFrameType
    {
        None = 0x00,
        SetSeed = 0x01,
        JoinGame = 0x02,

        MAX = 0xFF, // NOTE:
        // You absolutely cannot have a higher ID than this without
        // changing the message header to multi-byte for ALL frame types!!!
    }
}
