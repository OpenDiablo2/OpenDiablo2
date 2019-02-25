#ifndef OPENDIABLO2_D2INPUT_H
#define OPENDIABLO2_D2INPUT_H

#include <memory>

namespace OpenDiablo2::System {

class D2Input {
public:
	typedef std::unique_ptr<D2Input> Ptr;
	D2Input();
	void ProcessEvents();
	bool QuitIsRequested();
private:
	bool quitIsRequested = false;
};

}

#endif //OPENDIABLO2_D2INPUT_H
