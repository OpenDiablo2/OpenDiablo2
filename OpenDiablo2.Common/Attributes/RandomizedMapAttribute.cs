using System;

namespace OpenDiablo2.Common.Attributes
{
    [AttributeUsage(AttributeTargets.Class, Inherited = false, AllowMultiple = false)]
    public sealed class RandomizedMapAttribute : Attribute
    {
        readonly string mapName;
        public string MapName => mapName;
        
        public RandomizedMapAttribute(string mapName)
        {
            this.mapName = mapName;
        }

        
        
    }
}
