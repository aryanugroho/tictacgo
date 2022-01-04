package domain

import (
	"errors"
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
}

type TictactoeUseCase interface {
	Play(position int) (bool, string, error)
}

const (
	playerO    = "O"
	playerX    = "X"
	playerDraw = ""
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
