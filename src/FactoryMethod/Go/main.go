// FactoryMethod: Goではクラス継承ベースのファクトリは不要。
// シンプルなインターフェースとファクトリ関数で同じ「生成と登録」のロジックを実現する。
package main

import "fmt"

// Product is the interface for things created by a factory.
type Product interface {
	Use()
}

// IDCard is a concrete product.
type IDCard struct {
	Owner string
}

func (c *IDCard) Use() {
	fmt.Printf("%sを使います。\n", c)
}

func (c *IDCard) String() string {
	return fmt.Sprintf("[IDCard:%s]", c.Owner)
}

// NewIDCard is a factory function that creates and registers an IDCard.
// In Go, a simple function replaces the entire Factory/FactoryMethod class hierarchy.
func NewIDCard(owner string) *IDCard {
	fmt.Printf("%sのカードを作ります。\n", owner)
	card := &IDCard{Owner: owner}
	fmt.Printf("%sを登録しました。\n", card)
	return card
}

func main() {
	card1 := NewIDCard("Hiroshi Yuki")
	card2 := NewIDCard("Tomura")
	card3 := NewIDCard("Hanako Sato")
	card1.Use()
	card2.Use()
	card3.Use()
}
