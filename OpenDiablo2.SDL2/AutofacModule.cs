using Autofac;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.SDL2_
{
    public sealed class AutofacModule : Module
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        protected override void Load(ContainerBuilder builder)
        {
            log.Info("Configuring OpenDiablo2.Core service implementations.");

            builder.RegisterType<SDL2RenderWindow>().AsImplementedInterfaces().SingleInstance();
            builder.RegisterType<SDL2MusicPlayer>().AsImplementedInterfaces().SingleInstance();
            
        }
    }
}
