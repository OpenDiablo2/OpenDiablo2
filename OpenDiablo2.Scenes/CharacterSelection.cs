using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{

    [Scene("Select Character")]
    public sealed class CharacterSelection : IScene
    {
        static readonly log4net.ILog log =
            log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IPaletteProvider paletteProvider;
        private readonly IMPQProvider mpqProvider;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IMusicProvider musicProvider;
        private readonly ISceneManager sceneManager;
        private readonly ITextDictionary textDictionary;
        private readonly IKeyboardInfoProvider keyboardInfoProvider;
        private readonly IGameState gameState;

        private ISprite backgroundSprite, largeButtonSprite, largeButtonSprite2;

        public CharacterSelection(IRenderWindow renderWindow, IPaletteProvider paletteProvider, IMPQProvider mpqProvider, IMouseInfoProvider mouseInfoProvider, IMusicProvider musicProvider, ISceneManager sceneManager, ITextDictionary textDictionary, IKeyboardInfoProvider keyboardInfoProvider, IGameState gameState)
        {
            this.renderWindow = renderWindow;
            this.paletteProvider = paletteProvider;
            this.mpqProvider = mpqProvider;
            this.mouseInfoProvider = mouseInfoProvider;
            this.musicProvider = musicProvider;
            this.sceneManager = sceneManager;
            this.textDictionary = textDictionary;
            this.keyboardInfoProvider = keyboardInfoProvider;
            this.gameState = gameState;
            
            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectionBackground, Palettes.Sky);
            largeButtonSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectionTallButton, Palettes.Sky);
            largeButtonSprite2 = renderWindow.LoadSprite(ResourcePaths.CharacterSelectionTallButton, Palettes.Sky);
        }

        public void Update(long ms)
        {
            var point = new Point(33, 527);
            var buttonSize = largeButtonSprite.FrameSize;
            
            var mouseX = mouseInfoProvider.MouseX - 33;
            var mouseY = mouseInfoProvider.MouseY - 527 + buttonSize.Height;

            if (mouseX > 0 && mouseX < buttonSize.Width && mouseY > 0 && mouseY < buttonSize.Height)
            {
                log.Info("redraw as pressed");
                renderWindow.Draw(largeButtonSprite, 1, new Point(33, 527));
            }
            else
            {
                log.Info("redraw as normal");
                renderWindow.Draw(largeButtonSprite, 0, new Point(33, 527));
            }
        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);
            
            renderWindow.Draw(largeButtonSprite, new Point(33, 527));
            renderWindow.Draw(largeButtonSprite2, new Point(433, 527));
        }

        private void OnExitClicked()
        {
            sceneManager.ChangeScene("Main Menu");

        }

        public void Dispose()
        {
            backgroundSprite.Dispose();
        }
    }
}