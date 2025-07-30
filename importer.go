package imageimporter

import (
	fr "github.com/Ulbora/FileReader"
)

// Importer  Importer
type Importer interface {
	ReadImages(files *fr.CsvFiles) *[]RemoteImage
}

// ImageReader ImageReader
type ImageReader interface {
	ReadImage(name string, remoteURL string)
}

//go mod init github.com/Ulbora/ImageImporter
