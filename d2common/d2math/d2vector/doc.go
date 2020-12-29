// Package d2vector provides an implementation of a 2D Euclidean vector using float64 to store the two values.
/*
Vector uses d2math.Epsilon for approximate equality and comparison. Note: SetLength, Reflect, ReflectSurface and Rotate
do not (per their unit tests) return exact values but ones within Epsilon range of the expected value.*/
package d2vector
