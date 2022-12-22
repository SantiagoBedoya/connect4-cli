package app

import "github.com/SantiagoBedoya/connect4/internal/models"

// Run Initialize connect4 game
func Run() {
	game := &models.Game{
		P1:           models.PlayerPosition{X: 0, Y: 0},
		P2:           models.PlayerPosition{X: 0, Y: 0},
		PlayerToggle: 1,
		Matrix:       [6][7]string{},
	}
	game.Start()
}
