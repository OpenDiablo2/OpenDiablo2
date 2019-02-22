#include <OpenDiablo2.OS/D2Window.h>
#include <spdlog/spdlog.h>
#include <GLFW/glfw3.h>

namespace OpenDiablo2 {
	namespace OS {

		D2Window::D2Window() = default;

		void
		D2Window::Initialize() {
			spdlog::debug("Initializing D2Window");

			spdlog::debug("Initializing GLFW");
			if (!glfwInit()) {
				spdlog::error("Initializing D2Window");
				throw std::runtime_error(
					 "GLFW could not initialize the host window.");
			}

			spdlog::debug("Creating GLFW window");
			glfwWindowHint(GLFW_RESIZABLE, 0);
			glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
			glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 0);
			glfwWindow = glfwCreateWindow(800, 600, "OpenDiablo 2", nullptr, nullptr);
		}

		void
		D2Window::Finalize() {
			spdlog::debug("Destroying GLFW window");
			glfwDestroyWindow(glfwWindow);

			spdlog::debug("Terminating GLFW");
			glfwTerminate();
		}

		bool
		D2Window::WindowStillOpen() {
			return !glfwWindowShouldClose(glfwWindow);
		}

		void
		D2Window::PollEvents() {
			glfwPollEvents();
		}

		void
		D2Window::FlipBuffer() {
			glfwSwapBuffers(glfwWindow);
		}
	}
}
