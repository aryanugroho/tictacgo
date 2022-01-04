package tictactoe

import (
	"fmt"

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

	// human play
	err := u.tictactoe.Play(position)
	if err != nil {
		return false, "", err
	}
	u.tictactoe.SwitchPlayer()

	if !u.tictactoe.IsOver() {
		// computer play
		compPosition := u.tictactoe.GetComputerPosition()
		if compPosition > 0 {
			fmt.Println("Computer position: ", compPosition)
			err = u.tictactoe.Play(compPosition)
			if err != nil {
				return false, "", err
			}
		}

		u.tictactoe.SwitchPlayer()
	}

	anyWinner, winner := u.tictactoe.CheckForWinner()

	return anyWinner, winner, nil
}

func (u *Tictactoe) GetBoard() []string {
	return u.tictactoe.GetBoard()
}

func (u *Tictactoe) IsOver() bool {
	return u.tictactoe.IsOver()
}
