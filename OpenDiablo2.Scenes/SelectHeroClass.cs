using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{

    enum eHeroStance
    {
        Idle,
        IdleSelected,
        Approaching,
        Selected,
        Retreating
    }

    class HeroRenderInfo : IDisposable
    {
        public ISprite IdleSprite, IdleSelectedSprite, ForwardWalkSprite, ForwardWalkSpriteOverlay, SelectedSprite, SelectedSpriteOverlay, BackWalkSprite, BackWalkSpriteOverlay;
        public eHeroStance Stance;
        public long ForwardWalkTimeMs, BackWalkTimeMs;
        public long SpecialFrameTime;
        public Rectangle SelectionBounds = new Rectangle();

        public void Dispose()
        {
            IdleSprite?.Dispose();
            IdleSelectedSprite?.Dispose();
            ForwardWalkSprite?.Dispose();
            ForwardWalkSpriteOverlay?.Dispose();
            SelectedSprite?.Dispose();
            SelectedSpriteOverlay?.Dispose();
            BackWalkSprite?.Dispose();
            BackWalkSpriteOverlay?.Dispose();
        }
    }

    [Scene(eSceneType.SelectHeroClass)]
    public sealed class SelectHeroClass : IScene
    {
        private readonly IRenderWindow renderWindow;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly ISceneManager sceneManager;
        private readonly ITextDictionary textDictionary;
        private readonly IKeyboardInfoProvider keyboardInfoProvider;
        private readonly IGameState gameState;
        private readonly ISoundProvider soundProvider;

        private bool showEntryUi = false;
        private eHero? selectedHero = null;
        private float secondTimer;
        private int sfxChannel = -1;
        private int sfxChannel2 = -1;
        private readonly ISprite backgroundSprite, campfireSprite;
        private readonly IFont headingFont;
        private readonly IFont heroDescFont;
        private readonly IFont uiFont;
        private readonly ILabel headingLabel, heroClassLabel, heroDesc1Label, heroDesc2Label, heroDesc3Label, characterNameLabel;
        private readonly IButton exitButton, okButton;
        private readonly ITextBox characterNameTextBox;
        private readonly Dictionary<eHero, HeroRenderInfo> heroRenderInfo = new Dictionary<eHero, HeroRenderInfo>();

        private Dictionary<string, byte[]> sfxDictionary;

        public SelectHeroClass(
            IRenderWindow renderWindow,
            IMouseInfoProvider mouseInfoProvider,
            ISceneManager sceneManager,
            ISoundProvider soundProvider,
            Func<eButtonType, IButton> createButton,
            Func<ITextBox> createTextBox,
            ITextDictionary textDictionary,
            IKeyboardInfoProvider keyboardInfoProvider,
            IMPQProvider mpqProvider,
            IGameState gameState
            )
        {
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sceneManager = sceneManager;
            this.textDictionary = textDictionary;
            this.keyboardInfoProvider = keyboardInfoProvider;
            this.soundProvider = soundProvider;
            this.gameState = gameState;
            sfxDictionary = new Dictionary<string, byte[]>();

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBackground, Palettes.Fechar);
            campfireSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectCampfire, Palettes.Fechar, new Point(380, 335));
            campfireSprite.Blend = true;

            heroRenderInfo[eHero.Barbarian] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselected, Palettes.Fechar, new Point(400, 330)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianUnselectedH, Palettes.Fechar, new Point(400, 330)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianForwardWalk, Palettes.Fechar, new Point(400, 330)),
                ForwardWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianForwardWalkOverlay, Palettes.Fechar, new Point(400, 330)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianSelected, Palettes.Fechar, new Point(400, 330)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectBarbarianBackWalk, Palettes.Fechar, new Point(400, 330)),
                SelectionBounds = new Rectangle(364, 201, 90, 170),
                ForwardWalkTimeMs = 2500,
                BackWalkTimeMs = 1000
            };

            heroRenderInfo[eHero.Sorceress] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselected, Palettes.Fechar, new Point(626, 352)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressUnselectedH, Palettes.Fechar, new Point(626, 352)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressForwardWalk, Palettes.Fechar, new Point(626, 352)),
                ForwardWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressForwardWalkOverlay, Palettes.Fechar, new Point(626, 352)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressSelected, Palettes.Fechar, new Point(626, 352)),
                SelectedSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressSelectedOverlay, Palettes.Fechar, new Point(626, 352)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressBackWalk, Palettes.Fechar, new Point(626, 352)),
                BackWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecSorceressBackWalkOverlay, Palettes.Fechar, new Point(626, 352)),
                SelectionBounds = new Rectangle(580, 240, 65, 160),
                ForwardWalkTimeMs = 2300,
                BackWalkTimeMs = 1200
            };
            heroRenderInfo[eHero.Sorceress].SelectedSpriteOverlay.Blend = true;
            heroRenderInfo[eHero.Sorceress].ForwardWalkSpriteOverlay.Blend = true;
            heroRenderInfo[eHero.Sorceress].BackWalkSpriteOverlay.Blend = true;


            heroRenderInfo[eHero.Necromancer] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselected, Palettes.Fechar, new Point(300, 335)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectNecromancerUnselectedH, Palettes.Fechar, new Point(300, 335)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerForwardWalk, Palettes.Fechar, new Point(300, 335)),
                ForwardWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerForwardWalkOverlay, Palettes.Fechar, new Point(300, 335)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerSelected, Palettes.Fechar, new Point(300, 335)),
                SelectedSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerSelectedOverlay, Palettes.Fechar, new Point(300, 335)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerBackWalk, Palettes.Fechar, new Point(300, 335)),
                BackWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecNecromancerBackWalkOverlay, Palettes.Fechar, new Point(300, 335)),
                SelectionBounds = new Rectangle(265, 220, 55, 175),
                ForwardWalkTimeMs = 2000,
                BackWalkTimeMs = 1500,
            };
            heroRenderInfo[eHero.Necromancer].ForwardWalkSpriteOverlay.Blend = true;
            heroRenderInfo[eHero.Necromancer].BackWalkSpriteOverlay.Blend = true;
            heroRenderInfo[eHero.Necromancer].SelectedSpriteOverlay.Blend = true;

            heroRenderInfo[eHero.Paladin] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselected, Palettes.Fechar, new Point(521, 338)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectPaladinUnselectedH, Palettes.Fechar, new Point(521, 338)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecPaladinForwardWalk, Palettes.Fechar, new Point(521, 338)),
                ForwardWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecPaladinForwardWalkOverlay, Palettes.Fechar, new Point(521, 338)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecPaladinSelected, Palettes.Fechar, new Point(521, 338)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecPaladinBackWalk, Palettes.Fechar, new Point(521, 338)),
                SelectionBounds = new Rectangle(490, 210, 65, 180),
                ForwardWalkTimeMs = 3400,
                BackWalkTimeMs = 1300
            };

            heroRenderInfo[eHero.Amazon] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselected, Palettes.Fechar, new Point(100, 339)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAmazonUnselectedH, Palettes.Fechar, new Point(100, 339)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecAmazonForwardWalk, Palettes.Fechar, new Point(100, 339)),
                //ForwardWalkSpriteOverlay = renderWindow.LoadSprite(ResourcePaths.CharacterSelecAmazonForwardWalkOverlay, Palettes.Fechar, new Point(100, 339)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecAmazonSelected, Palettes.Fechar, new Point(100, 339)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelecAmazonBackWalk, Palettes.Fechar, new Point(100, 339)),
                SelectionBounds = new Rectangle(70, 220, 55, 200),
                ForwardWalkTimeMs = 2200,
                BackWalkTimeMs = 1500
            };

            heroRenderInfo[eHero.Assassin] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselected, Palettes.Fechar, new Point(231, 365)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinUnselectedH, Palettes.Fechar, new Point(231, 365)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinForwardWalk, Palettes.Fechar, new Point(231, 365)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinSelected, Palettes.Fechar, new Point(231, 365)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectAssassinBackWalk, Palettes.Fechar, new Point(231, 365)),
                SelectionBounds = new Rectangle(175, 235, 50, 180),
                ForwardWalkTimeMs = 3800,
                BackWalkTimeMs = 1500
            };

            heroRenderInfo[eHero.Druid] = new HeroRenderInfo
            {
                Stance = eHeroStance.Idle,
                IdleSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselected, Palettes.Fechar, new Point(720, 370)),
                IdleSelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidUnselectedH, Palettes.Fechar, new Point(720, 370)),
                ForwardWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidForwardWalk, Palettes.Fechar, new Point(720, 370)),
                SelectedSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidSelected, Palettes.Fechar, new Point(720, 370)),
                BackWalkSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectDruidBackWalk, Palettes.Fechar, new Point(720, 370)),
                SelectionBounds = new Rectangle(680, 220, 70, 195),
                ForwardWalkTimeMs = 4800,
                BackWalkTimeMs = 1500
            };

            headingFont = renderWindow.LoadFont(ResourcePaths.Font30, Palettes.Units);
            heroDescFont = renderWindow.LoadFont(ResourcePaths.Font16, Palettes.Units);
            uiFont = renderWindow.LoadFont(ResourcePaths.Font16, Palettes.Units);

            headingLabel = renderWindow.CreateLabel(headingFont);
            headingLabel.Text = textDictionary.Translate("strSelectHeroClass");
            headingLabel.Location = new Point(400 - (headingLabel.TextArea.Width / 2), 17);

            heroClassLabel = renderWindow.CreateLabel(headingFont);
            heroClassLabel.Text = "";
            heroClassLabel.Location = new Point(400 - (heroClassLabel.TextArea.Width / 2), 65);

            heroDesc1Label = renderWindow.CreateLabel(heroDescFont);
            heroDesc2Label = renderWindow.CreateLabel(heroDescFont);
            heroDesc3Label = renderWindow.CreateLabel(heroDescFont);

            characterNameLabel = renderWindow.CreateLabel(uiFont);
            characterNameLabel.Text = textDictionary.Translate("strCharacterName");
            characterNameLabel.Location = new Point(320, 475);
            characterNameLabel.TextColor = Color.FromArgb(216, 196, 128);

            exitButton = createButton(eButtonType.Medium);
            exitButton.Text = textDictionary.Translate("strExit");
            exitButton.Location = new Point(33, 540);
            exitButton.OnActivate = OnExitClicked;

            okButton = createButton(eButtonType.Medium);
            okButton.Text = textDictionary.Translate("strOk");
            okButton.Location = new Point(630, 540);
            okButton.OnActivate = OnOkclicked;
            okButton.Enabled = false;

            characterNameTextBox = createTextBox();
            characterNameTextBox.Text = "";
            characterNameTextBox.Location = new Point(320, 493);

            Parallel.ForEach(new[]
            {
                ResourcePaths.SFXAmazonSelect,
                ResourcePaths.SFXAssassinSelect,
                ResourcePaths.SFXBarbarianSelect,
                ResourcePaths.SFXDruidSelect,
                ResourcePaths.SFXNecromancerSelect,
                ResourcePaths.SFXPaladinSelect,
                ResourcePaths.SFXSorceressSelect,

                ResourcePaths.SFXAmazonDeselect,
                ResourcePaths.SFXAssassinDeselect,
                ResourcePaths.SFXBarbarianDeselect,
                ResourcePaths.SFXDruidDeselect,
                ResourcePaths.SFXNecromancerDeselect,
                ResourcePaths.SFXPaladinDeselect,
                ResourcePaths.SFXSorceressDeselect
            }, (path => sfxDictionary.Add(path, mpqProvider.GetBytes(path))));
        }

        private void StopSfx()
        {
            if (sfxChannel > -1)
                soundProvider.StopSfx(sfxChannel);

            if (sfxChannel2 > -1)
                soundProvider.StopSfx(sfxChannel);
        }

        private void OnOkclicked()
        {
            StopSfx();

            // TODO: Support other session types
            // TODO: support other difficulty types
            gameState.Initialize(characterNameTextBox.Text, selectedHero.Value, eSessionType.Local, eDifficulty.NORMAL);
        }

        private void OnExitClicked()
        {
            StopSfx();

            var heros = Enum.GetValues(typeof(eHero)).Cast<eHero>();
            foreach (var hero in heros)
            {
                heroRenderInfo[hero].SpecialFrameTime = 0;
                heroRenderInfo[hero].Stance = eHeroStance.Idle;
            }
            showEntryUi = false;
            keyboardInfoProvider.KeyPressCallback = null;
            characterNameTextBox.Text = "";
            okButton.Enabled = false;
            selectedHero = null;

            sceneManager.ChangeScene(eSceneType.SelectCharacter);
        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);

            RenderHeros();

            renderWindow.Draw(campfireSprite, (int)(campfireSprite.TotalFrames * secondTimer));
            renderWindow.Draw(headingLabel);
            if (selectedHero.HasValue)
            {
                renderWindow.Draw(heroClassLabel);
                renderWindow.Draw(heroDesc1Label);
                renderWindow.Draw(heroDesc2Label);
                renderWindow.Draw(heroDesc3Label);
            }

            exitButton.Render();

            if (showEntryUi)
            {
                renderWindow.Draw(characterNameLabel);
                okButton.Render();
                characterNameTextBox.Render();
            }
        }

        private void RenderHeros()
        {
            var heros = Enum.GetValues(typeof(eHero)).Cast<eHero>().Skip(1); // skip NONE
            foreach (var hero in heros)
                if (heroRenderInfo[hero].Stance == eHeroStance.Idle || heroRenderInfo[hero].Stance == eHeroStance.IdleSelected)
                    RenderHero(hero);

            foreach (var hero in heros)
                if (heroRenderInfo[hero].Stance != eHeroStance.Idle && heroRenderInfo[hero].Stance != eHeroStance.IdleSelected)
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
                    {
                        var framePct = renderInfo.SpecialFrameTime / (float)renderInfo.ForwardWalkTimeMs;
                        renderWindow.Draw(renderInfo.ForwardWalkSprite, (int)(renderInfo.ForwardWalkSprite.TotalFrames * framePct));
                        if (renderInfo.ForwardWalkSpriteOverlay != null)
                            renderWindow.Draw(renderInfo.ForwardWalkSpriteOverlay, (int)(renderInfo.ForwardWalkSpriteOverlay.TotalFrames * framePct));
                    }
                    break;
                case eHeroStance.Selected:
                    {
                        var framePct = renderInfo.SpecialFrameTime / (float)1000;
                        renderWindow.Draw(renderInfo.SelectedSprite, (int)(renderInfo.SelectedSprite.TotalFrames * framePct));
                        if (renderInfo.SelectedSpriteOverlay != null)
                            renderWindow.Draw(renderInfo.SelectedSpriteOverlay, (int)(renderInfo.SelectedSpriteOverlay.TotalFrames * framePct));
                    }
                    break;
                case eHeroStance.Retreating:
                    {
                        var framePct = renderInfo.SpecialFrameTime / (float)renderInfo.BackWalkTimeMs;
                        renderWindow.Draw(renderInfo.BackWalkSprite, (int)(renderInfo.BackWalkSprite.TotalFrames * framePct));
                        if (renderInfo.BackWalkSpriteOverlay != null)
                            renderWindow.Draw(renderInfo.BackWalkSpriteOverlay, (int)(renderInfo.BackWalkSpriteOverlay.TotalFrames * framePct));
                    }
                    break;
            }

        }

        private void OnKeyPressed(char charcode)
        {
            if (charcode == '\b')
            {
                if (characterNameTextBox.Text.Length == 0)
                    return;

                characterNameTextBox.Text = characterNameTextBox.Text.Substring(0, characterNameTextBox.Text.Length - 1);
                okButton.Enabled = characterNameTextBox.Text.Length >= 2;
                return;
            }

            if (characterNameTextBox.Text.Length >= 15 || !selectedHero.HasValue)
                return;

            if (!"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".Contains(charcode))
                return;
            characterNameTextBox.Text += charcode;
            okButton.Enabled = characterNameTextBox.Text.Length >= 2;
        }

        public void Update(long ms)
        {
            float seconds = ms / 1500f;
            secondTimer += seconds;
            while (secondTimer >= 1f)
                secondTimer -= 1f;

            if (keyboardInfoProvider.KeyPressCallback == null)
                keyboardInfoProvider.KeyPressCallback = OnKeyPressed;

            // Don't update hero selection if one of them is walking to or from the campfire
            var canSelect = heroRenderInfo.All(x => x.Value.Stance == eHeroStance.Idle || x.Value.Stance == eHeroStance.IdleSelected || x.Value.Stance == eHeroStance.Selected);

            foreach (var hero in Enum.GetValues(typeof(eHero)).Cast<eHero>())
                UpdateHeroSelectionHover(hero, ms, canSelect);

            if (selectedHero.HasValue && heroRenderInfo.All(x => x.Value.Stance == eHeroStance.Idle))
            {
                selectedHero = null;
            }

            exitButton.Update();
            okButton.Update();
            characterNameTextBox.Update(ms);
        }

        private void UpdateHeroSelectionHover(eHero hero, long ms, bool canSelect)
        {
            if(hero == eHero.None)
            {
                return;
            }
            var renderInfo = heroRenderInfo[hero];
            if (renderInfo.Stance == eHeroStance.Approaching)
            {
                renderInfo.SpecialFrameTime += ms;
                if (renderInfo.SpecialFrameTime >= renderInfo.ForwardWalkTimeMs)
                {
                    renderInfo.Stance = eHeroStance.Selected;
                    renderInfo.SpecialFrameTime = 0;
                }

                return;
            }

            if (renderInfo.Stance == eHeroStance.Retreating)
            {
                renderInfo.SpecialFrameTime += ms;
                if (renderInfo.SpecialFrameTime >= renderInfo.BackWalkTimeMs)
                {
                    renderInfo.Stance = eHeroStance.Idle;
                    renderInfo.SpecialFrameTime = 0;
                }

                return;
            }

            if (renderInfo.Stance == eHeroStance.Selected)
            {
                renderInfo.SpecialFrameTime += ms;
                while (renderInfo.SpecialFrameTime >= 1000)
                    renderInfo.SpecialFrameTime -= 1000;
                return;
            }

            if (!canSelect)
                return;

            // No need to highlight a hero if they are next to the campfire
            if (renderInfo.Stance == eHeroStance.Selected)
                return;

            var mouseX = mouseInfoProvider.MouseX;
            var mouseY = mouseInfoProvider.MouseY;

            var b = renderInfo.SelectionBounds;
            var mouseHover = (mouseX >= b.Left) && (mouseX <= b.Left + b.Width) && (mouseY >= b.Top) && (mouseY <= b.Top + b.Height);

            if (mouseHover && mouseInfoProvider.LeftMouseDown)
            {
                showEntryUi = true;
                renderInfo.Stance = eHeroStance.Approaching;
                renderInfo.SpecialFrameTime = 0;


                foreach (var ri in heroRenderInfo)
                {
                    if (ri.Value.Stance != eHeroStance.Selected)
                        continue;

                    PlayHeroDeselected(ri.Key);
                    ri.Value.Stance = eHeroStance.Retreating;
                    ri.Value.SpecialFrameTime = 0;
                    break;
                }

                selectedHero = hero;
                UpdateHeroText();
                PlayHeroSelected(hero);

                return;
            }

            heroRenderInfo[hero].Stance = mouseHover ? eHeroStance.IdleSelected : eHeroStance.Idle;

            if (selectedHero == null && mouseHover)
            {
                selectedHero = hero;
                UpdateHeroText();
            }

        }

        private void PlayHeroSelected(eHero hero)
        {
            switch (hero)
            {
                case eHero.Barbarian:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXBarbarianSelect]);
                    break;
                case eHero.Necromancer:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXNecromancerSelect]);
                    break;
                case eHero.Paladin:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXPaladinSelect]);
                    break;
                case eHero.Assassin:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXAssassinSelect]);
                    break;
                case eHero.Sorceress:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXSorceressSelect]);
                    break;
                case eHero.Amazon:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXAmazonSelect]);
                    break;
                case eHero.Druid:
                    sfxChannel = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXDruidSelect]);
                    break;
                default:
                    break;
            }
        }

        private void PlayHeroDeselected(eHero hero)
        {
            switch (hero)
            {
                case eHero.Barbarian:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXBarbarianDeselect]);
                    break;
                case eHero.Necromancer:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXNecromancerDeselect]);
                    break;
                case eHero.Paladin:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXPaladinDeselect]);
                    break;
                case eHero.Assassin:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXAssassinDeselect]);
                    break;
                case eHero.Sorceress:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXSorceressDeselect]);
                    break;
                case eHero.Amazon:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXAmazonDeselect]);
                    break;
                case eHero.Druid:
                    sfxChannel2 = soundProvider.PlaySfx(sfxDictionary[ResourcePaths.SFXDruidDeselect]);
                    break;
                default:
                    break;
            }
        }

        private void SetDescLabels(string descKey)
        {
            var heroDesc = textDictionary.Translate(descKey);
            var parts = StringUtils.SplitIntoLinesWithMaxWidth(heroDesc, 37);
            heroDesc1Label.Text = parts.Count > 0 ? parts[0] : "";
            heroDesc2Label.Text = parts.Count > 1 ? parts[1] : "";
            heroDesc3Label.Text = parts.Count > 2 ? parts[2] : "";
        }

        private void UpdateHeroText()
        {
            if (selectedHero == null)
                return;

            switch (selectedHero.Value)
            {
                case eHero.Barbarian:
                    heroClassLabel.Text = textDictionary.Translate("strBarbarian");
                    SetDescLabels("strBarbDesc");
                    break;
                case eHero.Necromancer:
                    heroClassLabel.Text = textDictionary.Translate("strNecromancer");
                    SetDescLabels("strNecroDesc");
                    break;
                case eHero.Paladin:
                    heroClassLabel.Text = textDictionary.Translate("strPaladin");
                    SetDescLabels("strPalDesc");
                    break;
                case eHero.Assassin:
                    heroClassLabel.Text = textDictionary.Translate("strAssassin");
                    heroDesc1Label.Text = "Schooled in the Material Arts. Her";
                    heroDesc2Label.Text = "mind and body are deadly weapons.";
                    heroDesc3Label.Text = "";
                    break;
                case eHero.Sorceress:
                    heroClassLabel.Text = textDictionary.Translate("strSorceress");
                    SetDescLabels("strSorcDesc");
                    break;
                case eHero.Amazon:
                    heroClassLabel.Text = textDictionary.Translate("strAmazon");
                    SetDescLabels("strAmazonDesc");
                    break;
                case eHero.Druid:
                    heroClassLabel.Text = textDictionary.Translate("strDruid");
                    heroDesc1Label.Text = "Commanding the forces of nature, he";
                    heroDesc2Label.Text = "summons wild beasts and raging";
                    heroDesc3Label.Text = "storms to his side.";
                    break;
            }

            heroClassLabel.Location = new Point(400 - (heroClassLabel.TextArea.Width / 2), 65);
            heroDesc1Label.Location = new Point(400 - (heroDesc1Label.TextArea.Width / 2), 100);
            heroDesc2Label.Location = new Point(400 - (heroDesc2Label.TextArea.Width / 2), 115);
            heroDesc3Label.Location = new Point(400 - (heroDesc3Label.TextArea.Width / 2), 130);
        }

        public void Dispose()
        {
            backgroundSprite.Dispose();
            campfireSprite.Dispose();
            headingFont.Dispose();
            headingLabel.Dispose();
            sfxDictionary.Clear();

            foreach (var hri in heroRenderInfo)
                hri.Value.Dispose();

            heroRenderInfo.Clear();
        }

    }
}
