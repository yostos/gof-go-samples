// Strategy pattern is unnecessary in Go. First-class functions replace the Strategy
// interface and concrete classes, passing behavior directly as function values.
package main

import (
	"fmt"
	"math/rand"
)

// Hand represents a janken (rock-paper-scissors) hand.
type Hand int

const (
	Rock     Hand = 0
	Scissors Hand = 1
	Paper    Hand = 2
)

var handNames = [3]string{"グー", "チョキ", "パー"}

func (h Hand) String() string {
	return handNames[h]
}

func GetHand(value int) Hand {
	return Hand(value)
}

// IsStrongerThan returns true if h beats other.
func (h Hand) IsStrongerThan(other Hand) bool {
	return (int(h)+1)%3 == int(other)
}

// StrategyFunc is a function type that returns the next hand and accepts
// a study callback for learning from results.
// nextHand returns the next hand to play.
// study is called after each round with win=true if the player won.
type StrategyFunc func() Hand
type StudyFunc func(win bool)

// NewWinningStrategy returns a strategy that repeats the last hand if it won.
func NewWinningStrategy(seed int64) (nextHand StrategyFunc, study StudyFunc) {
	r := rand.New(rand.NewSource(seed))
	won := false
	prevHand := Hand(0)

	nextHand = func() Hand {
		if !won {
			prevHand = GetHand(r.Intn(3))
		}
		return prevHand
	}
	study = func(win bool) {
		won = win
	}
	return
}

// NewProbStrategy returns a strategy that uses probability based on past history.
func NewProbStrategy(seed int64) (nextHand StrategyFunc, study StudyFunc) {
	r := rand.New(rand.NewSource(seed))
	prevHandValue := 0
	currentHandValue := 0
	history := [3][3]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	getSum := func(hv int) int {
		return history[hv][0] + history[hv][1] + history[hv][2]
	}

	nextHand = func() Hand {
		bet := r.Intn(getSum(currentHandValue))
		handvalue := 0
		if bet < history[currentHandValue][0] {
			handvalue = 0
		} else if bet < history[currentHandValue][0]+history[currentHandValue][1] {
			handvalue = 1
		} else {
			handvalue = 2
		}
		prevHandValue = currentHandValue
		currentHandValue = handvalue
		return GetHand(handvalue)
	}
	study = func(win bool) {
		if win {
			history[prevHandValue][currentHandValue]++
		} else {
			history[prevHandValue][(currentHandValue+1)%3]++
			history[prevHandValue][(currentHandValue+2)%3]++
		}
	}
	return
}

// Player tracks a janken player's name, strategy, and record.
type Player struct {
	name      string
	nextHand  StrategyFunc
	study     StudyFunc
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, nextHand StrategyFunc, study StudyFunc) *Player {
	return &Player{name: name, nextHand: nextHand, study: study}
}

func (p *Player) NextHand() Hand {
	return p.nextHand()
}

func (p *Player) Win() {
	p.study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.gameCount++
}

func (p *Player) String() string {
	return fmt.Sprintf("[%s:%d games, %d win, %d lose]",
		p.name, p.gameCount, p.winCount, p.loseCount)
}

func main() {
	// Create two players with different strategies passed as functions.
	next1, study1 := NewWinningStrategy(314)
	next2, study2 := NewProbStrategy(15)

	player1 := NewPlayer("Taro", next1, study1)
	player2 := NewPlayer("Hana", next2, study2)

	for i := 0; i < 10000; i++ {
		hand1 := player1.NextHand()
		hand2 := player2.NextHand()
		if hand1.IsStrongerThan(hand2) {
			fmt.Printf("Winner:%s\n", player1)
			player1.Win()
			player2.Lose()
		} else if hand2.IsStrongerThan(hand1) {
			fmt.Printf("Winner:%s\n", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}
	fmt.Println("Total result:")
	fmt.Println(player1)
	fmt.Println(player2)
}
