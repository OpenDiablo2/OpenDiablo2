using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Threading;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class GameEngine : IGameEngine, IPaletteProvider, ISceneManager
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly GlobalConfiguration globalConfig;
        private readonly IMPQProvider mpqProvider;
        private readonly Func<IRenderWindow> getRenderWindow;
        private readonly Func<string, IScene> getScene;
        private readonly Func<IResourceManager> getResourceManager;
        private readonly Func<IGameState> getGameState;

        private IScene currentScene;
        private IScene nextScene = null;
        private ISprite mouseSprite;

        private readonly MPQ[] MPQs;

        private Dictionary<string, SoundEntry> soundTable = new Dictionary<string, SoundEntry>();
        public Dictionary<string, Palette> PaletteTable { get; private set; } = new Dictionary<string, Palette>();


        public GameEngine(
            GlobalConfiguration globalConfig,
            IMPQProvider mpqProvider,
            Func<IRenderWindow> getRenderWindow,
            Func<string, IScene> getScene,
            Func<IResourceManager> getResourceManager,
            Func<IGameState> getGameState
            )
        {
            this.globalConfig = globalConfig;
            this.mpqProvider = mpqProvider;
            this.getRenderWindow = getRenderWindow;
            this.getScene = getScene;
            this.getResourceManager = getResourceManager;
            this.getGameState = getGameState;

            MPQs = mpqProvider.GetMPQs().ToArray();
        }

        private void LoadPalettes()
        {
            log.Info("Loading palettes");
            var paletteFiles = MPQs.SelectMany(x => x.Files).Where(x => x.StartsWith("data\\global\\palette\\") && x.EndsWith(".dat"));
            foreach (var paletteFile in paletteFiles)
            {
                var paletteNameParts = paletteFile.Split('\\');
                var paletteName = paletteNameParts[paletteNameParts.Count() - 2];
                PaletteTable[paletteName] = getResourceManager().GetPalette(paletteFile);
            }
        }

        private void LoadSoundData()
        {
            log.Info("Loading sound configuration data");
            var soundDescFile = mpqProvider.GetTextFile("data\\global\\excel\\Sounds.txt");
            
            foreach (var row in soundDescFile.Skip(1).Where(x => !String.IsNullOrWhiteSpace(x)))
            {
                var soundEntry = row.ToSoundEntry();
                soundTable[soundEntry.Handle] = soundEntry;
            }
            
        }

        public void Run()
        {
            var renderWindow = getRenderWindow();

            LoadPalettes();
            LoadSoundData();

            mouseSprite = renderWindow.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units);
            var cursor = renderWindow.LoadCursor(mouseSprite, 0, new Point(0, 3));
            renderWindow.MouseCursor = cursor;
            
            currentScene = getScene("Main Menu");
            var lastTicks = renderWindow.GetTicks();
            while (getRenderWindow().IsRunning)
            {
                var curTicks = renderWindow.GetTicks();
                var ms = curTicks - lastTicks;

                if (ms < 0)
                    continue;

                if (ms < 33)
                {
                    Thread.Sleep(33 - (int)ms);
                    continue;
                }

                // Prevent falco-punch updates
                if (ms > 1000)
                {
                    lastTicks = renderWindow.GetTicks();
                    continue;
                }

                lastTicks = curTicks;

                lock (getGameState().ThreadLocker)
                {
                    getGameState().Update(ms);
                    getRenderWindow().Update();
                    currentScene.Update(ms);

                    if (nextScene!= null)
                    {
                        currentScene = nextScene;
                        nextScene = null;
                        continue;
                    }

                    renderWindow.Clear();
                    currentScene.Render();

                    renderWindow.Sync();
                }
            }
        }

        public void Dispose()
        {
            currentScene?.Dispose();
        }

        public void ChangeScene(string sceneName)
            => nextScene = getScene(sceneName);
    }
}
