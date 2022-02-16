package game

import (
	"errors"
	"fmt"
	"strings"
)

// Board represents a 3x3 tic-tac-toe game board.
type Board struct {
	grid [][]*Mark
}

// NewBoard is the constructor for Board.
func NewBoard() *Board {
	return &Board{
		grid: [][]*Mark{
			{nil, nil, nil},
			{nil, nil, nil},
			{nil, nil, nil},
		},
	}
}

// ApplyMark applies a Mark to the given board grid space indicated by the row
// and column number. This is the programmatic equivalent to drawing an X or O
// on a physical tic-tac-toe grid :).
func (b *Board) ApplyMark(row, col int, mark Mark) error {
	if row < 0 || row > 2 {
		return fmt.Errorf("row %d is out of range [0, 2]", row)
	}
	if col < 0 || col > 2 {
		return fmt.Errorf("col %d is out of range [0, 2]", col)
	}
	if b.grid[row][col] != nil {
		return errors.New("space is occupied")
	}
	b.grid[row][col] = &mark
	return nil
}

// IsThreeInARow returns true if the grid has a given Mark (X or O) placed
// three-in-a-row along any row, column, or diagonal.
func (b *Board) IsThreeInARow(mark Mark) bool {
	// Diagonals
	if b.spaceEquals(0, 0, mark) && b.spaceEquals(1, 1, mark) && b.spaceEquals(2, 2, mark) {
		return true
	}
	if b.spaceEquals(0, 2, mark) && b.spaceEquals(1, 1, mark) && b.spaceEquals(2, 0, mark) {
		return true
	}
	for i := 0; i < 3; i++ {
		// Horizontal
		if b.spaceEquals(i, 0, mark) && b.spaceEquals(i, 1, mark) && b.spaceEquals(i, 2, mark) {
			return true
		}
		// Vertical
		if b.spaceEquals(0, i, mark) && b.spaceEquals(1, i, mark) && b.spaceEquals(2, i, mark) {
			return true
		}
	}
	return false
}

// IsFull returns true if all the board spaces are currently marked.
func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.grid[i][j] == nil {
				return false
			}
		}
	}
	return true
}

// ToString returns a string representation of the game board.
func (b *Board) ToString() string {
	builder := strings.Builder{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			builder.WriteString(b.grid[i][j].ToString())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b *Board) spaceEquals(row, col int, mark Mark) bool {
	return b.grid[row][col] != nil && *b.grid[row][col] == mark
}
