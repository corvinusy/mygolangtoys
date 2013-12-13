package main

import (
    "fmt"
    "github.com/cznic/mathutil"
)

const THRESHOLD = 100
const LIMIT = 1e9

func main() {

    var mark bool

    count := 1

    for i := uint32(2); i <= LIMIT; i++ {
        fts := mathutil.FactorInt(i)
        if len(fts) > 25 {
            continue
        }

        mark = true
        for _, ft := range fts {
            if ft.Prime > THRESHOLD {
                mark = false
                break
            }
        }

        if mark {
            count += 1
        }
    }

    fmt.Println("limit =", LIMIT, "THRESHOLD =", THRESHOLD, "count = ", count)
    
}
/*----------------------------------------------------------------------------*/
