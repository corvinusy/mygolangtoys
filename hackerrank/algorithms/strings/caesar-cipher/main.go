package main

import "fmt"

func main() {
	var (
		s    []byte
		n, k byte
	)
	fmt.Scan(&n)
	fmt.Scanln(&s)
	fmt.Scan(&k)
	fmt.Println(encode(s, k))
}

func encode(s []byte, k byte) string {
	for i := range s {
		switch {
		case s[i] >= 'a' && s[i] <= 'z':
			s[i] = (s[i]-'a'+k)%26 + 'a'
		case s[i] >= 'A' && s[i] <= 'Z':
			s[i] = (s[i]-'A'+k)%26 + 'A'
		}
	}
	return string(s)
}
