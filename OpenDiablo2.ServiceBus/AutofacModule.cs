/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

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
                var att = type.GetCustomAttributes(true).First(x => (x is MessageFrameAttribute)) as MessageFrameAttribute;
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
