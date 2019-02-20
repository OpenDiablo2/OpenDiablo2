#ifndef OPENDIABLO2_WINDOW_H
#define OPENDIABLO2_WINDOW_H

#include <memory>
#include <GLFW/glfw3.h>


namespace OpenDiablo2 { namespace OS {

class D2Window {
public:
	D2Window();
	void Run();
private:
	GLFWwindow* glfwWindow;

	void
	InitializeWindow();

	void
	FinalizeWindow();
};

typedef std::unique_ptr<D2Window> D2WindowPtr;

}}

#endif //OPENDIABLO2_WINDOW_H
