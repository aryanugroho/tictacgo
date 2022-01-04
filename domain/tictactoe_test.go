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
				board:       [9]string{playerX, playerO, playerO, playerO, playerX, playerX, playerX, playerO, playerO},
				stepCounter: 9,
				player:      playerX,
			},
			anyWinner: true,
			winner:    noplayer,
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

func TestTictactoe_play(t *testing.T) {
	tests := []struct {
		tictactoe    Game
		movePosition int
		wantErr      error
	}{
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerO, playerO, playerO, playerX, "", playerX, "", "", ""},
				stepCounter: 5,
				player:      playerX,
			},
			movePosition: 5,
			wantErr:      nil,
		},
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerX, playerX, playerX, playerO, "", playerO, "", "", ""},
				stepCounter: 5,
				player:      playerO,
			},
			movePosition: 2,
			wantErr:      ErrInvalidPosition,
		},
	}

	for i, test := range tests {
		err := test.tictactoe.Play(test.movePosition)
		if err != test.wantErr {
			t.Error("test", i+1, "failed, expect wantErr", test.wantErr, "got", err)
		}
	}
}

func TestTictactoe_switchPlayer(t *testing.T) {
	tests := []struct {
		tictactoe  Game
		nextPlayer string
	}{
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerO, playerO, playerO, playerX, "", playerX, "", "", ""},
				stepCounter: 5,
				player:      playerX,
			},
			nextPlayer: playerO,
		},
		{
			tictactoe: &Tictactoe{
				board:       [9]string{playerX, playerX, playerX, playerO, "", playerO, "", "", ""},
				stepCounter: 5,
				player:      playerO,
			},
			nextPlayer: playerX,
		},
	}

	for i, test := range tests {
		nextPlayer := test.tictactoe.SwitchPlayer()
		if nextPlayer != test.nextPlayer {
			t.Error("test", i+1, "failed, expect wantErr", test.nextPlayer, "got", nextPlayer)
		}
	}
}

func TestTictactoe_GetComputerPosition(t *testing.T) {
	tests := []struct {
		tictactoe Game
		wantErr   error
	}{
		{
			tictactoe: &Tictactoe{
				board:  [9]string{playerX, "", playerO, "", "", playerO, "", "", ""},
				player: playerX,
			},
			wantErr: nil,
		},
		{
			tictactoe: &Tictactoe{
				board:  [9]string{playerX, playerO, playerO, "", "", playerO, "", "", ""},
				player: playerX,
			},
			wantErr: nil,
		},
		{
			tictactoe: &Tictactoe{
				board:  [9]string{playerX, playerO, playerO, playerO, playerX, playerO, "", "", ""},
				player: playerX,
			},
			wantErr: nil,
		},
	}

	for i, test := range tests {
		compPosition := test.tictactoe.GetComputerPosition()
		err := test.tictactoe.Play(compPosition)
		if err != test.wantErr {
			t.Error("test", i+1, "failed, expect wantErr", test.wantErr, "got", err)
		}
	}
}
