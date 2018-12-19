package bed2x

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Bed3i interface {
	Chr() string
	Start() int
	End() int
}
type Bed6i interface {
	Bed3i
	Id() string
	Score() float64
	Strand() string
}
type Bed6 struct {
	chr    string
	start  int
	end    int
	id     string
	score  float64
	strand string
}

func (b *Bed6) Chr() string {
	return b.chr
}
func (b *Bed6) Start() int {
	return b.start
}
func (b *Bed6) End() int {
	return b.end
}
func (b *Bed6) Id() string {
	return b.id
}
func (b *Bed6) Score() float64 {
	return b.score
}
func (b *Bed6) Strand() string {
	return b.strand
}

type Bed12 struct {
	chr         string
	start       int
	end         int
	id          string
	score       float64
	strand      string
	thickStart  int
	thickEnd    int
	itemRgb     string
	blockCount  int
	blockSizes  []int
	blockStarts []int
}

func (b *Bed12) Chr() string {
	return b.chr
}
func (b *Bed12) Start() int {
	return b.start
}
func (b *Bed12) End() int {
	return b.end
}
func (b *Bed12) Id() string {
	return b.id
}
func (b *Bed12) Score() float64 {
	return b.score
}
func (b *Bed12) Strand() string {
	return b.strand
}
func (b *Bed12) ItemRgb() string {
	return b.itemRgb
}
func (b *Bed12) ThickStart() int {
	return b.thickStart
}
func (b *Bed12) ThickEnd() int {
	return b.thickEnd
}
func (b *Bed12) BlockCount() int {
	return b.blockCount
}
func (b *Bed12) BlockSizes() []int {
	return b.blockSizes
}
func (b *Bed12) BlockStarts() []int {
	return b.blockStarts
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func (b *Bed12) _sliceBed12(start int, end int, suffix string) (*Bed12, error) {
	chr := b.chr
	if start < b.start {
		start = b.start
	}
	if end > b.end {
		end = b.end
	}
	strand := b.strand
	id := b.id + "_" + suffix
	score := b.score
	itemRgb := b.itemRgb
	thickStart := max(start, b.thickStart)
	thickEnd := min(end, b.thickEnd)
	blockCount := 0
	sliceBlockStarts := make([]int, 0)
	sliceBlockSizes := make([]int, 0)
	/*init max min */
	sliceStart := end
	sliceEnd := start
	for i := 0; i < b.blockCount; i++ {
		exonStart := b.blockStarts[i] + b.start
		exonEnd := exonStart + b.blockSizes[i]
		exonSliceStart := max(start, exonStart)
		exonSliceEnd := min(end, exonEnd)
		if exonSliceStart < exonSliceEnd {
			if sliceStart > exonSliceStart {
				sliceStart = exonSliceStart
			}
			if sliceEnd < exonSliceEnd {
				sliceEnd = exonSliceEnd
			}
			blockCount += 1
			sliceBlockStarts = append(sliceBlockStarts, exonSliceStart-start)
			sliceBlockSizes = append(sliceBlockSizes, exonSliceEnd-exonSliceStart)
		}
	}
	if start < sliceStart {
		offset := sliceStart - start
		start = sliceStart
		for i := 0; i < blockCount; i++ {
			sliceBlockStarts[i] -= offset
		}
	}
	if end > sliceEnd {
		end = sliceEnd
	}
	if blockCount == 0 {
		return nil, errors.New("wrong slice")
	}
	return &Bed12{chr, start, end, id, score, strand, thickStart, thickEnd, itemRgb, blockCount, sliceBlockSizes, sliceBlockStarts}, nil
}
func (b *Bed12) CDS() (*Bed12, error) {
	return b._sliceBed12(b.thickStart, b.thickEnd, "cds")
}
func (b *Bed12) UTR5() (*Bed12, error) {
	switch b.strand {
	case "+":
		return b._sliceBed12(b.start, b.thickStart, "utr5")
	case "-":
		return b._sliceBed12(b.thickEnd, b.end, "utr5")
	}
	return nil, errors.New("wrong slice")
}
func (b *Bed12) UTR3() (*Bed12, error) {
	switch b.strand {
	case "+":
		return b._sliceBed12(b.thickEnd, b.end, "utr3")
	case "-":
		return b._sliceBed12(b.start, b.thickStart, "utr3")
	}
	return nil, errors.New("wrong slice")
}
func intArray(a []int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}
func (b *Bed12) String() string {
	s := fmt.Sprintf("%s\t%d\t%d", b.chr, b.start, b.end)
	s += fmt.Sprintf("\t%s\t%.1f\t%s", b.id, b.score, b.strand)
	s += fmt.Sprintf("\t%d\t%d\t%s", b.thickStart, b.thickEnd, b.itemRgb)
	s += fmt.Sprintf("\t%d\t%s\t%s", b.blockCount, intArray(b.blockSizes), intArray(b.blockStarts))
	return s
}

func (b *Bed6) String() string {
	s := fmt.Sprintf("%s\t%d\t%d", b.chr, b.start, b.end)
	s += fmt.Sprintf("\t%s\t%.1f\t%s", b.id, b.score, b.strand)
	return s
}

func (b *Bed12) Exons() ([]*Bed6, error) {
	e := make([]*Bed6, b.blockCount)
	step := 1
	j := 0
	if b.strand == "-" {
		step = -1
		j = b.blockCount - 1
	}
	for i := 0; i < b.blockCount; i++ {
		id := fmt.Sprintf("%s_exon_%d", b.id, j+1)
		e[j] = &Bed6{b.chr, b.start + b.blockStarts[i], b.start + b.blockStarts[i] + b.blockSizes[i], id, float64(0.0), b.strand}
		j += step
	}
	return e, nil
}

func (b *Bed6) Exons() ([]*Bed6, error) {
	e := make([]*Bed6, 1)
	id := fmt.Sprintf("%s_exon_1", b.id)
	e[0] = &Bed6{b.chr, b.start, b.end, id, float64(0.0), b.strand}
	return e, nil
}

func (b *Bed12) Introns() ([]*Bed6, error) {
	e := make([]*Bed6, b.blockCount-1)
	step := 1
	j := 0
	if b.strand == "-" {
		step = -1
		j = b.blockCount - 2
	}
	for i := 0; i < b.blockCount-1; i++ {
		id := fmt.Sprintf("%s_intron_%d", b.id, j+1)
		e[j] = &Bed6{b.chr, b.start + b.blockStarts[i] + b.blockSizes[i], b.start + b.blockStarts[i+1], id, float64(0.0), b.strand}
		j += step
	}
	return e, nil
}
func Promoter(b Bed6i, up int, down int) (*Bed6, error) {
	var start int
	var end int
	if b.Strand() == "-" {
		start = b.End() - down
		end = b.End() + up
		if start < 0 {
			down = b.Start()
		}
	} else {
		start = b.Start() - up
		end = b.Start() + down
		if start < 0 {
			up = b.Start()
		}
	}
	id := fmt.Sprintf("%s_promoter_up%d_down%d", b.Id(), up, down)
	return &Bed6{b.Chr(), start, end, id, float64(0.0), b.Strand()}, nil
}
func Upstream(b Bed6i, bp int) (*Bed6, error) {
	var start int
	var end int
	id := fmt.Sprintf("%s_up%d", b.Id(), bp)
	if b.Strand() == "-" {
		start = b.End()
		end = b.End() + bp
	} else {
		start = b.Start() - bp
		end = b.Start()
	}
	if start < 0 {
		start = 0
		id = fmt.Sprintf("%s_up%d", b.Id(), b.Start())
	}
	return &Bed6{b.Chr(), start, end, id, float64(0.0), b.Strand()}, nil

}
func Downstream(b Bed6i, bp int) (*Bed6, error) {
	var start int
	var end int
	id := fmt.Sprintf("%s_down%d", b.Id(), bp)
	if b.Strand() == "-" {
		start = b.Start() - bp
		end = b.Start()
	} else {
		start = b.End()
		end = b.End() + bp
	}
	if start < 0 {
		start = 0
		id = fmt.Sprintf("%s_down%d", b.Id(), b.Start())
	}
	return &Bed6{b.Chr(), start, end, id, float64(0.0), b.Strand()}, nil
}
func Tss(b Bed6i) (*Bed6, error) {
	var pos int
	if b.Strand() == "-" {
		pos = b.End() - 1
	} else {
		pos = b.Start()
	}
	return &Bed6{b.Chr(), pos, pos + 1, b.Id() + "_tss", float64(0.0), b.Strand()}, nil
}
func Tts(b Bed6i) (*Bed6, error) {
	var pos int
	if b.Strand() == "-" {
		pos = b.Start()
	} else {
		pos = b.End() - 1
	}
	return &Bed6{b.Chr(), pos, pos + 1, b.Id() + "_tts", float64(0.0), b.Strand()}, nil
}
