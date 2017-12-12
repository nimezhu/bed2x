# bed2x
[![Build Status](https://travis-ci.org/nimezhu/bed2x.svg?branch=master)](https://travis-ci.org/nimezhu/bed2x)
[![Releases](https://img.shields.io/github/release/nimezhu/bed2x.svg)](https://github.com/nimezhu/bed2x/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nimezhu/bed2x)](https://goreportcard.com/report/github.com/nimezhu/bed2x)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/nimezhu/bed2x)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/bed2x/current/linux/bed2x)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/bed2x/current/win64/bed2x.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](http://genome.compbio.cs.cmu.edu/~xiaopenz/bed2x/current/mac/bed2x)
## Functions
- get upstream/downstream/intron/exon/promoter/tts/tss/utr3/utr5/cds annotation bed based on input bed12/bed6 bigbed/tabix/gzip/ascii files.
- fetch cDNA sequence for bed12 format file.
- support stdin/stdout pipe
## Install
### Install from go
```
go get github.com/nimezhu/bed2x/...
```
### Download Binaries
[Download Link](http://genome.compbio.cs.cmu.edu/~xiaopenz/bed2x/current)

## Usage Examples
```
bed2x exon [file.bb or file.bed or file.bed.gz]  > file.exon.bed
bed2x promoter file.bb | bed2x seq -g genome.2bit > file.promoter.fa
bed2x read file.bb | head -10 | bed2 seq -g http://hgdownload.soe.ucsc.edu/goldenPath/hg19/bigZips/hg19.2bit
```
