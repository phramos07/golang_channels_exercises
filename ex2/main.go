package main

/*
You are tasked with implementing a simple parallel image processing program. The program takes a list of image file paths as input and applies a grayscale filter to each image concurrently using Go routines and channels. The processed images are then saved to an output directory.

DETAILS:
In the applyGrayscale function:
- you need to iterate over the pixels of the input image and apply the grayscale formula to each pixel. You can use the color.Gray type to set the pixel values in the grayscale image.

In the processImage function:

- Open the input image file using os.Open.
- Apply the grayscale filter using the applyGrayscale function.
- Create the output file using os.Create.
- Encode and save the grayscale image to the output file using the appropriate image format package.

In the main function:
- Iterate over the results channel and print the processing status for each image.
*/

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import necessary image format packages
	_ "image/png"
	"sync"
)

// Define a function to apply a grayscale filter to an image
func applyGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// TODO: Iterate over the image pixels and apply the grayscale formula
	// Tips: iterate over the bounds. bounds.Min.Y and bounds.Max.Y indicate pixel-width and
	// bouns.Min.X and bounds.Max.X indicate pixel-height.

	// Tip: to get one pixel, img.At(x, y)

	// Tip: To convert one pixel to gray: color.GrayModel.Convert(pixel)

	// Tip: To set every pixel on an image: gayImg.Set(x, y, pixel)

	return grayImg
}

// Tips: results channel will have string messages indicating Success or failure
// Sucess: "Processed file at <filepath>"
// Error: "Error when processing <inputfile>"
func processImage(inputPath, outputPath string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// TODO: Open the input image file using os.Open(filename)

	// TODO: The the image object by decoding the file. image.Decode(inputFile)

	// TODO: Apply grayscale filter using the applyGrayscale function

	// TODO: Create the output file
	// It should be in the same folder. os.Create(filepath.Join(outputPath, filepath.Base(inputPath)))

	// TODO: Encode and save the grayscale image
	// Check the extension of the image.
	// if filepath.Ext(inputFile) == ".png" -> png.Encode(outputfile, grayimg)
	// if filepath.Ext(inputFile) == ".jpg" -> jpg.Encode(outputfile, grayimg)

	// TODO: Send a message to the results channel indicating success or failure
}

func main() {
	inputPaths := []string{"file1.jpg", "file2.jpg", "file3.jpg"} // Add image file paths
	outputPath := "output/"                                       // Create the output directory

	results := make(chan string)
	var wg sync.WaitGroup

	for _, inputPath := range inputPaths {
		wg.Add(1)
		go processImage(inputPath, outputPath, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Iterate over the results channel and print the processing status
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("Processing complete.")
}
