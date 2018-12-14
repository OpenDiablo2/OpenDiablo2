using OpenDiablo2.Common;
using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Scenes
{
    [Scene(eSceneType.Credits)]
    public sealed class Credits : IScene
    {
        private class LabelItem
        {
            public ILabel Label { get; set; }
            public bool IsHeading { get; set; }
            public bool Avaiable { get; set; }
        }

        private const float secondsPerCycle = (40f / 1000f);

        private readonly IRenderWindow renderWindow;
        private readonly ISceneManager sceneManager;
        private readonly IMPQProvider mpqProvider;

        private bool doneWithCredits = false;
        private int cyclesTillNextLine = 0;
        private float cycleTime = 0f;
        private IFont textFont;
        private ISprite backgroundSprite;
        private IButton btnExit;
        private Stack<string> creditsText;
        private List<LabelItem> fontLabels = new List<LabelItem>();

        public Credits(
            IRenderWindow renderWindow,
            ISceneManager sceneManager,
            IMPQProvider mpqProvider,
            Func<eButtonType, IButton> createButton
            )
        {
            this.renderWindow = renderWindow;
            this.sceneManager = sceneManager;
            this.mpqProvider = mpqProvider;

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.CreditsBackground, Palettes.Sky);

            btnExit = createButton(eButtonType.Medium);
            btnExit.Text = "Exit".ToUpper();
            btnExit.Location = new Point(20, 550);
            btnExit.OnActivate = OnActivateClicked;

            textFont = renderWindow.LoadFont(ResourcePaths.FontFormal10, Palettes.Static);

            creditsText = new Stack<string>((
                File.ReadAllLines(Path.Combine(Path.GetDirectoryName(Environment.GetCommandLineArgs().First()), "credits.txt"))
                    .Concat(mpqProvider.GetTextFile(ResourcePaths.CreditsText))
                ).Reverse());
        }

        private void OnActivateClicked()
            => sceneManager.ChangeScene(eSceneType.MainMenu);

        private void AddNextItem()
        {
            if (!creditsText.Any())
            {
                doneWithCredits = true;
                return;
            }

            var text = creditsText.Pop().Trim();
            if (text.Trim().Length == 0)
            {
                cyclesTillNextLine = 18;
                return;
            }
            var isHeading = text.StartsWith("*");
            var isNextHeading = creditsText.Any() && creditsText.Peek().StartsWith("*");
            var isNextSpace = creditsText.Any() && creditsText.Peek().Trim().Length == 0;
            var label = GetNewFontLabel(isHeading);
            label.Text = isHeading ? text.Substring(1) : text;
            var isDoubled = false;
            if (!isHeading && !isNextHeading && !isNextSpace)
            {
                isDoubled = true;

                // Gotta go side by side
                label.Location = new Point(390 - label.TextArea.Width, 605);

                var text2 = creditsText.Pop().Trim();
                isNextHeading = creditsText.Any() && creditsText.Peek().StartsWith("*");
                var label2 = GetNewFontLabel(isHeading);
                label2.Text = text2;

                label2.Location = new Point(410, 605);
            }
            else
            {
                label.Location = new Point(400 - (label.TextArea.Width / 2), 605);
            }

            if (isHeading && isNextHeading)
                cyclesTillNextLine = 40;
            else if (isNextHeading)
                cyclesTillNextLine = isDoubled ? 40 : 70;
            else if (isHeading)
                cyclesTillNextLine = 40;
            else
                cyclesTillNextLine = 18;
        }

        public void Render()
        {
            renderWindow.Draw(backgroundSprite, 4, 3, 0);
            btnExit.Render();
            foreach (var label in fontLabels.Where(x => !x.Avaiable).Select(x => x.Label))
                renderWindow.Draw(label);
        }

        public void Update(long ms)
        {

            cycleTime += (ms / 1000f);
            while (cycleTime >= secondsPerCycle)
            {
                cycleTime -= secondsPerCycle;
                if (!doneWithCredits && (--cyclesTillNextLine <= 0))
                    AddNextItem();

                foreach (var fontLabel in fontLabels.Where(x => !x.Avaiable))
                {
                    if (fontLabel.Label.Location.Y - 1 <= -15)
                    {
                        fontLabel.Avaiable = true;
                        continue;
                    }
                    fontLabel.Label.Location = new Point(fontLabel.Label.Location.X, fontLabel.Label.Location.Y - 1);
                }
            }

            btnExit.Update();
        }

        private ILabel GetNewFontLabel(bool isHeading)
        {
            var labelItem = fontLabels.FirstOrDefault(x => x.Avaiable && x.IsHeading == isHeading);
            if (labelItem != null)
            {
                labelItem.Avaiable = false;
                return labelItem.Label;
            }

            var newLabelItem = new LabelItem
            {
                Avaiable = false,
                IsHeading = isHeading,
                Label = renderWindow.CreateLabel(textFont)
            };

            newLabelItem.Label.TextColor = isHeading
                ? Color.FromArgb(255, 88, 82)
                : Color.FromArgb(198, 178, 150);

            fontLabels.Add(newLabelItem);
            return newLabelItem.Label;
        }


        public void Dispose()
        {
            backgroundSprite?.Dispose();

            foreach (var labelItem in fontLabels)
                labelItem.Label.Dispose();

            textFont.Dispose();
        }
    }
}
