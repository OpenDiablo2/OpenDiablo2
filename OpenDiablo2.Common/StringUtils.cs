using System.Collections.Generic;
using System.Text;

namespace OpenDiablo2.Common
{
    public static class StringUtils
    {
        public static List<string> SplitIntoLinesWithMaxWidth(string fullSentence, int maxChars)
        {
            var lines = new List<string>();
            var line = new StringBuilder();
            var totalLength = 0;
            var words = fullSentence.Split(' ');
            foreach (var word in words)
            {
                totalLength += 1 + word.Length;
                if (totalLength > maxChars)
                {
                    totalLength = word.Length;
                    lines.Add(line.ToString());
                    line = new StringBuilder();
                }
                else
                {
                    line.Append(' ');
                }

                line.Append(word);
            }

            if (line.Length > 0)
            {
                lines.Add(line.ToString());
            }

            return lines;
        }
    }
}