#ifndef OPENDIABLO2_WINDOW_H
#define OPENDIABLO2_WINDOW_H

#include <memory>


class GLFWwindow;

namespace OpenDiablo2 { namespace OS {

class D2Window {
public:
	D2Window();
	void Initialize();
	void Finalize();
	bool WindowStillOpen();
	void FlipBuffer();
	void PollEvents();
private:
	GLFWwindow* glfwWindow;
};

typedef std::unique_ptr<D2Window> D2WindowPtr;

}}

#endif //OPENDIABLO2_WINDOW_H
