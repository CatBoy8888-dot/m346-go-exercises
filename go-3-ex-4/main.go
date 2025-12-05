package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	cards := map[rune][]string{}
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	// TODO: Loop over suits and ranks to output all combinations.

	// for _,rank := range ranks{
	// for _,suit := range suits{
	// 	fmt.Printf("%c%c\t", suit,rank)
	// }
	// fmt.Println()
	// }

	for _, suit := range suits {
    var suitCards []string
    for _, rank := range ranks {
        suitCards = append(suitCards, fmt.Sprintf("%c%c", suit, rank))
    }
    cards[suit] = suitCards
}

	for suit, suitCards := range cards{
		fmt.Printf("%c: %v\n", suit, suitCards)
	}
}

