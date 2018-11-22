using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Attributes
{
    [AttributeUsage(AttributeTargets.Class, Inherited = false, AllowMultiple = false)]
    public sealed class SceneAttribute : Attribute
    {
        public SceneAttribute(string sceneName) => SceneName = sceneName;
        public string SceneName { get; }
    }
}
