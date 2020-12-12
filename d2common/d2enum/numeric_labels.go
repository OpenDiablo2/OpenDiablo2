package d2enum

// there are labels for "numeric labels (see AssetManager.TranslateLabel)
const (
	CancelLabel = iota
	CopyrightLabel
	AllRightsReservedLabel
	SinglePlayerLabel
	_
	OtherMultiplayerLabel
	ExitGameLabel
	CreditsLabel
	CinematicsLabel

	ViewAllCinematicsLabel
	EpilogueLabel
	SelectCinematicLabel

	_
	TCPIPGameLabel
	TCPIPOptionsLabel
	TCPIPHostGameLabel
	TCPIPJoinGameLabel
	TCPIPEnterHostIPLabel
	TCPIPYourIPLabel
	TipHostLabel
	TipJoinLabel
	IPNotFoundLabel

	CharNameLabel
	HardCoreLabel
	SelectHeroClassLabel
	AmazonDescr
	NecromancerDescr
	BarbarianDescr
	SorceressDescr
	PaladinDescr

	_

	HellLabel
	NightmareLabel
	NormalLabel
	SelectDifficultyLabel

	_

	DelCharConfLabel
	OpenLabel

	_

	YesLabel
	NoLabel

	_

	ExitLabel
	OKLabel
)

// BaseLabelNumbers returns base label value (#n in english string table table)
func BaseLabelNumbers(idx int) int {
	baseLabelNumbers := []int{
		// main menu labels
		1612, // CANCEL
		1613, // (c) 2000 Blizzard Entertainment
		1614, // All Rights Reserved.
		1620, // SINGLE PLAYER
		1621, // BATTLE.NET
		1623, // OTHER MULTIPLAYER
		1625, // EXIT DIABLO II
		1627, // CREDITS
		1639, // CINEMATICS

		// cinematics menu labels
		1640, // View All Earned Cinematics
		1659, // Epilogue
		1660, // SELECT CINEMATICS

		// multiplayer labels
		1663, // OPEN BATTLE.NET
		1666, // TCP/IP GAME
		1667, // TCP/IP Options
		1675, // HOST GAME
		1676, // JOIN GAME
		1678, // Enter Host IP Address to Join Game
		1680, // Your IP Address is:
		1689, // Tip: host game
		1690, // Tip: join game
		1691, // Cannot detect a valid TCP/IP address.
		1694, // Character Name
		1696, // Hardcore
		1697, // Select Hero Class

		1698, // amazon description
		1704, // nec description
		1709, // barb description
		1710, // sorc description
		1711, // pal description
		/*in addition, as many elements as the value
		  of the highest modifier must be listed*/
		1712,

		/* here, should be labels used to battle.net multiplayer, but they are not used yet,
		   therefore I don't list them here.*/

		// difficulty levels:
		1800, // Hell
		1864, // Nightmare
		1865, // Normal
		1867, // Select Difficulty

		1869, // not used, for locales with +1 mod
		1878, // delete char confirm
		1881, // Open
		1889, // char name is currently taken (not used)
		1896, // YES
		1925, // NO

		1926, // not used, for locales with +1 mod

		970, // EXIT
		971, // OK
		1612,
	}

	return baseLabelNumbers[idx]
}
