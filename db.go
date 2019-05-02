package bed2x

import (
	"errors"
)

type Db struct {
	data     map[string]Bed6i
	binindex *BinIndexMap
	alter    map[string][]Bed6i
}

func NewDb() Db {
	return Db{make(map[string]Bed6i), NewBinIndexMap(), make(map[string][]Bed6i)}
}

func (db Db) Load(fn string) error {
	ch, err := IterBed6(fn)
	if err != nil {
		return err
	}
	for b := range ch {
		if _, ok := db.data[b.Id()]; !ok {
			db.data[b.Id()] = b //TODO: Rename b.Id() if exists
		} else {
			//TODO
			if _, ok := db.alter[b.Id()]; !ok {
				db.alter[b.Id()] = make([]Bed6i, 0, 0)
			}
			db.alter[b.Id()] = append(db.alter[b.Id()], b)
		}
		db.binindex.Insert(b)
	}
	return nil
}

//* TODO Get All Bed6i with same Id */
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

func (db Db) query(chr string, start int, end int) (<-chan NamedRangeI, error) {
	return db.binindex.QueryRegion(chr, start, end)
}
