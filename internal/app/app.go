package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/mattn/go-tty"
)

// PlayerPosition define player keyboard position
type PlayerPosition struct {
	X int
	Y int
}

// Game define attributes for connect4 game
type Game struct {
	P1           PlayerPosition
	P2           PlayerPosition
	PlayerToggle int
	Matrix       [6][7]string
}

// Start listen keyboard events and take action
func (g *Game) Start() {
	clearCMD()
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	for {
		fmt.Printf("Player %d is playing...\n", g.PlayerToggle)
		fmt.Println("(a: left, d: right, w: up, s: down)")
		g.PrintBoard()
		g.HandleMoves(tty)
		clearCMD()
	}
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
		if g.PlayerToggle == 1 {
			g.Matrix[p1.Y][p1.X] = "x"
			g.PlayerToggle = 2
		} else {
			g.Matrix[p1.Y][p1.X] = "o"
			g.PlayerToggle = 1
		}
	default:
		fmt.Println("Key not valid, use (a,s,d,w)")
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
}

// Run Initialize connect4 game
func Run() {
	game := &Game{
		P1:           PlayerPosition{0, 0},
		P2:           PlayerPosition{0, 0},
		PlayerToggle: 1,
		Matrix:       [6][7]string{},
	}
	game.Start()
}

func clearCMD() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
