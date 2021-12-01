package piscine

func StrRev(s string) string {
	r := []rune(s)
	i := len([]rune(s))
	for _, c := range s {
		i--
		r[i] = c
	}
	return string(r)
}
