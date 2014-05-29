package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Sudoku struct {
	lines [9][9]rune
}

func main() {

	var ()

	//read file into source

	content, err := ioutil.ReadFile("sudoku.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content), "\n")

	gridstrs := strings.Split(source, "Grid") // gridstrs = 10 lines of 1 grid

	//create array and init it
	grids := make([]Sudoku, len(gridstrs)-1)

	for k := 1; k < len(gridstrs); k++ {
		strs := strings.Split(gridstrs[k], "\n") // str = 1 line of gridstr
		for i := 1; i < len(strs); i++ {
			for j, a := range strs[i] {
				grids[k-1].lines[i-1][j] = a
			}
		}
	}
	/*
	      //check data
	   	for i, g := range grids {
	   		fmt.Println("Grid", i)
	   		for _, l := range g.lines {
	   			for _, v := range l {
	   				fmt.Printf("%2d", v)
	   			}
	   			fmt.Println()
	   		}
	   	}
	*/
	result := 0
	for _, g := range grids {
		solve(g, &result)
	}

	fmt.Println(result)

}

/*-----------------------------------------------------------------------------*/
func solve(g Sudoku, result *int) bool {

	// create solution matrix

	v, res := get_solution_matrix(g, result)
	if res {
		return true
	}

	// try solutions recursively

	for i, line := range v {
		for j, _ := range line {
			if len(v[i][j]) == 0 {
				break
			}

			if len(v[i][j]) == 1 {
				g.lines[i][j] = v[i][j][0]
			}

			if len(v[i][j]) > 1 {
				for _, a := range v[i][j] {
					g.lines[i][j] = a
					if solve(g, result) {
						return true
					}

				}
			}
		}
	}
	return false
}

/*-----------------------------------------------------------------------------*/
func get_solution_matrix(g Sudoku, result *int) ([9][9][]rune, bool) {
	rmaps := make([]map[rune]bool, 9)
	cmaps := make([]map[rune]bool, 9)
	bmaps := make([]map[rune]bool, 9)

	for i, _ := range g.lines {
		rmaps[i] = make(map[rune]bool, 0)
		cmaps[i] = make(map[rune]bool, 0)
		bmaps[i] = make(map[rune]bool, 0)
	}

	for i, line := range g.lines {
		for j, a := range line {
			if a != '0' {
				rmaps[i][a] = true
				cmaps[j][a] = true
				bmaps[3*(i/3)+j/3][a] = true
			}
		}
	}

	// create available variants for all cells

	var v [9][9][]rune

	sum := 0

	for i, line := range g.lines {
		for j, _ := range line {
			v[i][j] = make([]rune, 0)

			if g.lines[i][j] != '0' {
				v[i][j] = append(v[i][j], g.lines[i][j])
			} else {
				for k := '1'; k <= '9'; k++ {
					if !rmaps[i][k] && !cmaps[j][k] && !bmaps[3*(i/3)+j/3][k] {
						v[i][j] = append(v[i][j], k)
						rmaps[i][k] = true
						cmaps[j][k] = true
						bmaps[3*(i/3)+j/3][k] = true
					}
				}
			}

			if len(v[i][j]) == 0 {
				return v, false
			}

			sum += len(v[i][j])
		}
	}

	if sum == 81 {
		for i, line := range v {
			for j, _ := range line {
				fmt.Printf("%2d", v[i][j][0]-'0')
			}
			fmt.Println()
		}

		*result += int((v[0][0][0]-'0')*100 + (v[0][1][0]-'0')*10 + (v[0][2][0] - '0'))
		fmt.Println("result =", *result)
		return v, true
	}

	return v, false
}
