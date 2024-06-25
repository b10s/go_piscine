package printalphabet

import "ft"

func PrintAlphabet() {
	var x int
	for i := 0; i < 26; i++ {
		x = 'a' + i
		ft.PrintRune(rune(x))
	}
	ft.PrintRune('\n')
}
