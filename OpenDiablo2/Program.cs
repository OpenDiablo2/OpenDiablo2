using CommandLine;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Autofac;
using OpenDiablo2.Common.Models;
using System.Reflection;
using OpenDiablo2.Common.Interfaces;
using System.Diagnostics;

namespace OpenDiablo2
{
    static class Program
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        static void Main(string[] args)
        {
            log.Info("OpenDiablo 2: The Free and Open Source Diablo 2 clone!");

#if !DEBUG
            try
            {
#endif
            BuildContainer()
                .ResolveCommandLineOptions(args)
                .Resolve<IGameEngine>()
                .Run();
#if !DEBUG
            }
            catch (Exception ex)
            {
                log.Fatal("Uncaught exception detected, the game has been terminated!", ex);
            }
#endif
        }

        static IContainer BuildContainer() => new ContainerBuilder()
            .RegisterLocalTypes()
            .LoadAssemblyModules()
            .Build();


        static IContainer ResolveCommandLineOptions(this IContainer container, IEnumerable<string> args)
        {
            var globalConfiguration = container.Resolve<GlobalConfiguration>();

            Parser.Default.ParseArguments<CommandLineOptions>(args).WithParsed<CommandLineOptions>(o =>
            {
                globalConfiguration.BaseDataPath = Path.GetFullPath(o.DataPath ?? Directory.GetCurrentDirectory());
            });

            return container;
        }

        static ContainerBuilder RegisterLocalTypes(this ContainerBuilder containerBuilder)
        {
            containerBuilder.RegisterType<GlobalConfiguration>().AsSelf().SingleInstance();

            containerBuilder.Register<Func<string, IScene>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sceneName) => componentContext.ResolveKeyed<IScene>(sceneName);
            });

            return containerBuilder;
        }

        static ContainerBuilder LoadAssemblyModules(this ContainerBuilder containerBuilder)
        {
            var filesToLoad = Directory.GetFiles(Directory.GetCurrentDirectory(), "*.dll");
            foreach (var file in filesToLoad)
            {
                try
                {
                    var assembly = Assembly.LoadFrom(file);
                    containerBuilder.RegisterAssemblyModules(assembly);

                }
                catch { /* Silently ignore assembly load errors as not all DLLs are our modules... */ }
            }
            return containerBuilder;
        }
    }
}
