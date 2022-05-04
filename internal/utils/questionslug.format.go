package utils

//@utils
// Return first n chars of a string
// https://stackoverflow.com/a/41604514/17126147
func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
