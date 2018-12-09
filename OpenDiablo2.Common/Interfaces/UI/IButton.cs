using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    /// <summary>
    /// A callback that is executed the button is activated (clicked).
    /// </summary>
    public delegate void OnActivateDelegate();

    /// <summary>
    /// A callback that is executed if the button is toggled.
    /// </summary>
    /// <param name="isToggled">When enabled, this is true, otherwise false</param>
    public delegate void OnToggleDelegate(bool isToggled);

    /// <summary>
    /// Represents a visual button on the screen.
    /// </summary>
    public interface IButton : IDisposable
    {
        /// <summary>
        /// Assigning a function to this property will cause that function to be called
        /// when a button is pressed.
        /// </summary>
        OnActivateDelegate OnActivate { get; set; }

        /// <summary>
        /// If false, the button is visually darkened, and will ignore all user input.
        /// </summary>
        bool Enabled { get; set; }

        /// <summary>
        /// If true, the button is pushed down, false otherwise. Only valid for toggle buttons.
        /// </summary>
        bool Toggled { get; }

        /// <summary>
        /// The position of the button on the screen.
        /// </summary>
        Point Location { get; set; }

        /// <summary>
        /// Area from upper left corner that reacts to clicking
        /// </summary>
        Size ClickableRect { get; set; }

        /// <summary>
        /// Indicates if button sprite should react to Toggle and Activate on hover
        /// </summary>
        bool AllowFrameChange { get; set; }

        /// <summary>
        /// Assigning a function to this property will cause that function to be called
        /// when the button is toggled on or off.
        /// </summary>
        OnToggleDelegate OnToggle { get; set; }

        /// <summary>
        /// Toggle the button. Only valid for toggle buttons.
        /// </summary>
        bool Toggle();

        string Text { get; set; }

        /// <summary>
        /// Allows the button to update its internal state.
        /// Call this in the Update method of your scene.
        /// </summary>
        void Update();

        /// <summary>
        /// Renders the button to the screen.
        /// Call this in the render method of your scene.
        /// </summary>
        void Render();
    }
}
