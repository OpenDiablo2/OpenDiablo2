#ifndef OPENDIABLO2_GAME_COMMON_PALETTE_H
#define OPENDIABLO2_GAME_COMMON_PALETTE_H

#include <string>

namespace OpenDiablo2::Game::Common {
class D2Palette {
public:
	typedef std::string Entry;
	const Entry Act1 = "ACT1";
	const Entry Act2 = "ACT2";
	const Entry Act3 = "ACT3";
	const Entry Act4 = "ACT4";
	const Entry Act5 = "ACT5";
	const Entry EndGame = "EndGame";
	const Entry EndGame2 = "EndGame2";
	const Entry Fechar = "fechar";
	const Entry Loading = "loading";
	const Entry Menu0 = "Menu0";
	const Entry Menu1 = "menu1";
	const Entry Menu2 = "menu2";
	const Entry Menu3 = "menu3";
	const Entry Menu4 = "menu4";
	const Entry Sky = "Sky";
	const Entry Static = "STATIC";
	const Entry Trademark = "Trademark";
	const Entry Units = "Units";
private:
	Palette() {}
};

}

#endif // OPENDIABLO2_GAME_COMMON_PALETTE_H
