using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;

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
        private ISprite backgroundSprite, campfireSprite;
        private IFont headingFont;
        private ILabel headingLabel;

        public SelectHeroClass(
            IRenderWindow renderWindow,
            IPaletteProvider paletteProvider,
            IMPQProvider mpqProvider,
            IMouseInfoProvider mouseInfoProvider,
            IMusicProvider musicProvider,
            ISceneManager sceneManager
            )
        {
            this.renderWindow = renderWindow;
            this.paletteProvider = paletteProvider;
            this.mpqProvider = mpqProvider;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sceneManager = sceneManager;

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBackground, Palettes.Fechar);
            campfireSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectCampfire, Palettes.Fechar, new System.Drawing.Point(380, 335));

            headingFont = renderWindow.LoadFont(ResourcePaths.Font30, Palettes.Units);
            headingLabel = renderWindow.CreateLabel(headingFont);
            headingLabel.Text = "Select Hero Class";
            headingLabel.Location = new System.Drawing.Point(400 - (headingLabel.TextArea.Width / 2), 20);


        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);
            renderWindow.Draw(campfireSprite, (int)(campfireSprite.TotalFrames * secondTimer));
            //renderWindow.Draw(headingLabel);
        }

        public void Update(long ms)
        {
            float seconds = ((float)ms / 1000f);
            secondTimer += seconds;
            while (secondTimer >= 1f)
                secondTimer -= 1f;

        }

        public void Dispose()
        {
            
        }
    }
}
