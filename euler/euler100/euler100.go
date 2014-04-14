package main

import (
    "fmt"
    "math/big"
)

//iteration limit

const limit = 30

func main() {

    z := big.NewRat(1, 2)
    zz := new(big.Rat)
    z1 := new(big.Int)
    z2 := new(big.Int)

    prevDenom := big.NewInt(21)
    blue := big.NewInt(85)
    denom := big.NewInt(120)

    for i := 0; i < limit; i++ {

        newDenom := new(big.Int)
        newBlue := new(big.Int)

        newDenom.Mul(denom, denom)
        newBlue.Mul(blue, denom)

        newDenom.Div(newDenom, prevDenom)
        newBlue.Div(newBlue, prevDenom)

        prevDenom.Set(denom)

        fmt.Println("iteration", i, newBlue, newDenom)

        //check
        z1.Sub(newBlue, big.NewInt(1))
        z1.Mul(z1, newBlue)

        z2.Sub(newDenom, big.NewInt(1))
        z2.Mul(z2, newDenom)

        zz.SetFrac(z1, z2)

        for {
            // pair tweaking arrival

            for zz.Cmp(z) == 1 { // if x/y > 1/2 than denom++

                newDenom.Add(newDenom, big.NewInt(1))

                z1.Sub(newBlue, big.NewInt(1))
                z1.Mul(z1, newBlue)

                z2.Sub(newDenom, big.NewInt(1))
                z2.Mul(z2, newDenom)

                zz.SetFrac(z1, z2)
            }

            for zz.Cmp(z) == -1 { // if x/y < 1/2 than numer++

                newBlue.Add(newBlue, big.NewInt(1))

                z1.Sub(newBlue, big.NewInt(1))
                z1.Mul(z1, newBlue)

                z2.Sub(newDenom, big.NewInt(1))
                z2.Mul(z2, newDenom)

                zz.SetFrac(z1, z2)

            }

            if zz.Cmp(z) == 0 {
                fmt.Println("result", newBlue, newDenom)

                blue.Set(newBlue)
                denom.Set(newDenom)

                if newDenom.Cmp(big.NewInt(1e12)) == 1 {
                    return
                }

                break
            }

        }
    }

}
