// Memento pattern: in Go, value semantics makes save/restore trivial via struct copy.
// No separate Memento class is needed -- just copy the Gamer struct.
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var fruitNames = []string{"リンゴ", "ぶどう", "バナナ", "みかん"}

// Gamer holds the game state. Since it uses only value types (int, []string copy),
// saving state is just a struct copy -- no Memento class needed.
type Gamer struct {
	Money  int
	Fruits []string
}

// NewGamer creates a gamer with initial money.
func NewGamer(money int) Gamer {
	return Gamer{Money: money}
}

// Bet advances the game by one turn (dice roll).
func (g *Gamer) Bet() {
	dice := rand.Intn(6) + 1
	switch dice {
	case 1:
		g.Money += 100
		fmt.Println("所持金が増えました。")
	case 2:
		g.Money /= 2
		fmt.Println("所持金が半分になりました。")
	case 6:
		f := g.getFruit()
		fmt.Printf("フルーツ(%s)をもらいました。\n", f)
		g.Fruits = append(g.Fruits, f)
	default:
		fmt.Println("何も起こりませんでした。")
	}
}

// Save creates a snapshot of the current state.
// In Go, this is just a struct copy with a filtered copy of the fruits slice.
func (g *Gamer) Save() Gamer {
	snapshot := Gamer{Money: g.Money}
	for _, f := range g.Fruits {
		// Only save "delicious" fruits
		if strings.HasPrefix(f, "おいしい") {
			snapshot.Fruits = append(snapshot.Fruits, f)
		}
	}
	return snapshot
}

// Restore replaces the current state with a saved snapshot.
func (g *Gamer) Restore(snapshot Gamer) {
	g.Money = snapshot.Money
	// Copy the slice to avoid sharing
	g.Fruits = make([]string, len(snapshot.Fruits))
	copy(g.Fruits, snapshot.Fruits)
}

func (g Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.Money, g.Fruits)
}

func (g *Gamer) getFruit() string {
	f := fruitNames[rand.Intn(len(fruitNames))]
	if rand.Intn(2) == 0 {
		return "おいしい" + f
	}
	return f
}

func main() {
	gamer := NewGamer(100)
	snapshot := gamer.Save()

	for i := 0; i < 100; i++ {
		fmt.Printf("==== %d\n", i)
		fmt.Printf("現状:%s\n", gamer)

		gamer.Bet()

		fmt.Printf("所持金は%d円になりました。\n", gamer.Money)

		if gamer.Money > snapshot.Money {
			fmt.Println("※だいぶ増えたので、現在の状態を保存しておこう！")
			snapshot = gamer.Save()
		} else if gamer.Money < snapshot.Money/2 {
			fmt.Println("※だいぶ減ったので、以前の状態を復元しよう！")
			gamer.Restore(snapshot)
		}

		time.Sleep(100 * time.Millisecond)
		fmt.Println()
	}
}
