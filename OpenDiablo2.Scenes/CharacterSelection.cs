using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Scenes
{
    [Scene("Select Character")]
    public sealed class CharacterSelection : IScene
    {
        static readonly log4net.ILog log =
            log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;

        private readonly ISprite backgroundSprite;
        private readonly IButton createNewCharacterButton, deleteCharacterButton, exitButton, okButton;

        public CharacterSelection(IRenderWindow renderWindow,
            ISceneManager sceneManager, ITextDictionary textDictionary, Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CharacterSelectionBackground, Palettes.Sky);
            createNewCharacterButton = createButton(eButtonType.Tall);
            // TODO: use strCreateNewCharacter -- need to get the text to split after 10 chars though.
            createNewCharacterButton.Text = textDictionary.Translate("strCreateNewCharacter");// "Create New".ToUpper();
            createNewCharacterButton.Location = new Point(33, 467);
            createNewCharacterButton.OnActivate = () => sceneManager.ChangeScene("Select Hero Class");

            deleteCharacterButton = createButton(eButtonType.Tall);
            deleteCharacterButton.Text = textDictionary.Translate("strDelete");
            deleteCharacterButton.Location = new Point(433, 467);

            exitButton = createButton(eButtonType.Medium);
            exitButton.Text = textDictionary.Translate("strExit");
            exitButton.Location = new Point(33, 540);
            exitButton.OnActivate = () => sceneManager.ChangeScene("Main Menu");

            okButton = createButton(eButtonType.Medium);
            okButton.Text = textDictionary.Translate("strOk");
            okButton.Location = new Point(630, 540);
            okButton.Enabled = false;
        }

        public void Update(long ms)
        {
            createNewCharacterButton.Update();
            deleteCharacterButton.Update();
            exitButton.Update();
            okButton.Update();
        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);

            createNewCharacterButton.Render();
            deleteCharacterButton.Render();
            exitButton.Render();
            okButton.Render();
        }

        public void Dispose()
        {
            backgroundSprite.Dispose();
            createNewCharacterButton.Dispose();
            deleteCharacterButton.Dispose();
            exitButton.Dispose();
            okButton.Dispose();
        }
    }
}