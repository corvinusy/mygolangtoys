package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"sort"
)

type Game struct {
	handN string
	handS string
	comboN byte
	comboS byte
	rankS int
	rankN int
}


func main() {

	content, err := ioutil.ReadFile("poker.txt")
	if err != nil {
		panic("File not read")
	}

	source := string(content)

	lines := strings.Split(source, "\n")

	const (
		MAX = 1
		PAIR = 2
		PAIRS = 3
		TRIPLE = 4
		STREET = 5
		FLASH = 6
		FULLHOUSE = 7
		SQUARE = 8
		FLASHSTREET = 9
		FLASHROYAL = 10
	)

	var games []Game

	for _, l := range lines {
		var game Game
		game.handN = l[:15]
		game.handS = l[15:]
		games = append(games, game)
		find_combo(handN)
	}

}
/*-----------------------------------------------------------------------------*/
func find_combo(h string) string {

	strs := strings.Split(h, " ")
	sort.Strings(strs)

	if (strs[1][1] == strs[4][1]) && (strs[1][1] == strs[7][1]) && 
		(strs[1][1] == strs[10][1]) && (strs[1][1] == strs[13][1]) {
		//one-suiter

		if (strs[0][0] - strs[3][0] == 1) && (strs[3][0] - strs[6][0] == 1) && 
			(strs[6][0] - strs[9][0] == 1) && (strs[9][0] - strs[12][0] == 1) {
			// flash-street
			if strs[0] == 'A' {
				//flash-royal
				return 
			}
		}
	}

}
