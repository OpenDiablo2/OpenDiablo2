package d2render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngleToDirection_16Directions(t *testing.T) {

	numberOfDirections := 16

	angle := 45.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 22.5
	}

	angle = 50.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 22.5
	}

	angle = 40.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 22.5
	}

}

func TestAngleToDirection_8Directions(t *testing.T) {

	numberOfDirections := 8

	angle := 45.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 45
	}

	angle = 50.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 45
	}

	angle = 40.0
	for i := 0; i < numberOfDirections; i++ {
		assert.Equal(t, i, angleToDirection(angle, numberOfDirections))
		angle += 45
	}

}

func TestAngleToDirection_1Direction(t *testing.T) {
	angle := 0.0
	for i := 0; i < 120; i++ {
		assert.Equal(t, 0, angleToDirection(angle, 1))
		angle += 3
	}
}

func TestAngleToDirection_0Directions(t *testing.T) {
	angle := 0.0
	for i := 0; i < 120; i++ {
		assert.Equal(t, 0, angleToDirection(angle, 0))
		angle += 3
	}
}
