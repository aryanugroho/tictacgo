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

const (
	playerO    = "O"
	playerX    = "X"
	playerDraw = ""
)

func NewTictactoe() Game {
	return &Tictactoe{
		player: playerO,
		isOver: false,
	}
}

func (u *Tictactoe) CheckForWinner() (bool, string) {
	anyWinner := false
	i := 0

	for i < 9 {
		anyWinner = u.board[i] == u.board[i+1] && u.board[i+1] == u.board[i+2] && u.board[i] != ""
		if !anyWinner {
			i += 3
		} else {
			return true, u.board[i]
		}
	}

	i = 0
	for i < 3 {
		anyWinner = u.board[i] == u.board[i+3] && u.board[i+3] == u.board[i+6] && u.board[i] != ""
		if !anyWinner {
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
