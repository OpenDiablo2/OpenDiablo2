using System;
using Autofac;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus
{
    public sealed class AutofacModule : Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterType<LocalSessionManager>().AsSelf().InstancePerLifetimeScope();

            builder.Register<Func<eSessionType, ISessionManager>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sessionType) =>
                {
                    switch (sessionType)
                    {
                        case eSessionType.Local:
                            return componentContext.Resolve<LocalSessionManager>();
                        case eSessionType.Server:
                        case eSessionType.Remote:
                        default:
                            throw new ApplicationException("Unsupported session type.");
                    }
                };
            });
        }
    }
}
