#ifndef OPENDIABLO2_GAME_COMMON_PALETTE_H
#define OPENDIABLO2_GAME_COMMON_PALETTE_H

#include <string>

namespace OpenDiablo2::Game::Common {

class Palette {
public:
	const std::string Act1 = "ACT1";
	const std::string Act2 = "ACT2";
	const std::string Act3 = "ACT3";
	const std::string Act4 = "ACT4";
	const std::string Act5 = "ACT5";
	const std::string EndGame = "EndGame";
	const std::string EndGame2 = "EndGame2";
	const std::string Fechar = "fechar";
	const std::string Loading = "loading";
	const std::string Menu0 = "Menu0";
	const std::string Menu1 = "menu1";
	const std::string Menu2 = "menu2";
	const std::string Menu3 = "menu3";
	const std::string Menu4 = "menu4";
	const std::string Sky = "Sky";
	const std::string Static = "STATIC";
	const std::string Trademark = "Trademark";
	const std::string Units = "Units";
private:
	Palette() {}
};

}

#endif // OPENDIABLO2_GAME_COMMON_PALETTE_H
