package Colors

import (
	"fmt"
	"image"
	_ "image/png" // import this package to decode PNGs
	"os"
)

// test
func GetColor() {
	reader, err := os.Open("downloadedImage.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Open output file for writing
	file, errs := os.Create("colorsPixCount.txt")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			values := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
			file.WriteString(values)

			values = ""
		}
	}
}
