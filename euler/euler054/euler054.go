package main

import (
    "fmt"
	"io/ioutil"
	"strings"
)

type Card struct {
	rank int
	suit int
}

type Game struct {
	Nstr string
	Sstr string
	handN [5]Card
	handS [5]Card
	comboN byte
	comboS byte
	rankN int
	rankS int
	winner int
}

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


func main() {

	content, err := ioutil.ReadFile("poker.txt")
	if err != nil {
		panic("File 'poker.txt' not found")
	}

	source := string(content)

	lines := strings.Split(source, "\n")

	var games []Game

	for _, l := range lines {
		var game Game
		if len(l) < 1 {
			break
		}
		game.Nstr = l[:15]
		game.handN = convert2cards(game.Nstr)
		game.Sstr = l[15:]
		game.handS = convert2cards(game.Sstr)
		game.comboN, game.rankN = find_combo(game.handN)
		game.comboS, game.rankS = find_combo(game.handS)
		games = append(games, game)
	}

	count := 0 
	for _, g := range games {
		if (g.rankN < 0) || (g.rankS < 0) {
			fmt.Println(g)
		
			break
		}
		if g.comboN > g.comboS {
			count ++
		}
		if g.comboN == g.comboS {
			if g.rankN > g.rankS {
				count++
			}
			if g.rankN == g.rankS {
				fmt.Println(g)
			}
		}
	}

	fmt.Println(count)

}
/*-----------------------------------------------------------------------------*/
func convert2cards(s string ) [5]Card {
	var cs [5]Card

	for i, a := range s {
		if i%3 == 0 {
			switch a {
			case 'A': cs[i/3].rank = 12
			case 'K': cs[i/3].rank = 11
			case 'Q': cs[i/3].rank = 10
			case 'J': cs[i/3].rank = 9
			case 'T': cs[i/3].rank = 8

			default: cs[i/3].rank = int(a - '2')

			}
		}
		if i%3 == 1 {
			switch a {
			case 'S': cs[i/3].suit = 3
			case 'H': cs[i/3].suit = 2
			case 'D': cs[i/3].suit = 1
			case 'C': cs[i/3].suit = 0
			}
		}
	}

	return cs
}
/*-----------------------------------------------------------------------------*/
func find_combo(h [5]Card) (byte, int) {

	var (
		b bool
		r int
	)

	//check FLASHROYAL

	b, r = is_flashroyal(h)

	if b {
		return FLASHROYAL, r
	}

	//check FLASHSTREET

	b, r = is_flashstreet(h)

	if b {
		return FLASHSTREET, r
	}

	//check SQUARE

	b, r = is_square(h)

	if b {
		return SQUARE, r
	}
	
	//check FULLHOUSE

	b, r = is_fullhouse(h)

	if b {
		return FULLHOUSE, r
	}

	//check FLASH

	b, r = is_flash(h)

	if b {
		return FLASH, r
	}

	//check STREET

	b, r = is_street(h)

	if b {
		return STREET, r
	}

	//check TRIPLE

	b, r = is_triple(h)

	if b {
		return TRIPLE, r
	}

	//check PAIRS

	b, r = is_pairs(h)

	if b {
//		fmt.Println(h, r)
		return PAIRS, r
	}

	//check PAIR

	b, r = is_pair(h)

	if b {
		return PAIR, r
	}

	// check MAX
	b, r = is_max(h)

	if b {
		return MAX, r
	}

	return 0, 0

}
/*-----------------------------------------------------------------------------*/
func is_flashroyal(cs [5]Card) (bool, int) {

	s := cs[0].suit

	rs := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		if n.suit != s {
			return false, -1
		}
		rs[n.rank]++
	}

	if (rs[12] == 1) && (rs[11] == 1) && (rs[10] == 1) && (rs[9] == 1) && (rs[8] == 1) {
		return true, 12
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_flashstreet(cs [5]Card) (bool, int) {

	s := cs[0].suit

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		if n.suit != s {
			return false, -1
		}
		r[n.rank]++ 
	}

	for i:=0; i < 9; i++ {

		if (r[i] == 1) && (r[i+1] == 1) && (r[i+2] == 1) && (r[i+3] == 1) && (r[i+4] == 1) {
			return true, i+4
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_flash(cs [5]Card) (bool, int) {

	s := cs[0].suit

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		if n.suit != s {
			return false, -1
		}
		r[n.rank]++ 
	}

	retc := 0

	for i:=1; i < 13; i++ {

		if (r[i] == 1)  {
			retc = i
		}
	}

	return true, retc
}
/*-----------------------------------------------------------------------------*/
func is_street(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 9; i++ {
		if (r[i] == 1) && (r[i+1] == 1) && (r[i+2] == 1) && (r[i+3] == 1) && (r[i+4] == 1) {
			return true, i+4
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_square(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 13; i++ {
		if (r[i] == 4) {
			return true, i
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_fullhouse(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 13; i++ {
		if (r[i] == 3) {
			for j:=0; j < 13; j++ {
				if r[j] == 2 {
					return true, i
				}
			}
			
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_triple(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 13; i++ {
		if (r[i] == 3) {
			return true, i
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_pairs(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 13; i++ {
		if (r[i] == 2) {
			for j:=i+1; j < 13; j++ {
				if r[j] == 2 {
					return true, j
				}
			}
			
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_pair(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=0; i < 13; i++ {
		if (r[i] == 2) {
			return true, i
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
func is_max(cs [5]Card) (bool, int) {

	r := [...]int {0,0,0,0,0,0,0,0,0,0,0,0,0}

	for _, n := range cs {
		r[n.rank]++
	}

	for i:=12; i >=0; i-- {
		if (r[i] == 1) {
			return true, i
		}
	}
	return false, -1
}
/*-----------------------------------------------------------------------------*/
