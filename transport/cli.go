package transport

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aryanugroho/tictacgo/domain"
	"github.com/aryanugroho/tictacgo/usecase/tictactoe"
)

type CliTransport struct {
	useCase domain.TictactoeUseCase
}

func NewCliTransport() Transport {
	return &CliTransport{
		useCase: tictactoe.NewTictactoeUseCase(),
	}
}

func (u *CliTransport) drawBoard() {
	for i, v := range u.useCase.GetBoard() {
		if v == "" {
			fmt.Print(" ")
		} else {
			fmt.Print(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print("|")

		}

	}
}

func (u *CliTransport) next() bool {
	var movePosition int
	fmt.Println("Enter position: ")
	fmt.Scan(&movePosition)

	anyWinner, winner, err := u.useCase.Play(movePosition)
	if err != nil {
		if err == domain.ErrInvalidPosition {
			fmt.Println("Invalid position, try another move")
		}
	}

	if anyWinner {
		fmt.Println(winner, "WIN!")
	}
	return anyWinner
}

func (u *CliTransport) clearBoard() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (u *CliTransport) Run() {
	for !u.useCase.IsOver() {
		u.drawBoard()
		anyWinner := u.next()
		if anyWinner {
			break
		}
	}
	u.drawBoard()
}
