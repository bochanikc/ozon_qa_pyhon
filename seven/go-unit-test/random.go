package unit

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g
type Dice interface {
	Throw() int
}

type DiceTrue struct{}

func (d DiceTrue) Throw() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)
	fmt.Printf("%d", n)
	return n
}

func IsWin(bet int, dice Dice) bool {
	return dice.Throw() == bet
}
