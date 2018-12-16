using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.ServiceBus.Message_Frames.Server
{
    [MessageFrame(eMessageFrameType.LocatePlayers)]
    public sealed class MFLocatePlayers : IMessageFrame
    {
        public IEnumerable<LocationDetails> LocationDetails { get; set; }

        public void LoadFrom(BinaryReader br)
        {
            var count = br.ReadUInt16();
            var result = new List<LocationDetails>();

            for (var i = 0; i < count; i++)
            {
                result.Add(Common.Models.LocationDetails.FromBytes(br));
            }

            LocationDetails = result;
        }

        public void WriteTo(BinaryWriter bw)
        {
            bw.Write((UInt16)LocationDetails.Count());
            foreach (var locationDetail in LocationDetails)
                locationDetail.WriteBytes(bw);
        }

        public MFLocatePlayers() { }
        public MFLocatePlayers(IEnumerable<LocationDetails> locationDetails)
        {
            this.LocationDetails = locationDetails;
        }

        public void Process(int clientHash, ISessionEventProvider sessionEventProvider)
            => sessionEventProvider.OnLocatePlayers(clientHash, LocationDetails);
    }
}
