package bed2x

import "github.com/biogo/hts/sam"

/* TODO TEST
 */
func CigarToCoords(cigar sam.Cigar, offset int) ([]int, []int) {
	exonStarts := []int{offset}
	iE := 0
	exonLengths := []int{0}
	state := 0
	for _, i := range cigar {
		cigarType := i & 0xf
		cigarLen := int(i >> 4)
		if cigarType == 0 || cigarType == 7 || cigarType == 8 {
			exonLengths[iE] += cigarLen
			state = 1
		} else if cigarType == 2 {
			exonLengths[iE] += cigarLen
			state = 1
		} else if cigarType == 3 {
			if state == 1 {
				exonStarts = append(exonStarts, exonStarts[iE]+cigarLen)
				exonLengths = append(exonLengths, 0)
				iE += 1
			} else {
				exonStarts[iE] += cigarLen

			}
			state = 0
		}
	}
	return exonStarts, exonLengths
}
