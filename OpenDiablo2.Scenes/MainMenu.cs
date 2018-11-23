using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Core.UI;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace OpenDiablo2.Scenes
{
    [Scene("Main Menu")]
    public class MainMenu : IScene
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IPaletteProvider paletteProvider;
        private readonly IMPQProvider mpqProvider;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IMusicProvider musicProvider;
        private readonly ISceneManager sceneManager;

        private float logoFrame;
        private ISprite backgroundSprite, diabloLogoLeft, diabloLogoRight, diabloLogoLeftBlack, diabloLogoRightBlack;
        private IFont labelFont;
        private ILabel versionLabel, urlLabel;
        private Button btnSinglePlayer, btnExit, btnWebsite;

        public MainMenu(
            IRenderWindow renderWindow,
            IPaletteProvider paletteProvider,
            IMPQProvider mpqProvider,
            IMouseInfoProvider mouseInfoProvider,
            IMusicProvider musicProvider,
            ISceneManager sceneManager,
            Func<eButtonType, Button> createButton
            )
        {
            this.renderWindow = renderWindow;
            this.paletteProvider = paletteProvider;
            this.mpqProvider = mpqProvider;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sceneManager = sceneManager;

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.GameSelectScreen, Palettes.Sky);
            diabloLogoLeft = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, Palettes.Units, new Point(400, 120));
            diabloLogoLeft.Blend = true;
            diabloLogoRight = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoFireRight, Palettes.Units, new Point(400, 120));
            diabloLogoRight.Blend = true;
            diabloLogoLeftBlack = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, Palettes.Units, new Point(400, 120));
            diabloLogoRightBlack = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, Palettes.Units, new Point(400, 120));

            btnSinglePlayer = createButton(eButtonType.Wide);
            btnSinglePlayer.Text = "Single Player".ToUpper();
            btnSinglePlayer.Location = new Point(264, 290);
            btnSinglePlayer.OnActivate = OnSinglePlayerClicked;

            btnWebsite = createButton(eButtonType.Wide);
            btnWebsite.Text = "Visit Github".ToUpper();
            btnWebsite.Location = new Point(264, 460);
            btnWebsite.OnActivate = OnVisitWebsiteClicked;

            btnExit = createButton(eButtonType.Wide);
            btnExit.Text = "Exit Diablo II".ToUpper();
            btnExit.Location = new Point(264, 500);
            btnExit.OnActivate = OnExitClicked;

            labelFont = renderWindow.LoadFont(ResourcePaths.Font16, Palettes.Static);
            versionLabel = renderWindow.CreateLabel(labelFont, new Point(50, 555), "v0.01 Pre-Alpha");
            urlLabel = renderWindow.CreateLabel(labelFont, new Point(50, 569), "https://github.com/essial/OpenDiablo2/");
            urlLabel.TextColor = Color.Magenta;

            var loadingSprite = renderWindow.LoadSprite(ResourcePaths.LoadingScreen, Palettes.Loading, new Point(300, 400));

            renderWindow.Clear();
            renderWindow.Draw(loadingSprite);
            renderWindow.Sync();

            /*
            musicProvider.LoadSong(mpqProvider.GetStream("data\\global\\music\\introedit.wav"));

            musicProvider.PlaySong();
            */
        }

        private void OnVisitWebsiteClicked()
            => System.Diagnostics.Process.Start("https://github.com/essial/OpenDiablo2/");

        public void Render()
        {
            // Render the background
            renderWindow.Draw(backgroundSprite, 4, 3, 0);

            // Render the flaming diablo 2 logo
            renderWindow.Draw(diabloLogoLeftBlack, (int)(diabloLogoLeftBlack.TotalFrames * logoFrame));
            renderWindow.Draw(diabloLogoRightBlack, (int)(diabloLogoRightBlack.TotalFrames * logoFrame));
            renderWindow.Draw(diabloLogoLeft, (int)(diabloLogoLeft.TotalFrames * logoFrame));
            renderWindow.Draw(diabloLogoRight, (int)(diabloLogoRight.TotalFrames * logoFrame));

            // Render the text
            renderWindow.Draw(versionLabel);
            renderWindow.Draw(urlLabel);

            // Render the UI buttons
            btnSinglePlayer.Render();
            btnWebsite.Render();
            btnExit.Render();

            //wideButton.Location = new Point(264, 290);
            //renderWindow.Draw(wideButton, 2, 1, 0);
        }

        public void Update(long ms)
        {
            float seconds = ((float)ms / 1000f);
            logoFrame += seconds;
            while (logoFrame >= 1f)
                logoFrame -= 1f;

            btnSinglePlayer.Update();
            btnWebsite.Update();
            btnExit.Update();

        }

        public void Dispose()
        {
            backgroundSprite.Dispose();
            diabloLogoLeft.Dispose();
            diabloLogoRight.Dispose();
            diabloLogoLeftBlack.Dispose();
            diabloLogoRightBlack.Dispose();
            labelFont.Dispose();
            versionLabel.Dispose();
            urlLabel.Dispose();
            btnSinglePlayer.Dispose();
            btnExit.Dispose();
            btnWebsite.Dispose();
        }

        private void OnSinglePlayerClicked()
            => sceneManager.ChangeScene("Select Hero Class");

        private void OnExitClicked()
        {
            renderWindow.Quit();
        }
    }
}
