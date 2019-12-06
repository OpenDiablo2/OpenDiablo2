package d2render

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
