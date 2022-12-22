package models

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/mattn/go-tty"
)

// Game define attributes for connect4 game
type Game struct {
	P1           PlayerPosition
	P2           PlayerPosition
	PlayerToggle int
	Matrix       [6][7]string
	Message      string
	Moves        int
	ExistWinner  bool
}

// Start listen keyboard events and take action
func (g *Game) Start() {
	g.ClearCMD()
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	for {
		fmt.Printf("Player %d is playing...\n", g.PlayerToggle)
		fmt.Println("(a: left, d: right, w: up, s: down, spacebar: mark slot)")
		fmt.Printf("Moves: %d\n", g.Moves)
		g.PrintBoard()
		if g.ExistWinner {
			break
		}
		g.HandleMoves(tty)
		if g.Moves >= 7 {
			isP1Winner := g.VerifyBoard(1)
			isP2Winner := g.VerifyBoard(2)
			if isP1Winner {
				g.Message = "Player 1 is winner"
				g.ExistWinner = true
			}
			if isP2Winner {
				g.Message = "Player 2 is winner"
				g.ExistWinner = true
			}
		}
		g.ClearCMD()
	}
}

// VerifyBoard verify if player has 4 slot in line
func (g *Game) VerifyBoard(player int) bool {
	var char string
	if player == 1 {
		char = "x"
	} else {
		char = "o"
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if g.Matrix[i][j] == char &&
				g.Matrix[i][j+1] == char &&
				g.Matrix[i][j+2] == char &&
				g.Matrix[i][j+3] == char {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 6; j++ {
			if g.Matrix[i][j] == char &&
				g.Matrix[i+1][j] == char &&
				g.Matrix[i+2][j] == char &&
				g.Matrix[i+3][j] == char {
				return true
			}
		}
	}
	return false
}

// HandleMoves handle player moves
func (g *Game) HandleMoves(tty *tty.TTY) {
	r, err := tty.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	var p1 *PlayerPosition
	if g.PlayerToggle == 1 {
		p1 = &g.P1
	} else {
		p1 = &g.P2
	}
	switch string(r) {
	case "a":
		if p1.X > 0 {
			p1.X--
		}
	case "s":
		if p1.Y < 5 {
			p1.Y++
		}
	case "w":
		if p1.Y > 0 {
			p1.Y--
		}
	case "d":
		if p1.X < 6 {
			p1.X++
		}
	case " ":
		g.Moves++
		if g.PlayerToggle == 1 {
			if len(g.Matrix[p1.Y][p1.X]) > 0 {
				g.Message = "Slot already marked"
				return
			}
			g.Message = ""
			g.Matrix[p1.Y][p1.X] = "x"
			g.PlayerToggle = 2
		} else {
			if len(g.Matrix[p1.Y][p1.X]) > 0 {
				g.Message = "Slot already marked"
				return
			}
			g.Message = ""
			g.Matrix[p1.Y][p1.X] = "o"
			g.PlayerToggle = 1
		}
	default:
		g.Message = "Key not valid, use (a,s,d,w, spacebar)"
	}
}

// PrintBoard print connect4 board in console
func (g *Game) PrintBoard() {
	var colored *color.Color
	var playerPos PlayerPosition
	if g.PlayerToggle == 1 {
		playerPos = g.P1
		colored = color.New(color.FgRed).Add(color.Underline)
	} else {
		playerPos = g.P2
		colored = color.New(color.FgYellow).Add(color.Underline)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if i == playerPos.Y && j == playerPos.X {
				fmt.Printf("%s ", colored.Sprint("|*|"))
			} else {
				if len(g.Matrix[i][j]) > 0 {
					if g.Matrix[i][j] == "x" {
						c := color.New(color.FgRed)
						fmt.Printf("%s ", c.Sprint("|x|"))
					}
					if g.Matrix[i][j] == "o" {
						c := color.New(color.FgYellow)
						fmt.Printf("%s ", c.Sprint("|o|"))
					}
				} else {
					fmt.Print("| | ")
				}
			}
		}
		fmt.Println()
		fmt.Println("---------------------------")
	}
	fmt.Printf("Message: %s\n", g.Message)
}

// ClearCMD clean cmd output
func (g *Game) ClearCMD() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
