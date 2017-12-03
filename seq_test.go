package bed2x

import "testing"

func TestRC(t *testing.T) {
	seq := "atctt"
	t.Log(seq)
	t.Log(RC(seq))
	seq2 := "aatt"
	t.Log(seq2)
	t.Log(RC(seq2))
}
