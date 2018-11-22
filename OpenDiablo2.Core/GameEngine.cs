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
    public sealed class GameEngine : IGameEngine, IPaletteProvider
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IMPQProvider mpqProvider;
        private readonly Func<IRenderWindow> getRenderWindow;
        private readonly Func<string, IScene> getScene;
        private IScene currentScene;

        private readonly MPQ[] MPQs;

        private Dictionary<string, SoundEntry> soundTable = new Dictionary<string, SoundEntry>();
        public Dictionary<string, Palette> PaletteTable { get; private set; } = new Dictionary<string, Palette>();
        private Stopwatch sw = new Stopwatch();


        public GameEngine(IMPQProvider mpqProvider, Func<IRenderWindow> getRenderWindow, Func<string, IScene> getScene)
        {
            this.mpqProvider = mpqProvider;
            this.getRenderWindow = getRenderWindow;
            this.getScene = getScene;

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
                PaletteTable[paletteName] = Palette.LoadFromStream(mpqProvider.GetStream(paletteFile), paletteName);
            }
        }

        private void LoadSoundData()
        {
            log.Info("Loading sound configuration data");
            foreach (var soundDescFile in mpqProvider.GetTextFile("data\\global\\excel\\Sounds.txt"))
            {
                foreach (var row in soundDescFile.Skip(1).Where(x => !String.IsNullOrWhiteSpace(x)))
                {
                    var soundEntry = row.ToSoundEntry();
                    soundTable[soundEntry.Handle] = soundEntry;
                }
            }
        }

        public void Run()
        {
            LoadPalettes();
            LoadSoundData();

            currentScene = getScene("Main Menu");
            sw.Start();
            while (getRenderWindow().IsRunning)
            {
                while (sw.ElapsedMilliseconds < 16)
                    Thread.Sleep(1); // Oh yes we did

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
                
                currentScene.Render();
            }
        }

        public void Dispose()
        {

        }
    }
}
