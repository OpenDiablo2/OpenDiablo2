using System;
using System.IO;
using System.Linq;
using System.Reflection;
using Autofac;
using CommandLine;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2
{
    static class Program
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);
        private static GlobalConfiguration globalConfiguration;
        static void Main(string[] args)
        {
            log.Info("OpenDiablo 2: The Free and Open Source Diablo 2 clone!");

            Parser.Default.ParseArguments<CommandLineOptions>(args).WithParsed(o => globalConfiguration = new GlobalConfiguration
            {
                BaseDataPath = Path.GetFullPath(o.DataPath ?? Directory.GetCurrentDirectory()),
                MouseMode = o.HardwareMouse == true ? eMouseMode.Hardware : eMouseMode.Software,
                HardwareMouseScale = o.MouseScale
            }).WithNotParsed(o =>
            {
                log.Warn($"Could not parse command line options.");
                globalConfiguration = new GlobalConfiguration { BaseDataPath = Directory.GetCurrentDirectory(), MouseMode = eMouseMode.Software };
            }); ;

#if !DEBUG
            try
            {
#endif
            BuildContainer()
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


        static ContainerBuilder RegisterLocalTypes(this ContainerBuilder containerBuilder)
        {
            containerBuilder.Register(x => globalConfiguration).AsSelf().SingleInstance();

            containerBuilder.Register<Func<string, IScene>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sceneName) => componentContext.ResolveKeyed<IScene>(sceneName);
            });

            containerBuilder.Register<Func<eButtonType, IButton>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (buttonType) => componentContext.Resolve<IButton>(new NamedParameter("buttonLayout", ButtonLayout.Values[buttonType]));
            });

            /* Uncomment the below if we support multiple textbox types
            containerBuilder.Register<Func<TextBox>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return () => componentContext.Resolve<TextBox>();
            });
            */
            return containerBuilder;
        }

        static ContainerBuilder LoadAssemblyModules(this ContainerBuilder containerBuilder)
        {
            var filesToLoad = Directory.GetFiles(Directory.GetCurrentDirectory(), "*.dll").Where(x => Path.GetFileName(x).StartsWith("OpenDiablo2."));
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
