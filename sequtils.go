package bed2x

var cMap = map[rune]rune{
	'a': 't',
	't': 'a',
	'c': 'g',
	'g': 'c',
	'n': 'n',
	'A': 'T',
	'T': 'A',
	'G': 'C',
	'C': 'G',
	'N': 'N',
}

func RC(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		runes[i], runes[j] = cMap[runes[j]], cMap[runes[i]]
	}
	return string(runes)
}
