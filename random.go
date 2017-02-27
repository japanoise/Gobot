package main

import (
	"fmt"
	"math/rand"
)

/* Random chess opening */

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

/* Random I Ching hexagram */
var iching []string = []string{
	"\u4DC0 - HEXAGRAM FOR THE CREATIVE HEAVEN",
	"\u4DC1 - HEXAGRAM FOR THE RECEPTIVE EARTH",
	"\u4DC2 - HEXAGRAM FOR DIFFICULTY AT THE BEGINNING",
	"\u4DC3 - HEXAGRAM FOR YOUTHFUL FOLLY",
	"\u4DC4 - HEXAGRAM FOR WAITING",
	"\u4DC5 - HEXAGRAM FOR CONFLICT",
	"\u4DC6 - HEXAGRAM FOR THE ARMY",
	"\u4DC7 - HEXAGRAM FOR HOLDING TOGETHER",
	"\u4DC8 - HEXAGRAM FOR SMALL TAMING",
	"\u4DC9 - HEXAGRAM FOR TREADING",
	"\u4DCA - HEXAGRAM FOR PEACE",
	"\u4DCB - HEXAGRAM FOR STANDSTILL",
	"\u4DCC - HEXAGRAM FOR FELLOWSHIP",
	"\u4DCD - HEXAGRAM FOR GREAT POSSESSION",
	"\u4DCE - HEXAGRAM FOR MODESTY",
	"\u4DCF - HEXAGRAM FOR ENTHUSIASM",
	"\u4DD0 - HEXAGRAM FOR FOLLOWING",
	"\u4DD1 - HEXAGRAM FOR WORK ON THE DECAYED",
	"\u4DD2 - HEXAGRAM FOR APPROACH",
	"\u4DD3 - HEXAGRAM FOR CONTEMPLATION",
	"\u4DD4 - HEXAGRAM FOR BITING THROUGH",
	"\u4DD5 - HEXAGRAM FOR GRACE",
	"\u4DD6 - HEXAGRAM FOR SPLITTING APART",
	"\u4DD7 - HEXAGRAM FOR RETURN",
	"\u4DD8 - HEXAGRAM FOR INNOCENCE",
	"\u4DD9 - HEXAGRAM FOR GREAT TAMING",
	"\u4DDA - HEXAGRAM FOR MOUTH CORNERS",
	"\u4DDB - HEXAGRAM FOR GREAT PREPONDERANCE",
	"\u4DDC - HEXAGRAM FOR THE ABYSMAL WATER",
	"\u4DDD - HEXAGRAM FOR THE CLINGING FIRE",
	"\u4DDE - HEXAGRAM FOR INFLUENCE",
	"\u4DDF - HEXAGRAM FOR DURATION",
	"\u4DE0 - HEXAGRAM FOR RETREAT",
	"\u4DE1 - HEXAGRAM FOR GREAT POWER",
	"\u4DE2 - HEXAGRAM FOR PROGRESS",
	"\u4DE3 - HEXAGRAM FOR DARKENING OF THE LIGHT",
	"\u4DE4 - HEXAGRAM FOR THE FAMILY",
	"\u4DE5 - HEXAGRAM FOR OPPOSITION",
	"\u4DE6 - HEXAGRAM FOR OBSTRUCTION",
	"\u4DE7 - HEXAGRAM FOR DELIVERANCE",
	"\u4DE8 - HEXAGRAM FOR DECREASE",
	"\u4DE9 - HEXAGRAM FOR INCREASE",
	"\u4DEA - HEXAGRAM FOR BREAKTHROUGH",
	"\u4DEB - HEXAGRAM FOR COMING TO MEET",
	"\u4DEC - HEXAGRAM FOR GATHERING TOGETHER",
	"\u4DED - HEXAGRAM FOR PUSHING UPWARD",
	"\u4DEE - HEXAGRAM FOR OPPRESSION",
	"\u4DEF - HEXAGRAM FOR THE WELL",
	"\u4DF0 - HEXAGRAM FOR REVOLUTION",
	"\u4DF1 - HEXAGRAM FOR THE CAULDRON",
	"\u4DF2 - HEXAGRAM FOR THE AROUSING THUNDER",
	"\u4DF3 - HEXAGRAM FOR THE KEEPING STILL MOUNTAIN",
	"\u4DF4 - HEXAGRAM FOR DEVELOPMENT",
	"\u4DF5 - HEXAGRAM FOR THE MARRYING MAIDEN",
	"\u4DF6 - HEXAGRAM FOR ABUNDANCE",
	"\u4DF7 - HEXAGRAM FOR THE WANDERER",
	"\u4DF8 - HEXAGRAM FOR THE GENTLE WIND",
	"\u4DF9 - HEXAGRAM FOR THE JOYOUS LAKE",
	"\u4DFA - HEXAGRAM FOR DISPERSION",
	"\u4DFB - HEXAGRAM FOR LIMITATION",
	"\u4DFC - HEXAGRAM FOR INNER TRUTH",
	"\u4DFD - HEXAGRAM FOR SMALL PREPONDERANCE",
	"\u4DFE - HEXAGRAM FOR AFTER COMPLETION",
	"\u4DFF - HEXAGRAM FOR BEFORE COMPLETION",
}

func IChing() string {
	return PickString(iching)
}
