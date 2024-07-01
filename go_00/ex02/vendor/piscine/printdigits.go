package piscine

import "ft"

func PrintDigits() {
	var x int
	for i := 0; i < 10; i++ {
		x = '0' + i
		ft.PrintRune(rune(x))
	}
	ft.PrintRune('\n')
}
