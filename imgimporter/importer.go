package imgimporter

import (
	"fmt"
	"io/ioutil"

	fr "github.com/Ulbora/FileReader"
)

// ImageFileReader  ImageFileReader
type ImageFileReader interface {
	ReadImageURLs(files *fr.CsvFiles) *[]ImageRecord
}

// ImageReader ImageReader
type ImageReader interface {
	ReadImage(name string, remoteURL string)
}

// Importer Importer
type Importer interface {
	ProcessImages(file string)
}

// ImageImporter ImageImporter
type ImageImporter struct {
}

// ProcessImages ProcessImages
func (p *ImageImporter) ProcessImages(csvFile string) {
	var cr fr.CsvFileReader
	if csvFile != "" {
		sourceFile, err := ioutil.ReadFile(csvFile)
		fmt.Println("readFile err: ", err)
		rd := cr.GetNew()
		rec := rd.ReadCsvFile(sourceFile)
		fmt.Println("csv err: ", rec.CsvReadErr)
		fmt.Println("csv len: ", len(rec.CsvFileList))
		var rdr CvsFile
		ifr := rdr.New()
		irec := ifr.ReadImageURLs(rec)
		var rimg RemoteImage
		rir := rimg.New()
		for _, im := range *irec {
			rir.ReadImage(im.Name, im.ImageURL)
		}

	}

}

// New New
func (p *ImageImporter) New() Importer {
	return p
}

//go mod init github.com/Ulbora/ImageImporter
