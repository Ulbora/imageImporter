package main

import (
	"fmt"
	"os"

	ipp "github.com/Ulbora/imageImporter/imgimporter"
)

// ImageImporter ImageImporter
// type ImageImporter struct {
// }

func main() {
	//args := os.Args

	csvFile := os.Args[1:]

	fmt.Print("args: ")
	// fmt.Println(args)
	fmt.Println(csvFile[0])

	var imp ipp.ImageImporter
	impi := imp.New()
	impi.ProcessImages(csvFile[0])
}
