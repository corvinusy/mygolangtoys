package main

import "testing"

var twords []string = readFile("./input.txt")

const tword string = "Copernicus"

func BenchmarkContainString1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		containsString1(twords, tword)
	}
}

func BenchmarkContainString2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		containsString2(twords, tword)
	}
}
