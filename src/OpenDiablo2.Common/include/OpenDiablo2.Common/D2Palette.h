#ifndef OPENDIABLO2_COMMON_D2PALETTE_H
#define OPENDIABLO2_COMMON_D2PALETTE_H

#include <string>

namespace OpenDiablo2::Common {
class D2Palette {
public:
	typedef std::string Entry;
	static const Entry Act1;
	static const Entry Act2;
	static const Entry Act3;
	static const Entry Act4;
	static const Entry Act5;
	static const Entry EndGame;
	static const Entry EndGame2;
	static const Entry Fechar;
	static const Entry Loading;
	static const Entry Menu0;
	static const Entry Menu1;
	static const Entry Menu2;
	static const Entry Menu3;
	static const Entry Menu4;
	static const Entry Sky;
	static const Entry Static;
	static const Entry Trademark;
	static const Entry Units;
private:
	D2Palette() {}
};


const D2Palette::Entry D2Palette::Act1 = "ACT1";
const D2Palette::Entry D2Palette::Act2 = "ACT2";
const D2Palette::Entry D2Palette::Act3 = "ACT3";
const D2Palette::Entry D2Palette::Act4 = "ACT4";
const D2Palette::Entry D2Palette::Act5 = "ACT5";
const D2Palette::Entry D2Palette::EndGame = "EndGame";
const D2Palette::Entry D2Palette::EndGame2 = "EndGame2";
const D2Palette::Entry D2Palette::Fechar = "fechar";
const D2Palette::Entry D2Palette::Loading = "loading";
const D2Palette::Entry D2Palette::Menu0 = "Menu0";
const D2Palette::Entry D2Palette::Menu1 = "menu1";
const D2Palette::Entry D2Palette::Menu2 = "menu2";
const D2Palette::Entry D2Palette::Menu3 = "menu3";
const D2Palette::Entry D2Palette::Menu4 = "menu4";
const D2Palette::Entry D2Palette::Sky = "Sky";
const D2Palette::Entry D2Palette::Static = "STATIC";
const D2Palette::Entry D2Palette::Trademark = "Trademark";
const D2Palette::Entry D2Palette::Units = "Units";


}

#endif // OPENDIABLO2_COMMON_D2PALETTE_H
