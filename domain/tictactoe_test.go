package domain

import (
	"testing"
)

func TestTictactoe_CheckForWinner(t *testing.T) {
	tests := []struct {
		tictactoe Game
		anyWinner bool
		winner    string
		wantErr   error
	}{
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerO, playerO, playerO, playerX, "", playerX, "", "", ""},
				stepCounter: 5,
				player:      playerX,
			},
			anyWinner: true,
			winner:    playerO,
			wantErr:   nil,
		},
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerX, playerX, playerX, playerO, "", playerO, "", "", ""},
				stepCounter: 5,
				player:      playerO,
			},
			anyWinner: true,
			winner:    playerX,
			wantErr:   nil,
		},
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerO, playerX, playerO, playerO, playerX, playerO, playerX, playerO, playerX},
				stepCounter: 9,
				player:      playerX,
			},
			anyWinner: false,
			winner:    playerDraw,
			wantErr:   nil,
		},
	}

	for i, test := range tests {
		anyWinner, winner := test.tictactoe.CheckForWinner()
		if winner != test.winner {
			t.Error("test", i+1, "failed, expect winner", test.winner, "got", winner)
		}
		if anyWinner != test.anyWinner {
			t.Error("test", i+1, "failed, expect anyWinner", test.anyWinner, "got anyWinner", anyWinner)
		}

	}
}
