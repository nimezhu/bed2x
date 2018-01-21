package main

import (
	"fmt"
	"os"

	"github.com/nimezhu/bed2x"
	"github.com/urfave/cli"
)

const (
	VERSION = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "bed2x"
	app.Usage = "bed2x tools"
	app.EnableBashCompletion = true //TODO
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "read",
			Usage:  "get bed",
			Action: CmdRead,
		},
		{
			Name:   "utr3",
			Usage:  "get utr3",
			Action: CmdUTR3,
		},
		{
			Name:   "utr5",
			Usage:  "get utr5",
			Action: CmdUTR5,
		},
		{
			Name:   "cds",
			Usage:  "get cds",
			Action: CmdCDS,
		},
		{
			Name:   "exon",
			Usage:  "get exon",
			Action: CmdExon,
		},
		{
			Name:   "intron",
			Usage:  "get intron",
			Action: CmdIntron,
		},
		{
			Name:   "upstream",
			Usage:  "get upstream",
			Action: CmdUpstream,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "b,bp",
					Usage: "basepair number",
					Value: 1000,
				},
			},
		},
		{
			Name:   "downstream",
			Usage:  "get downstream",
			Action: CmdDownstream,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "b,bp",
					Usage: "basepair number",
					Value: 1000,
				},
			},
		},
		{
			Name:   "promoter",
			Usage:  "get promoter",
			Action: CmdPromoter,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "up,u",
					Usage: "basepair number",
					Value: 1000,
				},
				cli.IntFlag{
					Name:  "down,d",
					Usage: "basepair number",
					Value: 500,
				},
			},
		},
		{
			Name:   "tss",
			Usage:  "get tss",
			Action: CmdTss,
		},
		{
			Name:   "tts",
			Usage:  "get tts",
			Action: CmdTts,
		},
		{
			Name:   "seq",
			Usage:  "get seq",
			Action: CmdSeq,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "g,genome",
					Usage: "2bit file",
				},
			},
		},
		{
			Name:   "querybam",
			Usage:  "query bam",
			Action: CmdQueryBam,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "i,bam",
					Usage: "input bam",
				},
			},
		},
	}
	app.Run(os.Args)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/* read bigbed,gzip and text */
func dio(c *cli.Context) (string, string) {
	input := "STDIN"
	output := "STDOUT"
	if c.NArg() > 0 {
		input = c.Args().Get(0)
	}
	if c.NArg() > 1 {
		output = c.Args().Get(1)
	}
	return input, output
}
func CmdUTR3(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		u, err := b.UTR3()
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdUTR5(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		u, err := b.UTR5()
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdCDS(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		u, err := b.CDS()
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdTss(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed6(in)
	checkErr(err)
	for b := range ch {
		u, err := bed2x.Tss(b)
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdTts(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed6(in)
	checkErr(err)
	for b := range ch {
		u, err := bed2x.Tts(b)
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}

func CmdExon(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		u, err := b.Exons()
		if err == nil {
			for _, e := range u {
				fmt.Println(e)
			}
		}
	}
	return nil
}
func CmdIntron(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed12(in)
	checkErr(err)
	for b := range ch {
		u, err := b.Introns()
		if err == nil {
			for _, i := range u {
				fmt.Println(i)
			}
		}
	}
	return nil
}
func CmdUpstream(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed6(in)
	checkErr(err)
	bp := c.Int("bp")
	for b := range ch {
		u, err := bed2x.Upstream(b, bp)
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdDownstream(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed6(in)
	checkErr(err)
	bp := c.Int("bp")
	for b := range ch {
		u, err := bed2x.Downstream(b, bp)
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdPromoter(c *cli.Context) error {
	in, _ := dio(c)
	ch, err := bed2x.IterBed6(in)
	checkErr(err)
	up := c.Int("up")
	down := c.Int("down")
	for b := range ch {
		u, err := bed2x.Promoter(b, up, down)
		if err == nil {
			fmt.Println(u)
		}
	}
	return nil
}
func CmdRead(c *cli.Context) error {
	fn, _ := dio(c)
	ch, err := bed2x.IterBedLines(fn)
	checkErr(err)
	for line := range ch {
		fmt.Println(line)
	}
	return nil
}
