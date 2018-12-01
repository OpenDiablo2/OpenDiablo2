using System;
using System.Linq;
using Autofac;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.ServiceBus
{
    public sealed class AutofacModule : Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterType<SessionManager>().As<ISessionManager>().InstancePerLifetimeScope();
            builder.RegisterType<SessionServer>().As<ISessionServer>().InstancePerLifetimeScope();

            var types = ThisAssembly.GetTypes().Where(x => typeof(IMessageFrame).IsAssignableFrom(x) && x.IsClass);
            foreach (var type in types)
            {
                var att = type.GetCustomAttributes(true).First(x => typeof(MessageFrameAttribute).IsAssignableFrom(x.GetType())) as MessageFrameAttribute;
                builder
                    .RegisterType(type)
                    .Keyed<IMessageFrame>(att.FrameType)
                    .InstancePerDependency();
            }

            builder.Register<Func<eMessageFrameType, IMessageFrame>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (frameType) => componentContext.ResolveKeyed<IMessageFrame>(frameType);
            });

            builder.Register<Func<eSessionType, ISessionManager>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sessionType) => componentContext.Resolve<ISessionManager>(new NamedParameter("sessionType", sessionType));
            });

            builder.Register<Func<eSessionType, ISessionServer>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sessionType) => componentContext.Resolve<ISessionServer>(new NamedParameter("sessionType", sessionType));
            });

            
        }
    }
}
