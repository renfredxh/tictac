package tictac

import (
	"math/rand"

	"github.com/shurcooL/tictactoe"
	"golang.org/x/net/context"
)

type bestPlayer struct{}

func NewPlayer() (tictactoe.Player, error) {
	return bestPlayer{}, nil
}

func (bestPlayer) Name() string {
	return "Team best"
}

func (bestPlayer) Play(ctx context.Context, b tictactoe.Board) (tictactoe.Move, error) {
	freeSpaces := []tictactoe.State{}
	for _, c := range b.Cells {
		if c == tictactoe.F {
			freeSpaces = append(freeSpaces, c)
		}
	}
	return tictactoe.Move(freeSpaces[rand.Intn(len(freeSpaces))]), nil
}
