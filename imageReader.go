package imageimporter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
)

// RemoteImage image
type RemoteImage struct {
	// RemoteURL string
	// Name      string
	//ImageData
	// OutputDir string
}

// ReadImage ReadImage
func (m *RemoteImage) ReadImage(name string, remoteURL string) {
	resp, err := http.Get(remoteURL)
	if err != nil {
		log.Fatalf("Error fetching image: %v", err)
	}
	defer resp.Body.Close()

	img, format, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
	}

	outputFile, err := os.Create(name + "." + format) // Use the detected format for the extension
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Encode and write the image to the file based on its format
	switch format {
	case "jpeg":
		err = jpeg.Encode(outputFile, img, nil) // You can pass jpeg.Options here
	case "png":
		err = png.Encode(outputFile, img)
	// Add other cases for different image formats if needed (e.g., "gif")
	default:
		log.Fatalf("Unsupported image format: %s", format)
	}

	if err != nil {
		log.Fatalf("Error encoding and writing image: %v", err)
	}

	fmt.Printf("Image saved successfully as output_image.%s\n", format)

	fmt.Printf("Image format: %s\n", format)
	fmt.Printf("Image dimensions: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	// You can now work with the 'img' variable, which holds the decoded image data.
	// For example, you could save it to a local file, process it, etc.
}
