// Пакет popcount считает количество единиц в битовом представлении числа.
package popcount

var Pc [256]byte

func init() {
	for i := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(Pc[byte(x>>(0*8))] +
		Pc[byte(x>>(1*8))] +
		Pc[byte(x>>(2*8))] +
		Pc[byte(x>>(3*8))] +
		Pc[byte(x>>(4*8))] +
		Pc[byte(x>>(5*8))] +
		Pc[byte(x>>(6*8))] +
		Pc[byte(x>>(7*8))])
}
