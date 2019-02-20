#include <D2Window.h>
#define GLFW_INCLUDE_NONE

namespace OpenDiablo2 {
	namespace OS {

		D2Window::D2Window() {

		}

		void
		D2Window::Run() {
			InitializeWindow();
			while (!glfwWindowShouldClose(glfwWindow)) {
				glfwPollEvents();
			}
			FinalizeWindow();
		}

		void
		D2Window::InitializeWindow() {
			if (!glfwInit()) {
				throw std::runtime_error(
					 "GLFW could not initialize the host window.");
			}

			glfwWindowHint(GLFW_RESIZABLE, 0);
			glfwWindow = glfwCreateWindow(800, 600, "OpenDiablo 2", NULL, NULL);
		}

		void
		D2Window::FinalizeWindow() {
			glfwDestroyWindow(glfwWindow);
			glfwTerminate();
		}
	}
}
