package tictac

import (
	"math/rand"

	"github.com/shurcooL/tictactoe"
	"golang.org/x/net/context"
)

var me tictactoe.State

type bestPlayer struct{}

func NewPlayer() (tictactoe.Player, error) {
	return bestPlayer{}, nil
}

func (bestPlayer) Name() string {
	return "Team best"
}

func (bestPlayer) Play(ctx context.Context, b tictactoe.Board) (tictactoe.Move, error) {
	freeSpaces := []tictactoe.State{}
	for i, c := range b.Cells {
		if c == tictactoe.F {
			freeSpaces = append(freeSpaces, c)
			if i != 0 && b.Cells[i-1] == me {
				return tictactoe.Move(i), nil
			} else if i != 9 && b.Cells[i+1] == me {
				return tictactoe.Move(i), nil
			}
		}
	}
	if me == 0 {
		if len(freeSpaces) == 9 {
			me = tictactoe.X
		} else {
			me = tictactoe.O
		}
	}

	return tictactoe.Move(freeSpaces[rand.Intn(len(freeSpaces))]), nil
}
