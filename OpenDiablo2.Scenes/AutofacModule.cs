using System.Linq;
using Autofac;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{
    public sealed class AutofacModule : Module
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        protected override void Load(ContainerBuilder builder)
        {
            log.Info("Configuring OpenDiablo2.Scenes service implementations.");

            var types = ThisAssembly.GetTypes().Where(x => typeof(IScene).IsAssignableFrom(x) && x.IsClass);
            foreach (var type in types)
            {
                var att = type.GetCustomAttributes(true).First(x => (x is SceneAttribute)) as SceneAttribute;
                builder
                    .RegisterType(type)
                    .Keyed<IScene>(att.SceneType)
                    .InstancePerDependency();
            }
        }
    }
}
