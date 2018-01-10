package bed2x

import (
	"errors"

	"github.com/nimezhu/data"
)

type Db struct {
	data     map[string]Bed6i
	binindex *data.BinIndexMap
}

func NewDb() Db {
	return Db{make(map[string]Bed6i), data.NewBinIndexMap()}
}

func (db Db) Load(fn string) error {
	ch, err := IterBed6(fn)
	if err != nil {
		return err
	}
	for b := range ch {
		db.data[b.Id()] = b
		//t, _ := Tss(b)
		db.binindex.Insert(b)
	}
	return nil
}
func (db Db) Get(id string) (Bed6i, error) {
	if a, ok := db.data[id]; ok {
		return a, nil
	} else {
		return nil, errors.New("not found")
	}
}
func (db Db) Query(chr string, start int, end int) (<-chan Bed6i, error) {
	a, err := db.binindex.QueryRegion(chr, start, end)
	if err != nil {
		return nil, err
	}
	ch := make(chan Bed6i)
	go func() {
		for v := range a {
			//b, _ := db.Get(v.Id()) //TODO
			ch <- v.(*Bed6)
		}
		close(ch)
	}()
	return ch, nil
}

func (db Db) query(chr string, start int, end int) (<-chan data.NamedRangeI, error) {
	return db.binindex.QueryRegion(chr, start, end)
}
