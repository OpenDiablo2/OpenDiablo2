package d2calculation

// CalcString is a type of string often used in datafiles to specify
// a value that is calculated dynamically based on the stats of the relevant
// source, for instance a missile might have a movement speed of lvl*2
type CalcString string

// Issue #689
// info about calcstrings can be found here: https://d2mods.info/forum/kb/viewarticle?a=371
