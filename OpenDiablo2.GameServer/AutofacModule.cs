using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Autofac;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.GameServer_
{
    public sealed class AutofacModule : Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterType<GameServer>().As<IGameServer>().SingleInstance();
        }
    }
}
