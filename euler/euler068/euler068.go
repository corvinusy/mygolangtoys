package main

import (
	"fmt"
	"strconv"
)

/*

    6
      2  7
   1    3
 5  0  4  8
     9


*/

const LIMIT = 10
const RANK = 16

func main() {

	var (
		a      [LIMIT]int
		s      string
		result int = 0
	)

	template := make([][15]int, 10)

	template[9] = [...]int{9, 0, 1, 5, 1, 2, 6, 2, 3, 7, 3, 4, 8, 4, 0}
	template[8] = [...]int{8, 4, 0, 9, 0, 1, 5, 1, 2, 6, 2, 3, 7, 3, 4}
	template[7] = [...]int{7, 3, 4, 8, 4, 0, 9, 0, 1, 5, 1, 2, 6, 2, 3}
	template[6] = [...]int{6, 2, 3, 7, 3, 4, 8, 4, 0, 9, 0, 1, 5, 1, 2}
	template[5] = [...]int{5, 1, 2, 6, 2, 3, 7, 3, 4, 8, 4, 0, 9, 0, 1}

	for total := 10; total < 20; total++ {

		for a[0] = 1; a[0] <= LIMIT; a[0]++ {

			for a[1] = 1; a[1] <= LIMIT; a[1]++ {

				if is_repeated(1, a) {
					continue
				}

				if a[1]+a[0] > total-1 {
					break
				}

				for a[2] = 1; a[2] <= LIMIT; a[2]++ {

					if is_repeated(2, a) {
						continue
					}

					if a[2]+a[1] > total-1 {
						break
					}

					for a[3] = 1; a[3] <= LIMIT; a[3]++ {
						if is_repeated(3, a) {
							continue
						}

						if a[3]+a[2] > total-1 {
							break
						}

						for a[4] = 1; a[4] <= LIMIT; a[4]++ {
							if is_repeated(4, a) {
								continue
							}

							if a[4]+a[3] > total-1 {
								break
							}

							if a[4]+a[0] > total-1 {
								break
							}

							a[5] = total - a[1] - a[2]

							if is_repeated(5, a) {
								continue
							}

							a[6] = total - a[2] - a[3]

							if is_repeated(6, a) {
								continue
							}

							a[7] = total - a[3] - a[4]

							if is_repeated(7, a) {
								continue
							}

							a[8] = total - a[4] - a[0]

							if is_repeated(8, a) {
								continue
							}

							a[9] = total - a[0] - a[1]

							if is_repeated(9, a) {
								continue
							}

							if is_valid(a) {
								s = ""

								// find lowest end index
								l_ind := 5
								for i := 6; i < 10; i++ {
									if a[l_ind] > a[i] {
										l_ind = i
									}
								}

								for i, _ := range template[l_ind] {
									s += strconv.Itoa(a[template[l_ind][i]])
								}
								if len(s) == RANK {
									n, _ := strconv.Atoi(s)
									if result < n {
										result = n
									}
								}
								//fmt.Println(s, "\t", a)
								//break
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("result = ", result)
}

/*-----------------------------------------------------------------------------*/
func is_valid(a [LIMIT]int) bool {

	for i := 0; i < 10; i++ {
		if a[i] < 1 {
			return false
		}

		if a[i] > 10 {
			return false
		}
	}
	return true
}

/*-----------------------------------------------------------------------------*/
func is_repeated(n int, a [LIMIT]int) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if a[i] == a[j] {
				return true
			}
		}
	}
	return false
}

/*-----------------------------------------------------------------------------*/
