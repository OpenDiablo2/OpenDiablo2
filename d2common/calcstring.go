package d2common

// CalcString is a type of string often used in datafiles to specify
// a value that is calculated dynamically based on the stats of the relevant
// source, for instance a missile might have a movement speed of lvl*2
type CalcString string

// todo: the logic for parsing these should exist here
