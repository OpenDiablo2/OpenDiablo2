using Autofac;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Core.GameState_;
using OpenDiablo2.Core.UI;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core
{
    public sealed class AutofacModule : Module
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        protected override void Load(ContainerBuilder builder)
        {
            log.Info("Configuring OpenDiablo2.Core service implementations.");

            builder.RegisterType<GameEngine>().AsImplementedInterfaces().SingleInstance();
            builder.RegisterType<MPQProvider>().As<IMPQProvider>().SingleInstance();
            builder.RegisterType<ResourceManager>().As<IResourceManager>().SingleInstance();
            builder.RegisterType<TextDictionary>().As<ITextDictionary>().SingleInstance();
            builder.RegisterType<Button>().AsSelf().InstancePerDependency(); // TODO: Never register as Self() if we aren't in common...
            builder.RegisterType<TextBox>().AsSelf().InstancePerDependency(); // TODO: Never register as Self() if we aren't in common...
            builder.RegisterType<GameState>().AsSelf().SingleInstance(); // TODO: Never register as Self() if we aren't in common...
            builder.RegisterType<EngineDataManager>().As<IEngineDataManager>().SingleInstance();
        }
    }
}
