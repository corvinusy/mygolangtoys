package main

import (
    "fmt"
    "time"
    "math/rand"
)

const (
    LIMIT = 1e7 // rounds
    GOUT = 0
    JAIL = 10
    G2J = 30
    CC1 = 2
    CC2 = 17
    CC3 = 33
    CH1 = 7
    CH2 = 22
    CH3 = 36
    NEXTR = -1
    R1 = 5
    R2 = 15
    R3 = 25
    R4 = 35
    NEXTU = -2
    U1 = 12
    U2 = 28
    MINUS3 = -3
    STAND = -5
)

func main() {

    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    var a [40]int
    var dbls int
    
    p := 0
    a[0] = 1

    for i := 0; i < LIMIT; i++ {
        d, dbl := rollDices(r)

        if dbl {
            dbls += 1
        } else {
            dbls = 0
        }

        if dbls == 3 {
            p = JAIL
            d = 0
            dbls = 0
        }

        p = (p + d) % 40
        switch p {
        case G2J : p = 10
        case CH1, CH2, CH3 : { 
            switch rollCH(r) {
            case GOUT : p = GOUT
            case JAIL : p = JAIL
            case R1 : p = R1
            case 11 : p = 11
            case 24 : p = 24
            case 39 : p = 39
            case NEXTR: {
                switch {
                case p < R1 : p = R1
                case p < R2 : p = R2
                case p < R3 : p = R3
                case p < R4 : p = R4
                default : p = R1
                }
            } 
            case NEXTU: {
                switch {
                case p < U1 : p = U1
                case p < U2 : p = U2
                default : p = U1
                }               
            }

            case MINUS3: p = (p + 37) % 40
            }
        }
        case CC1, CC2, CC3 : {
            switch rollCC(r) {
            case JAIL : p = JAIL
            case GOUT : p = GOUT
            }
        }
            
        }

        a[p] += 1
    }

    for i := range a {
        if a[i] > 3e5 {
            fmt.Printf("%2d:%d \n", i, a[i])
        }
    }

}
/*----------------------------------------------------------------------------*/
func rollDices(r *rand.Rand) (int, bool) {
    
    d1 := r.Intn(4)
    d2 := r.Intn(4)
    return 2 + d1 + d2, d1 == d2
    
}
/*----------------------------------------------------------------------------*/
func rollCC(r *rand.Rand) int {

    d := r.Intn(16)

    if d == 1 {
        return JAIL
    }

    if d == 2 {
        return GOUT
    }

    return STAND
}
/*----------------------------------------------------------------------------*/
func rollCH(r *rand.Rand) int {

    d := r.Intn(16)

    switch d  {
    case 0: return GOUT
    case 1: return JAIL
    case 2: return 11
    case 3: return 24
    case 4: return 39
    case 5: return 5
    case 6, 7: return NEXTR
    case 8: return NEXTU
    case 9: return MINUS3
    default: return STAND
    }
}
