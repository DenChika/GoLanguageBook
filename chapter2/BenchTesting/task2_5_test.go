package BenchTesting

import "testing"

func init() {
	for i := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
}

func popCount5(x uint64) int {
	var sum int
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}

func popCount6(x uint64) int {
	return int(Pc[byte(x>>(0*8))] +
		Pc[byte(x>>(1*8))] +
		Pc[byte(x>>(2*8))] +
		Pc[byte(x>>(3*8))] +
		Pc[byte(x>>(4*8))] +
		Pc[byte(x>>(5*8))] +
		Pc[byte(x>>(6*8))] +
		Pc[byte(x>>(7*8))])
}

func BenchmarkPopCount5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount1(100000)
	}
}

func BenchmarkPopCount6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(100000)
	}
}
