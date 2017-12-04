package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aebruno/twobit"
	"github.com/nimezhu/bed2x"
	"github.com/nimezhu/netio"
	"github.com/urfave/cli"
)

func CmdSeq(c *cli.Context) error {
	fn, _ := dio(c)
	ch, err := bed2x.IterBedLines(fn)
	checkErr(err)
	format := 12
	signal := false
	genomeUri := c.String("g")
	f, err := netio.Open(genomeUri)
	checkErr(err)
	tb, err := twobit.NewReader(f)
	checkErr(err)
	for line := range ch {
		if signal == false {
			a := strings.Split(line, "\t")
			if len(a) < 12 {
				format = 6
			}
			signal = true
		}
		switch format {
		case 12:
			b, err := bed2x.ParseBed12(line)
			if err != nil {
				log.Println(err)
			} else {
				seq, err := tb.ReadRange(b.Chr(), b.Start(), b.End())
				if err == nil {
					fmt.Println(">%s", b.Id())
					l := 0
					for _, v := range b.BlockSizes() {
						l += v
					}
					s := make([]rune, l)
					seqrunes := []rune(string(seq))
					k := 0
					for i := 0; i < b.BlockCount(); i++ {
						offset := b.BlockStarts()[i]
						for j := 0; j < b.BlockSizes()[i]; j++ {
							s[k] = seqrunes[offset+j]
							k++
						}
					}
					r := string(s)
					if b.Strand() == "-" {
						fmt.Println(bed2x.RC(r))
					} else {
						fmt.Println(r)
					}
				}

			}
		case 6:
			b, err := bed2x.ParseBed6(line)
			if err != nil {
				log.Println(err)
			} else {
				seq, err := tb.ReadRange(b.Chr(), b.Start(), b.End())
				if err == nil {
					fmt.Println(">%s", b.Id())
					if b.Strand() == "-" {
						fmt.Println(bed2x.RC(string(seq)))
					} else {
						fmt.Println(string(seq))
					}
				}
			}
		}
	}
	return nil
}
