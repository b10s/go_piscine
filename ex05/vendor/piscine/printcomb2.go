package piscine

import "ft"

func PrintComb2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := i; k < 10; k++ {
				for n := j + 1; n < 10; n++ {
					ft.PrintRune(rune('0' + i))
					ft.PrintRune(rune('0' + j))
					ft.PrintRune(' ')
					ft.PrintRune(rune('0' + k))
					ft.PrintRune(rune('0' + n))
					if i != 9 || j != 8 {
						ft.PrintRune(',')
						ft.PrintRune(' ')
					}
				}
			}
		}
	}
}
