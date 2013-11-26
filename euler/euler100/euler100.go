package main

import (
    "fmt"
    "math/big"
)

//iteration limit

const LIMIT = 30

func main() {

    z := big.NewRat(1, 2)
    zz := new(big.Rat)
    z1 := new(big.Int)
    z2 := new(big.Int)

    prev_denom := big.NewInt(21)
    blue := big.NewInt(85)
    denom := big.NewInt(120)

    for i := 0; i < LIMIT; i++ {

        newdenom := new(big.Int)
        newblue := new(big.Int)

        newdenom.Mul(denom, denom)
        newblue.Mul(blue, denom)

        newdenom.Div(newdenom, prev_denom)
        newblue.Div(newblue, prev_denom)

        prev_denom.Set(denom)

        fmt.Println("iteration", i, newblue, newdenom)

        //check
        z1.Sub(newblue, big.NewInt(1))
        z1.Mul(z1, newblue)

        z2.Sub(newdenom, big.NewInt(1))
        z2.Mul(z2, newdenom)

        zz.SetFrac(z1, z2)

        for {
            // pair tweaking arrival

            for zz.Cmp(z) == 1 { // if x/y > 1/2 than denom++

                newdenom.Add(newdenom, big.NewInt(1))

                z1.Sub(newblue, big.NewInt(1))
                z1.Mul(z1, newblue)

                z2.Sub(newdenom, big.NewInt(1))
                z2.Mul(z2, newdenom)

                zz.SetFrac(z1, z2)
            }

            for zz.Cmp(z) == -1 { // if x/y < 1/2 than numer++

                newblue.Add(newblue, big.NewInt(1))

                z1.Sub(newblue, big.NewInt(1))
                z1.Mul(z1, newblue)

                z2.Sub(newdenom, big.NewInt(1))
                z2.Mul(z2, newdenom)

                zz.SetFrac(z1, z2)

            }

            if zz.Cmp(z) == 0 {
                fmt.Println("result", newblue, newdenom)

                blue.Set(newblue)
                denom.Set(newdenom)

                if newdenom.Cmp(big.NewInt(1e12)) == 1 {
                    return
                }

                break
            }

        }
    }

}
