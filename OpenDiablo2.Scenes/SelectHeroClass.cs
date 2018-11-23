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
    enum eHero
    {
        Amazon,
        Assassin,
        Necromancer,
        Barbarian,
        Paladin,
        Sorceress,
        Druid
    }

    enum eHeroStance
    {
        Idle,
        IdleSelected,
        Approaching,
        Selected,
        Retreating
    }

    struct HeroRenderInfo
    {
        public ISprite IdleSprite, IdleSelectedSprite, ApproacingSprite, SelectedSprite, RetreatingSprite;
        public eHeroStance Stance;
    }

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
        private Button exitButton;
        private Dictionary<eHero, HeroRenderInfo> heroRenderInfo = new Dictionary<eHero, HeroRenderInfo>();

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

            heroRenderInfo[eHero.Barbarian] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselected, Palettes.Fechar, new System.Drawing.Point(400, 330)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselectedH, Palettes.Fechar, new System.Drawing.Point(400, 330))
            };

            heroRenderInfo[eHero.Sorceress] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselected, Palettes.Fechar, new System.Drawing.Point(626, 352)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselectedH, Palettes.Fechar, new System.Drawing.Point(626, 352))
            };

            heroRenderInfo[eHero.Necromancer] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselected, Palettes.Fechar, new System.Drawing.Point(300, 335)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselectedH, Palettes.Fechar, new System.Drawing.Point(300, 335))
            };

            heroRenderInfo[eHero.Paladin] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselected, Palettes.Fechar, new System.Drawing.Point(521, 338)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselectedH, Palettes.Fechar, new System.Drawing.Point(521, 338))
            };

            heroRenderInfo[eHero.Amazon] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselected, Palettes.Fechar, new System.Drawing.Point(100, 339)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselectedH, Palettes.Fechar, new System.Drawing.Point(100, 339))
            };

            heroRenderInfo[eHero.Assassin] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselected, Palettes.Fechar, new System.Drawing.Point(231, 365)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselectedH, Palettes.Fechar, new System.Drawing.Point(231, 365))
            };

            heroRenderInfo[eHero.Druid] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselected, Palettes.Fechar, new System.Drawing.Point(720, 370)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselectedH, Palettes.Fechar, new System.Drawing.Point(720, 370))
            };

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

            RenderHeros();

            renderWindow.Draw(campfireSprite, (int)(campfireSprite.TotalFrames * secondTimer));
            renderWindow.Draw(headingLabel);
            exitButton.Render();
        }

        private void RenderHeros()
        {
            foreach (var hero in Enum.GetValues(typeof(eHero)).Cast<eHero>())
                RenderHero(hero);

        }

        private void RenderHero(eHero hero)
        {
            var renderInfo = heroRenderInfo[hero];
            switch (renderInfo.Stance)
            {
                case eHeroStance.Idle:
                    renderWindow.Draw(renderInfo.IdleSprite, (int)(renderInfo.IdleSprite.TotalFrames * secondTimer));
                    break;
                case eHeroStance.IdleSelected:
                    renderWindow.Draw(renderInfo.IdleSelectedSprite, (int)(renderInfo.IdleSelectedSprite.TotalFrames * secondTimer));
                    break;
                case eHeroStance.Approaching:
                    break;
                case eHeroStance.Selected:
                    break;
                case eHeroStance.Retreating:
                    break;
                default:
                    break;
            }

        }

        public void Update(long ms)
        {
            float seconds = ((float)ms / 1500f);
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
