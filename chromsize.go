package bed2x

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*

$ mysql  --user=genome --host=genome-mysql.cse.ucsc.edu -A -D hg19 -e 'select chrom,size from chromInfo limit 5'
+-------+-----------+
| chrom | size      |
+-------+-----------+
| chr1  | 249250621 |
| chr2  | 243199373 |
| chr3  | 198022430 |
| chr4  | 191154276 |
| chr5  | 180915260 |

*/
type chromInfo struct {
	Chrom string
	Size  int
}

func ChromSizes(genome string) error {
	db, err := gorm.Open("mysql", "genome@(genome-mysql.cse.ucsc.edu)/"+genome+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return err
	}
	rows, err := db.Table("chromInfo").Select("*").Rows() // (*sql.Rows, error)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var r chromInfo
		db.ScanRows(rows, &r)
		fmt.Printf("%s\t%d\n", r.Chrom, r.Size)
	}
	return nil
}
