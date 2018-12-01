using System;
using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.MessageBus;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.LocatePlayers)]
    public sealed class MFLocatePlayers : IMessageFrame
    {
        public IEnumerable<PlayerLocationDetails> LocationDetails { get; set; } 

        public byte[] Data
        {
            get
            {
                var result = new List<byte>();
                result.AddRange(BitConverter.GetBytes((UInt16)LocationDetails.Count()));
                result.AddRange(LocationDetails.SelectMany(x => x.GetBytes()));
                return result.ToArray();
            }

            set
            {
                var count = BitConverter.ToUInt16(value, 0);
                var result = new List<PlayerLocationDetails>();
                
                for(var i = 0; i < count; i++)
                {
                    result.Add(PlayerLocationDetails.FromBytes(value, 2 + (i * PlayerLocationDetails.SizeInBytes)));
                }

                LocationDetails = result;
            }
        }

        public MFLocatePlayers() { }
        public MFLocatePlayers(IEnumerable<PlayerLocationDetails> locationDetails)
        {
            this.LocationDetails = locationDetails;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnLocatePlayers(clientHash, LocationDetails);
    }
}
