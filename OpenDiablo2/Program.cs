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

        // We store the global configuration here so we can bind it and return it to anything that needs it.
        private static GlobalConfiguration globalConfiguration;

        static void Main(string[] args)
        {
            log.Info("OpenDiablo 2: The Free and Open Source Diablo 2 clone!\n" +
                     "This program comes with ABSOLUTELY NO WARRANTY.\n" +
                     "This is free software, and you are welcome to redistribute it\n" +
                     "under certain conditions; type `show c' for details.");

            // Parse the command-line arguments.
            Parser.Default.ParseArguments<CommandLineOptions>(args).WithParsed(o => globalConfiguration = new GlobalConfiguration
            {
                BaseDataPath = Path.GetFullPath(o.DataPath ?? Directory.GetCurrentDirectory()),
                MouseMode = o.HardwareMouse == true ? eMouseMode.Hardware : eMouseMode.Software,
                HardwareMouseScale = o.MouseScale,
                FullScreen = o.FullScreen
            }).WithNotParsed(o =>
            {
                log.Warn($"Could not parse command line options.");
                globalConfiguration = new GlobalConfiguration { BaseDataPath = Directory.GetCurrentDirectory(), MouseMode = eMouseMode.Software };
            }); ;

#if !DEBUG
            try
            {
#endif
            // Create the AutoFac DI container
            var container = BuildContainer();
            try
            {
                // Resolve the game engine
                using (var gameEngine = container.Resolve<IGameEngine>())
                {
                    // Start the game!
                    gameEngine.Run();
                }
            }
            finally
            {
                // Dispose the container, disposing any instantiated objects in the process
                container.Dispose();
            }


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

            containerBuilder.Register<Func<eSceneType, IScene>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (sceneType) => componentContext.ResolveKeyed<IScene>(sceneType);
            });

            containerBuilder.Register<Func<eButtonType, IButton>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (buttonType) => componentContext.Resolve<IButton>(new NamedParameter("buttonLayout", ButtonLayout.Values[buttonType]));
            });

            containerBuilder.Register<Func<ePanelFrameType, IPanelFrame>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (panelFrameType) => componentContext.Resolve<IPanelFrame>(new NamedParameter("panelFrameType", panelFrameType));
            });

            containerBuilder.Register<Func<eItemContainerType, IItemContainer>>(c =>
            {
                var componentContext = c.Resolve<IComponentContext>();
                return (itemContainerType) => componentContext.Resolve<IItemContainer>(new NamedParameter("itemContainerLayout", ItemContainerLayout.Values[itemContainerType]));
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
