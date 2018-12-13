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
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{
    [Scene(eSceneType.MainMenu)]
    public sealed class MainMenu : IScene
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISceneManager sceneManager;

        private float logoFrame;
        private readonly ISprite backgroundSprite, diabloLogoLeft, diabloLogoRight, diabloLogoLeftBlack, diabloLogoRightBlack;
        private readonly IFont labelFont;
        private readonly ILabel versionLabel, urlLabel;
        private readonly IButton btnSinglePlayer, btnExit, btnWebsite, btnCredits;

        public MainMenu(
            IRenderWindow renderWindow,
            ISceneManager sceneManager,
            IResourceManager resourceManager,
            ISoundProvider soundProvider,
            IMPQProvider mpqProvider,
            Func<eButtonType, IButton> createButton
            )
        {
            this.renderWindow = renderWindow;
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
            btnWebsite.Location = new Point(264, 330);
            btnWebsite.OnActivate = OnVisitWebsiteClicked;

            btnExit = createButton(eButtonType.Wide);
            btnExit.Text = "Exit Diablo II".ToUpper();
            btnExit.Location = new Point(264, 500);
            btnExit.OnActivate = OnExitClicked;

            btnCredits = createButton(eButtonType.Short);
            btnCredits.Text = "Credits".ToUpper(); /* TODO: We apparently need a 'half font' option... */
            btnCredits.Location = new Point(264, 470);
            btnCredits.OnActivate = OnCreditsClicked;

            labelFont = renderWindow.LoadFont(ResourcePaths.Font16, Palettes.Static);
            versionLabel = renderWindow.CreateLabel(labelFont, new Point(50, 555), "v0.02 Pre-Alpha");
            urlLabel = renderWindow.CreateLabel(labelFont, new Point(50, 569), "https://github.com/essial/OpenDiablo2/");
            urlLabel.TextColor = Color.Magenta;

            soundProvider.LoadSong(mpqProvider.GetStream(ResourcePaths.BGMTitle));
            soundProvider.PlaySong();
        }

        private void OnCreditsClicked()
            => sceneManager.ChangeScene(eSceneType.Credits);

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
            btnCredits.Render();

            //wideButton.Location = new Point(264, 290);
            //renderWindow.Draw(wideButton, 2, 1, 0);
        }

        public void Update(long ms)
        {
            float seconds = ms / 1000f;
            logoFrame += seconds;
            while (logoFrame >= 1f)
                logoFrame -= 1f;

            btnSinglePlayer.Update();
            btnWebsite.Update();
            btnExit.Update();
            btnCredits.Update();

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
            => sceneManager.ChangeScene(eSceneType.SelectHeroClass);

        private void OnExitClicked()
        {
            renderWindow.Quit();
        }
    }
}
