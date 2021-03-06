package bed2x

import "testing"

var (
	bed1 = &Bed12{"chr1", 0, 100, "noname", float64(0.0), "-", 50, 90, "0,0,0", 3, []int{10, 10, 10}, []int{0, 50, 90}}
	bed2 = &Bed12{"chr1", 0, 100, "noname", float64(0.0), "-", 5, 95, "0,0,0", 3, []int{10, 10, 10}, []int{0, 50, 90}}
)

func TestUtr(t *testing.T) {
	t.Log(bed1)
	t.Log(bed1.CDS())
	t.Log(bed1.UTR3())
	t.Log(bed1.UTR5())
	e, _ := bed1.Exons()
	for _, v := range e {
		t.Log(v)
	}
	introns, _ := bed1.Introns()
	for _, v := range introns {
		t.Log(v)
	}
	t.Log(bed2)
	t.Log(bed2.CDS())
	t.Log(bed2.UTR3())
	t.Log(bed2.UTR5())
	t.Log(Upstream(bed2, 1000))
}

func TestDb(t *testing.T) {
	db := NewDb()
	db.Load("test.bed")
	iter, _ := db.Query("chr1", 1, 20001)
	for v := range iter {
		t.Log(v)
	}
}

func TestDb2(t *testing.T) {
	db := NewDb()
	db.Load("chr1:1-30000")
	db.Load("chr1:200-30000")
	iter, _ := db.Query("chr1", 1, 20001)
	for v := range iter {
		t.Log(v)
	}
}
