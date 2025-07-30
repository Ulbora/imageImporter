package imgimporter

import (
	"fmt"

	fr "github.com/Ulbora/FileReader"
)

// ImageRecord ImageRecord
type ImageRecord struct {
	Name     string
	ImageURL string
}

// CvsFile CvsFile
type CvsFile struct {
}

// ReadImageURLs ReadImageURLs
func (p *CvsFile) ReadImageURLs(files *fr.CsvFiles) *[]ImageRecord {
	var rtn []ImageRecord
	if files != nil {
		fmt.Println("csv len: ", len(files.CsvFileList))
		for i, cf := range files.CsvFileList {
			if i == 0 {
				continue
			}
			fmt.Println("csv rec: ", cf)

			suc, newimg := p.parseRecord(cf)

			if suc {
				rtn = append(rtn, newimg)
			}

		}
	}

	return &rtn
}

func (p *CvsFile) parseRecord(cf []string) (bool, ImageRecord) {
	var rtn ImageRecord
	var suc = false
	if len(cf) > 0 {
		var n = cf[0]
		rtn.Name = n
		rtn.ImageURL = cf[1]
		suc = true
	}
	return suc, rtn
}

// New New
func (p *CvsFile) New() ImageFileReader {
	return p
}
