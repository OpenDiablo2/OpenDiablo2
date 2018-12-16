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

            builder.RegisterType<Cache>().As<ICache>().SingleInstance();
            builder.RegisterType<Button>().As<IButton>().InstancePerDependency();
            builder.RegisterType<EngineDataManager>().As<IEngineDataManager>().SingleInstance();
            builder.RegisterType<ItemManager>().As<IItemManager>().SingleInstance();
            builder.RegisterType<GameEngine>().AsImplementedInterfaces().SingleInstance();
            builder.RegisterType<GameState>().As<IGameState>().SingleInstance();
            builder.RegisterType<MapRenderer>().As<IMapRenderer>().SingleInstance();
            builder.RegisterType<GameHUD>().As<IGameHUD>().SingleInstance();
            builder.RegisterType<MiniPanel>().As<IMiniPanel>().InstancePerDependency();
            builder.RegisterType<PanelFrame>().As<IPanelFrame>().InstancePerDependency();
            builder.RegisterType<CharacterPanel>().AsImplementedInterfaces().InstancePerDependency();
            builder.RegisterType<InventoryPanel>().AsImplementedInterfaces().InstancePerDependency();
            builder.RegisterType<SkillsPanel>().AsImplementedInterfaces().InstancePerDependency();
            builder.RegisterType<ItemContainer>().As<IItemContainer>().InstancePerDependency();
            builder.RegisterType<MpqProvider>().As<IMPQProvider>().SingleInstance();
            builder.RegisterType<ResourceManager>().As<IResourceManager>().SingleInstance();
            builder.RegisterType<TextDictionary>().As<ITextDictionary>().SingleInstance();
            builder.RegisterType<TextBox>().As<ITextBox>().InstancePerDependency();
            builder.RegisterType<MobManager>().As<IMobManager>().SingleInstance(); // TODO: This needs to have client and server versions...
        }
    }
}
