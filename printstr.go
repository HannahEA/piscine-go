package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	
 
	for _, rune := range s{
		z01.PrintRune(rune)
	}
	
}
