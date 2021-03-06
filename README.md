# bed2x
[![Build Status](https://travis-ci.org/nimezhu/bed2x.svg?branch=master)](https://travis-ci.org/nimezhu/bed2x)
[![Releases](https://img.shields.io/github/release/nimezhu/bed2x.svg)](https://github.com/nimezhu/bed2x/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nimezhu/bed2x)](https://goreportcard.com/report/github.com/nimezhu/bed2x)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/nimezhu/bed2x)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)
## Functions
- get upstream/downstream/intron/exon/promoter/tts/tss/utr3/utr5/cds annotation bed based on input bed12/bed6 bigbed/tabix/gzip/ascii files.
- fetch cDNA sequence for bed12 format file.
- support stdin/stdout pipe
## Install
### Install from go
```
GO111MODULE=on go get github.com/nimezhu/bed2x/...
```
### Download Binaries
Download pre-compiled files for Linux, Windows or Mac OS X from [release](https://github.com/nimezhu/bed2x/releases).

## Usage Examples
```
bed2x exon [file.bb or file.bed or file.bed.gz]  > file.exon.bed
bed2x promoter file.bb | bed2x seq -g genome.2bit > file.promoter.fa
bed2x read file.bb | head -10 | bed2 seq -g http://hgdownload.soe.ucsc.edu/goldenPath/hg19/bigZips/hg19.2bit
bed2x read file.bb | head -10 | bed2 seq -g hg19
```
