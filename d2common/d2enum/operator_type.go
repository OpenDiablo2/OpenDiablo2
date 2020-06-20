package d2enum

// OperatorType is used for calculating dynamic property values
type OperatorType int // for dynamic properties

const (
	// OpDefault just adds the stat to the unit directly
	OpDefault = OperatorType(iota)

	// Op1 adds opstat.base * statvalue / 100 to the opstat.
	Op1

	// Op2 adds (statvalue * basevalue) / (2 ^ param) to the opstat
	// this does not work properly with any stat other then level because of the
	// way this is updated, it is only refreshed when you re-equip the item,
	// your character is saved or you level up, similar to passive skills, just
	// because it looks like it works in the item description
	// does not mean it does, the game just recalculates the information in the
	// description every frame, while the values remain unchanged serverside.
	Op2

	// Op3 is a percentage based version of op #2
	// look at op #2 for information about the formula behind it, just
	// remember the stat is increased by a percentage rather then by adding
	// an integer.
	Op3

	// Op4 works the same way op #2 works, however the stat bonus is
	// added to the item and not to the player (so that +defense per level
	// properly adds the defense to the armor and not to the character
	// directly!)
	Op4

	// Op5 works like op #4 but is percentage based, it is used for percentage
	// based increase of stats that are found on the item itself, and not stats
	// that are found on the character.
	Op5

	// Op6 works like for op #7, however this adds a plain bonus to the stat, and just
	// like #7 it also doesn't work so I won't bother to explain the arithmetic
	// behind it either.
	Op6

	// Op7 is used to increase a stat based on the current daytime of the game
	// world by a percentage, there is no need to explain the arithmetics
	// behind it because frankly enough it just doesn't work serverside, it
	// only updates clientside so this op is essentially useless.
	Op7

	// Op8 hardcoded to work only with maxmana, this will apply the proper amount
	// of mana to your character based on CharStats.txt for the amount of energy
	// the stat added (doesn't work for non characters)
	Op8

	// Op9 hardcoded to work only with maxhp and maxstamina, this will apply the
	// proper amount of maxhp and maxstamina to your character based on
	// CharStats.txt for the amount of vitality the stat added (doesn't work
	// for non characters)
	Op9

	// Op10 doesn't do anything, this has no switch case in the op function.
	Op10

	// Op11 adds opstat.base * statvalue / 100 similar to 1 and 13, the code just
	// does a few more checks
	Op11

	// Op12 doesn't do anything, this has no switch case in the op function.
	Op12

	// Op13 adds opstat.base * statvalue / 100 to the value of opstat, this is
	// useable only on items it will not apply the bonus to other unit types
	// (this is why it is used for +% durability, +% level requirement,
	// +% damage, +% defense ).
	Op13
)
