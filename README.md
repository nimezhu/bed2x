# bed2x
- get upstream/downstream/intron/exon/promoter/tts/tss/utr3/utr5/cds annotation bed based on input bed12/bed6 bigbed/tabix/gzip/ascii files.
- fetch cDNA sequence for bed12 format file.
- support stdin/stdout pipe
## Usage
```
bed2x exon file.bb/file.bed/file.bed.gz  > file.exon.bed
bed2x promoter file.bb | bed2x seq -g genome.2bit > file.promoter.fa
```
