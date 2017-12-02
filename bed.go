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
	Name() string
	Score() float64
	Strand() string
}
type Bed6 struct {
	chr    string
	start  int
	end    int
	name   string
	score  float64
	strand string
}
type Bed12 struct {
	chr         string
	start       int
	end         int
	name        string
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
func (b *Bed12) Name() string {
	return b.name
}
func (b *Bed12) Score() float64 {
	return b.score
}
func (b *Bed12) Strand() string {
	return b.strand
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
	name := b.name + "_" + suffix
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
	return &Bed12{chr, start, end, name, score, strand, thickStart, thickEnd, itemRgb, blockCount, sliceBlockSizes, sliceBlockStarts}, nil
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
	s += fmt.Sprintf("\t%s\t%.1f\t%s", b.name, b.score, b.strand)
	s += fmt.Sprintf("\t%d\t%d\t%s", b.thickStart, b.thickEnd, b.itemRgb)
	s += fmt.Sprintf("\t%d\t%s\t%s", b.blockCount, intArray(b.blockSizes), intArray(b.blockStarts))
	return s
}

func (b *Bed6) String() string {
	s := fmt.Sprintf("%s\t%d\t%d", b.chr, b.start, b.end)
	s += fmt.Sprintf("\t%s\t%.1f\t%s", b.name, b.score, b.strand)
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
		name := fmt.Sprintf("%s_Exon_%d", b.name, j+1)
		e[j] = &Bed6{b.chr, b.start + b.blockStarts[i], b.start + b.blockStarts[i] + b.blockSizes[i], name, float64(0.0), b.strand}
		j += step
	}
	return e, nil
}

func (b *Bed6) Exons() ([]*Bed6, error) {
	e := make([]*Bed6, 1)
	name := fmt.Sprintf("%s_Exon_1", b.name)
	e[0] = &Bed6{b.chr, b.start, b.end, name, float64(0.0), b.strand}
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
		name := fmt.Sprintf("%s_Intron_%d", b.name, j+1)
		e[j] = &Bed6{b.chr, b.start + b.blockStarts[i] + b.blockSizes[i], b.start + b.blockStarts[i+1], name, float64(0.0), b.strand}
		j += step
	}
	return e, nil
}
