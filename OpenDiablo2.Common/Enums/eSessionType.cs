using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    /// <summary>Defines the type of gameplay session we are running.</summary>
    public enum eSessionType
    {
        /// <summary>This session is an offline single player game.</summary>
        Local,

        /// <summary>This session is a multiplayer game, and this instance is the server.</summary>
        Server,

        /// <summary>This session is a multiplayer game, and this instance is a client connected to an external server.</summary>
        Remote
    }
}
