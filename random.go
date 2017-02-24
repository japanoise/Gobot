package main

import (
	"fmt"
	"math/rand"
)

var piece []string = []string{
	"Pawn", "Bishop", "Knight", "Queen", "Rook", "King", "Princess'",
	"Amazon", "Archbishop", "Unicorn", "Gold General", "Cannon",
}

var placeadj []string = []string{
	"American", "Aussie", "Andalusian", "Belarusian", "Burmese", "Belgian", "Canterbury", "Detroit", "Delhi",
	"Eastern", "Frankish", "Greek", "Haitian", "Indian", "Japanese", "Korean", "London", "Milan", "Newport",
	"Olympic", "Persian", "Qatar", "Russian", "Stalingrad", "Turkish", "Uruguay", "Uxbridge", "Vienna",
	"Western", "Yorkshire", "Zulu",
}

var adj []string = []string{
	"positional", "tactical", "brave", "cowardly", "aggressive", "passive", "odd-looking", "aesthetic",
	"foolish", "confusing", "solid", "drawish", "sharp", "interesting", "questionable", "playable",
	"hypermodern", "strategic", "easy to play", "tricky to play", "amateurish", "slow", "cramped for black",
}

var openterm []string = []string{
	"Attack", "Defence", "Opening", "Game", "Gambit", "Trap", "System",
}

var comment []string = []string{
	"One for the pros", "Rarely played these days", "Common across all levels", "An infamous trick", "Ugly but effective", "Not to be underestimated", "Rarely seen at master level, but common in club games", "The novice's first resort",
	"Not for the faint of heart",
}

func PickString(fr []string) string {
	return fr[rand.Intn(len(fr))]
}

func oname() string {
	retval := "The "
	if Yes() {
		retval += PickString(piece) + "'s"
		if Yes() {
			retval += " Pawn"
		}
		if Yes() {
			retval += " " + PickString(placeadj)
		}
	} else {
		retval += PickString(placeadj)
	}
	return fmt.Sprintf("%s %s.", retval,
		PickString(openterm))
}

func Opening() string {
	return fmt.Sprintf("%s %s. Reputation for being %s but %s.", oname(),
		PickString(comment), PickString(adj), PickString(adj))
}
