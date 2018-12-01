using System;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Attributes
{
    [AttributeUsage(AttributeTargets.All, Inherited = false, AllowMultiple = true)]
    public sealed class MessageFrameAttribute : Attribute
    {
        public eMessageFrameType FrameType { get; private set; }

        // This is a positional argument
        public MessageFrameAttribute(eMessageFrameType frameType)
        {
            this.FrameType = frameType;
        }


    }
}
