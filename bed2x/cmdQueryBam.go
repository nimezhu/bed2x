package main

import (
	"fmt"

	"github.com/biogo/hts/bam"
	"github.com/biogo/hts/sam"
	"github.com/nimezhu/bed2x"
	"github.com/nimezhu/netio"
	"github.com/urfave/cli"
)

type b interface {
	Start() int
	End() int
}

func overlap(bam b, bed b) bool {
	return bam.End() > bed.Start() && bam.Start() < bed.End()
}
func CmdQueryBam(c *cli.Context) error {
	bamUri := c.String("i")

	baiUri := bamUri + ".bai"
	baiReader, err := netio.Open(baiUri)

	checkErr(err)
	bai, err := bam.ReadIndex(baiReader)
	bamReader, err := netio.Open(bamUri)
	checkErr(err)
	bam1, err := bam.NewReader(bamReader, 0)
	checkErr(err)
	header := bam1.Header()
	refs := header.Refs()
	refMap := make(map[string]*sam.Reference)
	for _, v := range refs {
		refMap[v.Name()] = v
	}

	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		ref := refMap[b.Chr()]
		chunks, err := bai.Chunks(ref, b.Start(), b.End())
		iter, err := bam.NewIterator(bam1, chunks)
		checkErr(err)
		fmt.Printf("QR\t%s\n", b)
		for iter.Next() {
			read := iter.Record()
			if overlap(read, b) {
				//es, el := bed2x.CigarToCoords(reads.Cigar, reads.Start())
				//fmt.Println(es, el)
				bit, _ := read.MarshalText()
				fmt.Printf("HT\t%s\n", string(bit))
				fmt.Println("BED\t", bed2x.SamRecordToBed12(read).String())
			}
		}
	}
	return nil
}
