package main

import (
    "fmt"
)

func main() {

    const SIZE = 1001

    //create SIZE*SIZE 0-valued array

    spire := make([][]int, SIZE)
    for i := range spire {
        spire[i] = make([]int, SIZE)
        for j := range spire[i] {
            spire[i][j] = 0
        }
    }

    x := SIZE / 2
    y := SIZE / 2

    direction := 0 //0 - right, 1 - down, 2 - left, 3 - up

    for i := 1; i <= SIZE*SIZE; i++ {
        spire[x][y] = i
		if i == SIZE*SIZE {
			break
		}
        switch direction {
        case 0:
            {
                y = y + 1
                if spire[x+1][y] == 0 {
                    direction += 1
                }
            }
        case 1:
            {
                x = x + 1
                if spire[x][y-1] == 0 {
                    direction += 1
                }
            }
        case 2:
            {
                y = y - 1
                if spire[x-1][y] == 0 {
                    direction += 1
                }
            }
        case 3:
            {
                x = x - 1
                if spire[x][y+1] == 0 {
                    direction += 1
                }
            }
        }
        direction %= 4
    }

	sum := -1
    for i := range spire {
		sum = sum + spire[i][i] + spire[i][SIZE-1 - i]
    }

	fmt.Println(sum)

}
/*-----------------------------------------------------------------------------*/
