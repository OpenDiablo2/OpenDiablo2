using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace OpenDiablo2.Core
{
    public sealed class GameEngine : IGameEngine, IPaletteProvider, ISceneManager
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IMPQProvider mpqProvider;
        private readonly Func<IRenderWindow> getRenderWindow;
        private readonly Func<IMouseInfoProvider> getMouseInfoProvider;
        private readonly Func<string, IScene> getScene;
        private readonly Func<IResourceManager> getResourceManager;

        private IScene currentScene;
        private IScene nextScene = null;
        private ISprite mouseSprite;

        private readonly MPQ[] MPQs;

        private Dictionary<string, SoundEntry> soundTable = new Dictionary<string, SoundEntry>();
        public Dictionary<string, Palette> PaletteTable { get; private set; } = new Dictionary<string, Palette>();
        private Stopwatch sw = new Stopwatch();


        public GameEngine(
            IMPQProvider mpqProvider,
            Func<IRenderWindow> getRenderWindow,
            Func<IMouseInfoProvider> getMouseInfoProvider,
            Func<string, IScene> getScene,
            Func<IResourceManager> getResourceManager
            )
        {
            this.mpqProvider = mpqProvider;
            this.getRenderWindow = getRenderWindow;
            this.getMouseInfoProvider = getMouseInfoProvider;
            this.getScene = getScene;
            this.getResourceManager = getResourceManager;

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
            var mouseInfoProvider = getMouseInfoProvider();

            LoadPalettes();
            LoadSoundData();

            mouseSprite = renderWindow.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units);


            currentScene = getScene("Main Menu");
            sw.Start();
            while (getRenderWindow().IsRunning)
            {
                while (sw.ElapsedMilliseconds < 40)
                    Thread.Sleep((int)Math.Min(1, 40 -sw.ElapsedMilliseconds)); // The original runs at about 25 fps.

                var ms = sw.ElapsedMilliseconds;

                // Prevent falco-punch updates
                if (ms > 1000)
                {
                    sw.Restart();
                    continue;
                }
                sw.Restart();
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

                // Draw the mouse
                renderWindow.Draw(mouseSprite, new Point(mouseInfoProvider.MouseX, mouseInfoProvider.MouseY + 3));

                renderWindow.Sync();
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
