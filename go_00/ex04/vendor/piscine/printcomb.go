package piscine

import "ft"

func PrintComb() {
	var id, jd, kd int
	id, jd, kd = '0', '0', '0'
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				ft.PrintRune(rune(id + i))
				ft.PrintRune(rune(jd + j))
				ft.PrintRune(rune(kd + k))
				if i != 7 {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
}
