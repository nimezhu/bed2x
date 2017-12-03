package bed2x

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nimezhu/indexed"
	"github.com/nimezhu/indexed/bbi"
	"github.com/nimezhu/netio"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func fillLinesBigBed(r io.ReadSeeker, ch chan string) error {
	bwf := bbi.NewBbiReader(r)
	bwf.InitIndex() //smart initindex??
	bw := bbi.NewBigBedReader(bwf)
	iter := bw.Iter()
	idx2chr := make(map[int]string)
	for chr, idx := range bw.Genome.Chr2Idx {
		idx2chr[idx] = chr
	}
	for b := range iter {
		ch <- fmt.Sprintf("%s\t%d\t%d\t%s", idx2chr[b.Idx], b.From, b.To, b.Value)
	}
	return nil
}
func fillLines(r io.Reader, ch chan string) error {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	lines := string(c)
	l := strings.Split(lines, "\n")
	for _, v := range l {
		ch <- v
	}
	return nil
}

/*IterBedLines: input format could be bigbed,gzip and ascii text file
 */
func IterBedLines(fn string) (<-chan string, error) {
	ch := make(chan string)
	if fn == "STDIN" {
		bits, _ := ioutil.ReadAll(os.Stdin)
		lines := strings.Split(string(bits), "\n")
		go func() {
			for _, v := range lines {
				ch <- v
			}
			close(ch)
		}()
	} else {
		f, err := netio.Open(fn)
		if err != nil {
			return nil, err
		}
		format, _ := indexed.MagicReadSeeker(f)
		go func() {
			if format == "bigbed" {
				fillLinesBigBed(f, ch)
			} else if format == "gzip" {
				g, _ := gzip.NewReader(f)
				fillLines(g, ch)
			} else if format == "unknown" {
				fillLines(f, ch)
			}
			close(ch)
		}()
	}
	return ch, nil
}
func IterBed12(fn string) (<-chan *Bed12, error) {
	lines, err := IterBedLines(fn)
	if err != nil {
		return nil, err
	}
	ch := make(chan *Bed12)
	go func() {
		for line := range lines {
			b, err := ParseBed12(line)
			if err == nil {
				ch <- b
			} else {
				log.Println(err)
			}
		}
		close(ch)
	}()
	return ch, nil
}
func IterBed6(fn string) (<-chan *Bed6, error) {
	lines, err := IterBedLines(fn)
	if err != nil {
		return nil, err
	}
	ch := make(chan *Bed6)
	go func() {
		for line := range lines {
			b, err := ParseBed6(line)
			if err == nil {
				ch <- b
			} else {
				log.Println(err)
			}
		}
		close(ch)
	}()
	return ch, nil
}

var errorFormat = errors.New("wrong bed format")

func parseInts(a string) ([]int, error) {
	a = strings.Trim(a, ",")
	b := strings.Split(a, ",")
	r := make([]int, len(b))
	var err error
	for i, v := range b {
		r[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
func ParseBed6(line string) (*Bed6, error) {
	a := strings.Split(line, "\t")
	if len(a) < 6 {
		return nil, errors.New("less than 6 column")
	}
	chr := a[0]
	start, err := strconv.Atoi(a[1])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(a[2])
	if err != nil {
		return nil, err
	}
	name := a[3]
	score, err := strconv.ParseFloat(a[4], 64)
	if err != nil {
		return nil, err
	}
	strand := a[5]
	return &Bed6{chr, start, end, name, score, strand}, nil
}
func ParseBed12(line string) (*Bed12, error) {
	a := strings.Split(line, "\t")
	if len(a) < 12 {
		return nil, errors.New("less than 12 column")
	}
	chr := a[0]
	start, err := strconv.Atoi(a[1])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(a[2])
	if err != nil {
		return nil, err
	}
	name := a[3]
	score, err := strconv.ParseFloat(a[4], 64)
	if err != nil {
		return nil, err
	}
	strand := a[5]
	thickStart, err := strconv.Atoi(a[6])
	if err != nil {
		return nil, err
	}
	thickEnd, err := strconv.Atoi(a[7])
	if err != nil {
		return nil, err
	}
	itemRgb := a[8]
	blockCount, err := strconv.Atoi(a[9])
	if err != nil {
		return nil, err
	}
	blockSizes, err := parseInts(a[10])
	if err != nil {
		return nil, err
	}
	blockStarts, err := parseInts(a[11])
	if err != nil {
		return nil, err
	}
	return &Bed12{chr, start, end, name, score, strand, thickStart, thickEnd, itemRgb, blockCount, blockSizes, blockStarts}, nil
}