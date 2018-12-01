using Autofac;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Core.GameState_;
using OpenDiablo2.Core.Map_Engine;
using OpenDiablo2.Core.UI;

namespace OpenDiablo2.Core
{
    public sealed class AutofacModule : Module
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        protected override void Load(ContainerBuilder builder)
        {
            log.Info("Configuring OpenDiablo2.Core service implementations.");

            builder.RegisterType<Button>().As<IButton>().InstancePerDependency();
            builder.RegisterType<EngineDataManager>().As<IEngineDataManager>().SingleInstance();
            builder.RegisterType<GameEngine>().AsImplementedInterfaces().SingleInstance();
            builder.RegisterType<GameState>().As<IGameState>().SingleInstance();
            builder.RegisterType<MapEngine>().As<IMapEngine>().SingleInstance();
            builder.RegisterType<MiniPanel>().As<IMiniPanel>().InstancePerDependency();
            builder.RegisterType<PanelFrame>().As<IPanelFrame>().InstancePerDependency();
            builder.RegisterType<CharacterPanel>().As<ICharacterPanel>().InstancePerDependency();
            builder.RegisterType<InventoryPanel>().As<IInventoryPanel>().InstancePerDependency();
            builder.RegisterType<MPQProvider>().As<IMPQProvider>().SingleInstance();
            builder.RegisterType<ResourceManager>().As<IResourceManager>().SingleInstance();
            builder.RegisterType<TextDictionary>().As<ITextDictionary>().SingleInstance();
            builder.RegisterType<TextBox>().As<ITextBox>().InstancePerDependency();

            builder.RegisterType<MobManager>().As<IMobManager>().SingleInstance(); // TODO: This needs to have client and server versions...
        }
    }
}
