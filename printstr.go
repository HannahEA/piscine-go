package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	schangeable := []rune(s)

	for i := 0; i <= 11; i++ {
		z01.PrintRune(schangeable[i])
	}
	z01.PrintRune('\n')
}
