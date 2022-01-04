package tictactoe

import (
	"github.com/aryanugroho/tictacgo/domain"
)

type Tictactoe struct {
	tictactoe domain.Game
}

func NewTictactoeUseCase() domain.TictactoeUseCase {
	return &Tictactoe{
		tictactoe: domain.NewTictactoe(),
	}
}

func (u *Tictactoe) Play(position int) (bool, string, error) {

	err := u.tictactoe.Play(position)
	if err != nil {
		return false, "", err
	}

	u.tictactoe.SwitchPlayer()

	anyWinner, winner := u.tictactoe.CheckForWinner()

	return anyWinner, winner, nil
}

func (u *Tictactoe) GetBoard() []string {
	return u.tictactoe.GetBoard()
}

func (u *Tictactoe) IsOver() bool {
	return u.tictactoe.IsOver()
}
