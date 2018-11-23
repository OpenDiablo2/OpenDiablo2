using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Core.UI;

namespace OpenDiablo2.Scenes
{
    [Scene("Select Hero Class")]
    public sealed class SelectHeroClass : IScene
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IPaletteProvider paletteProvider;
        private readonly IMPQProvider mpqProvider;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IMusicProvider musicProvider;
        private readonly ISceneManager sceneManager;

        private float secondTimer;
        private ISprite backgroundSprite, campfireSprite,
            barbarianUnselected, barbarianUnselectedH,
            sorceressUnselected, sorceressUnselectedH,
            necromancerUnselected, necromancerUnselectedH,
            paladinUnselected, paladinUnselectedH,
            amazonUnselected, amazonUnselectedH,
            assassinUnselected, assassinUnselectedH,
            druidUnselected, druidUnselectedH;
        private IFont headingFont;
        private ILabel headingLabel;
        private Button exitButton;

        public SelectHeroClass(
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

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBackground, Palettes.Fechar);
            campfireSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectCampfire, Palettes.Fechar, new System.Drawing.Point(380, 335));

            barbarianUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselected, Palettes.Fechar, new System.Drawing.Point(400, 330));
            barbarianUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselectedH, Palettes.Fechar, new System.Drawing.Point(400, 330));

            sorceressUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselected, Palettes.Fechar, new System.Drawing.Point(626, 352));
            sorceressUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselectedH, Palettes.Fechar, new System.Drawing.Point(626, 352));

            necromancerUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselected, Palettes.Fechar, new System.Drawing.Point(300, 335));
            necromancerUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselectedH, Palettes.Fechar, new System.Drawing.Point(300, 335));

            paladinUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselected, Palettes.Fechar, new System.Drawing.Point(521, 338));
            paladinUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselectedH, Palettes.Fechar, new System.Drawing.Point(521, 338));

            amazonUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselected, Palettes.Fechar, new System.Drawing.Point(100, 339));
            amazonUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselectedH, Palettes.Fechar, new System.Drawing.Point(100, 339));

            assassinUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselected, Palettes.Fechar, new System.Drawing.Point(231, 365));
            assassinUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselectedH, Palettes.Fechar, new System.Drawing.Point(231, 365));

            druidUnselected = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselected, Palettes.Fechar, new System.Drawing.Point(720, 370));
            druidUnselectedH = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselectedH, Palettes.Fechar, new System.Drawing.Point(720, 370));

            headingFont = renderWindow.LoadFont(ResourcePaths.Font30, Palettes.Units);
            headingLabel = renderWindow.CreateLabel(headingFont);
            headingLabel.Text = "Select Hero Class";
            headingLabel.Location = new System.Drawing.Point(400 - (headingLabel.TextArea.Width / 2), 17);

            exitButton = createButton(eButtonType.Medium);
            exitButton.Text = "EXIT";
            exitButton.Location = new System.Drawing.Point(30, 540);
            exitButton.OnActivate = OnExitClicked;
        }

        private void OnExitClicked()
        {
            sceneManager.ChangeScene("Main Menu");
        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);

            renderWindow.Draw(barbarianUnselected, (int)(barbarianUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(sorceressUnselected, (int)(sorceressUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(necromancerUnselected, (int)(necromancerUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(paladinUnselected, (int)(paladinUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(amazonUnselected, (int)(amazonUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(assassinUnselected, (int)(assassinUnselected.TotalFrames * secondTimer));
            renderWindow.Draw(druidUnselected, (int)(druidUnselected.TotalFrames * secondTimer));

            renderWindow.Draw(campfireSprite, (int)(campfireSprite.TotalFrames * secondTimer));
            renderWindow.Draw(headingLabel);
            exitButton.Render();
        }

        public void Update(long ms)
        {
            float seconds = ((float)ms / 1000f);
            secondTimer += seconds;
            while (secondTimer >= 1f)
                secondTimer -= 1f;

            exitButton.Update();
        }

        public void Dispose()
        {
            backgroundSprite.Dispose();
            campfireSprite.Dispose();
            headingFont.Dispose();
            headingLabel.Dispose();
        }
    }
}
