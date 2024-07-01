package piscine

import "ft"

//import "fmt"

func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i < 10; i++ {
			ft.PrintRune(rune('0' + i))
			if i != 9 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			} else {
				ft.PrintRune('\n')
			}
		}
	} else {
		for i := 0; i < 10; i++ {
			inner(i, n-1, string('0'+i), i == (10-n))
		}
		ft.PrintRune('\n')
	}
}

func printbuf(buf string) {
	for _, v := range buf {
		ft.PrintRune(v)
	}
}

func inner(start, n int, buf string, last bool) {
	//fmt.Println("lst", last)
	for i := start + 1; i < 10; i++ {

		// base of recursion
		if n == 1 {
			printbuf(buf)
			ft.PrintRune(rune('0' + i))
			if !last || i != 9 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
			// go deeper
		} else {
			inner(i, n-1, buf+string('0'+i), last)
		}
	}
}
