package bed2x

import "testing"

func TestRC(t *testing.T) {
	seq := "atctt"
	t.Log("seq", seq)
	t.Log("rc", RC(seq))
}
func TestPosPattern(t *testing.T) {
	c := "chr1:1-2000042324"
	t.Log("?true", posPattern.MatchString(c))
	c2 := "chr1:1-2000042324U"
	t.Log("?false", posPattern.MatchString(c2))
	match := posPattern.FindStringSubmatch(c)
	t.Log("?chr1", match[1])
}
