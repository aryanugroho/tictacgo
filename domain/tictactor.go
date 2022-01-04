package domain

type Tictactoe struct {
	board        [9]string
	player       string
	stepCounter  int
	isVsComputer bool
	isOver       bool
}

type Game interface {
	CheckForWinner() (bool, string)
	ClearScreen()
	SwitchPlayer()
	play() error
	Next() int
	DrawBoard()
}
