package printreversealphabet

import "ft"

func PrintReverseAlphabet() {
	var x int
	for i := 0; i < 26; i++ {
		x = 'z' - i
		ft.PrintRune(rune(x))
	}
	ft.PrintRune('\n')
}
