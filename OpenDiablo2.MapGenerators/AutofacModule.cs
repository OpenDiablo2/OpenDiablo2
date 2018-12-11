using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Autofac;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.MapGenerators
{
    public sealed class AutofacModule : Module
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        protected override void Load(ContainerBuilder builder)
        {
            log.Info("Configuring OpenDiablo2.MapGenerators service implementations.");

            var types = ThisAssembly.GetTypes().Where(x => typeof(IRandomizedMapGenerator).IsAssignableFrom(x) && x.IsClass);
            foreach (var type in types)
            {
                var att = type.GetCustomAttributes(true).First(x => (x is RandomizedMapAttribute)) as RandomizedMapAttribute;
                builder
                    .RegisterType(type)
                    .Keyed<IRandomizedMapGenerator>(att.MapName)
                    .InstancePerDependency();
            }
        }
    }
}
