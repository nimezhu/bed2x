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
func sign(i int8) string {
	if i < 0 {
		return "-"
	}
	if i > 0 {
		return "+"
	}
	return "."
}
func SamRecordToBed12(s *sam.Record, chr string) *Bed12 {
	start := s.Start()
	exonStarts, exonLengths := CigarToCoords(s.Cigar, start)
	return &Bed12{
		chr,
		start,
		s.End(),
		s.Name,
		0,
		sign(s.Strand()),
		start,
		s.End(),
		"0",
		len(exonStarts),
		exonStarts,
		exonLengths,
	}
}
