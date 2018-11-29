using CommandLine;

namespace OpenDiablo2
{
    public sealed class CommandLineOptions
    {
        [Option('p', "datapath", Required = false, HelpText = "Specifies the root data path")]
        public string DataPath { get; set; }

        [Option("hwmouse", Default = false, Required = false, HelpText = "Use the hardware mouse instead of software")]
        public bool HardwareMouse { get; set; }

        [Option("mousescale", Default = 1, Required = false, HelpText = "When hardware mouse is enabled, this defines the pixel scale of the mouse. No effect for software mode")]
        public int MouseScale { get; set; }

        [Option('f', "fullscreen", Default = false, Required = false, HelpText = "When set, the game will start in full screen mode")]
        public bool FullScreen { get; set; }
    }
}
