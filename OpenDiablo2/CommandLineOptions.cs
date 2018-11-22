using CommandLine;

namespace OpenDiablo2
{
    public sealed class CommandLineOptions
    {
        [Option('p', "datapath", Required = false, HelpText = "Specifies the root data path")]
        public string DataPath { get; set; }
    }
}
