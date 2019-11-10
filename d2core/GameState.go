package d2core

import (
	"log"
	"time"
)

type GameState struct {
	Seed int64
}

func CreateGameState() *GameState {
	result := &GameState{
		Seed: time.Now().UnixNano(),
	}
	return result
}

func LoadGameState(path string) *GameState {
	log.Fatal("LoadGameState not implemented.")
	return nil
}
