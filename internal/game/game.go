// Package game represents a programmatic version of a tic-tac-toe game. The
// actual game logic is encapsulated within the Game type. Internally, Game
// manages a Board and implements the logic to alternate and apply turns
// between two players, as well as determine when the game is over and who is
// the winner (assuming there is one).
//
// Did you know that apparently these sorts of games can be traced back to
// ancient Egypt?
package game

import (
	"fmt"
)

// Game represents a game of tic-tac-toe and provides logic to manage a Board.
type Game struct {
	board *Board
}

// NewGame is the constructor for Game.
func NewGame() *Game {
	return &Game{board: NewBoard()}
}

// MainLoop starts and runs the Tic-Tac-Toe game in the foreground until
// completion or cancellation.
func (g *Game) MainLoop() error {
	var players = []Mark{X, O}
	for {
		for _, player := range players {
			fmt.Printf("%s's turn\n", player.ToString())
			if err := g.applyTurn(player); err != nil {
				return err
			}
			fmt.Println(g.board.ToString())
			if g.board.IsThreeInARow(player) {
				fmt.Printf("Game over - %s wins!\n", player.ToString())
				return nil
			}
			if g.board.IsFull() {
				fmt.Print("Game over - cat's game (draw)!")
				return nil
			}
		}
	}
}

func (g *Game) applyTurn(mark Mark) error {
	for {
		fmt.Print("Enter coordinates: ")
		row, col, err := scanInput()
		if err != nil {
			// The app will panic if non-ints are passed to stdin ¯\_(ツ)_/¯
			return err
		}
		if err := g.board.ApplyMark(row, col, mark); err != nil {
			// Errors from this function are non-fatal
			fmt.Printf("Error: %v; try again\n", err)
		} else {
			return nil
		}
	}
}

// TODO: better input validation? -ccampo 2022-02-16
func scanInput() (row, col int, err error) {
	if _, err = fmt.Scan(&row); err != nil {
		return 0, 0, err
	}
	if _, err = fmt.Scan(&col); err != nil {
		return 0, 0, err
	}
	return row, col, nil
}
