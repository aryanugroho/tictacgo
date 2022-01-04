package domain

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/aryanugroho/tictacgo/common/arrayutil"
)

type Tictactoe struct {
	board       [9]string
	player      string
	stepCounter int
	isOver      bool
}

type Game interface {
	CheckForWinner() (bool, string)
	SwitchPlayer() string
	Play(position int) error
	GetBoard() []string
	IsOver() bool
	GetComputerPosition() int
}

type TictactoeUseCase interface {
	Play(position int) (bool, string, error)
	GetBoard() []string
	IsOver() bool
}

const (
	playerO  = "O" // human
	playerX  = "X" // computer
	noplayer = ""  // empty
)

var (
	ErrInvalidPosition = errors.New("invalid position")
)

func NewTictactoe() Game {
	return &Tictactoe{
		player: playerO,
		isOver: false,
	}
}

func (u *Tictactoe) CheckForWinner() (bool, string) {
	i := 0
	for i < 9 {
		u.isOver = u.board[i] == u.board[i+1] && u.board[i+1] == u.board[i+2] && u.board[i] != ""
		if !u.isOver {
			i += 3
		} else {
			return true, u.board[i]
		}
	}

	i = 0
	for i < 3 {
		u.isOver = u.board[i] == u.board[i+3] && u.board[i+3] == u.board[i+6] && u.board[i] != ""
		if !u.isOver {
			i += 1
		} else {
			return true, u.board[i]
		}
	}

	if u.board[0] == u.board[4] && u.board[4] == u.board[8] && u.board[0] != "" {
		return true, u.board[0]
	}

	if u.board[2] == u.board[4] && u.board[4] == u.board[6] && u.board[2] != "" {
		return true, u.board[2]
	}

	if u.stepCounter == len(u.board) {
		return false, ""
	}

	return false, ""
}

func (u *Tictactoe) SwitchPlayer() string {
	if u.player == playerO {
		u.player = playerX
		return u.player
	}
	u.player = playerO
	return playerO
}

func (u *Tictactoe) Play(position int) error {
	if u.board[position-1] == "" {
		u.board[position-1] = u.player
		u.stepCounter += 1
		return nil
	}
	return ErrInvalidPosition
}

func (u *Tictactoe) GetBoard() []string {
	return u.board[:]
}

func (u *Tictactoe) IsOver() bool {
	return u.isOver
}

func (u *Tictactoe) GetComputerPosition() int {
	if u.stepCounter == len(u.board) {
		return -1
	}

	var choice int
	var emptyCell []int
	var humanCell []int
	var compCell []int

	// collect empty cell
	for i, v := range u.board {
		if v == noplayer {
			idxToPosition := i + 1
			emptyCell = append(emptyCell, idxToPosition)
		}
	}

	// collect human cell
	for i, v := range u.board {
		if v == playerO {
			idxToPosition := i + 1
			humanCell = append(humanCell, idxToPosition)
		}
	}
	anyPattern, pattern := isAnyPattern(humanCell)
	if anyPattern {
		fmt.Println("human", pattern)
		any := arrayutil.IsAny(pattern, emptyCell)
		if len(any) > 0 {
			return any[0]
		}

	}

	// collect comp cell
	for i, v := range u.board {
		if v == playerX {
			idxToPosition := i + 1
			compCell = append(compCell, idxToPosition)
		}
	}
	anyPattern, pattern = isAnyPattern(compCell)
	if anyPattern {
		fmt.Println("comp", pattern)
		any := arrayutil.IsAny(pattern, emptyCell)
		if len(any) > 0 {
			return any[0]
		}

	}

	fmt.Println("rand")
	cellChoice := rand.Intn(len(emptyCell))
	choice = emptyCell[cellChoice]
	return choice
}

func isAnyPattern(cell []int) (bool, []int) {

	// possible wins pattern
	var patterns [][]int
	patterns = append(patterns, []int{1, 2, 3})
	patterns = append(patterns, []int{4, 5, 6})
	patterns = append(patterns, []int{7, 8, 9})

	patterns = append(patterns, []int{1, 4, 6})
	patterns = append(patterns, []int{2, 5, 8})
	patterns = append(patterns, []int{3, 6, 9})

	patterns = append(patterns, []int{1, 5, 9})
	patterns = append(patterns, []int{3, 5, 7})

	for _, p := range patterns {
		contains := arrayutil.IsContains(p, cell, 2)
		if contains {
			return true, p
		}
	}
	return false, []int{}
}
