package domain

type Tictactoe struct {
	board       [9]string
	player      string
	stepCounter int
	isOver      bool
}

type Game interface {
	CheckForWinner() (bool, string)
	ClearScreen()
	SwitchPlayer()
	play() error
	Next() int
	DrawBoard()
}

func NewTictactoe() Game {
	return &Tictactoe{
		player: "O",
		isOver: false,
	}
}

func (u *Tictactoe) CheckForWinner() (bool, string) {
	return false, ""
}

func (u *Tictactoe) ClearScreen() {

}

func (u *Tictactoe) SwitchPlayer() {

}

func (u *Tictactoe) play() error {
	return nil
}

func (u *Tictactoe) Next() int {
	return 0
}

func (u *Tictactoe) DrawBoard() {

}
