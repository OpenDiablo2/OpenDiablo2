namespace OpenDiablo2.Common.Models
{
    public enum eMouseMode
    {
        Software,
        Hardware
    }

    public sealed class GlobalConfiguration
    {
        public string BaseDataPath { get; set; }
        public eMouseMode MouseMode { get; set; }
        public int HardwareMouseScale { get; set; }

    }
}
