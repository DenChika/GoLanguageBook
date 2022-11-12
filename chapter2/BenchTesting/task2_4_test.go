package BenchTesting

import "testing"

func init() {
	for i := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
}

func popCount3(x uint64) int {
	var sum int
	for i := 0; i < 64; i++ {
		if last := x; last%2 == 1 {
			sum++
		}
		x = x >> 1
	}
	return sum
}

func popCount4(x uint64) int {
	return int(Pc[byte(x>>(0*8))] +
		Pc[byte(x>>(1*8))] +
		Pc[byte(x>>(2*8))] +
		Pc[byte(x>>(3*8))] +
		Pc[byte(x>>(4*8))] +
		Pc[byte(x>>(5*8))] +
		Pc[byte(x>>(6*8))] +
		Pc[byte(x>>(7*8))])
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount1(100000)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(100000)
	}
}
