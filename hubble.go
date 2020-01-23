package main

import (
	"fmt"
	"log"
	"os"
	"image/color"
	"image"
	"image/png"
)


type lineRange struct{
    from int
    to int
}
var ymax, xmax int

func main() {
	// Get file
	file, err := os.Open("hubble.png")
	if err != nil {log.Fatal(err)} // Error case

	defer file.Close() // Close the file t the end of the code

	img, err := png.Decode(file)
	if err != nil {log.Fatal(err)}	// Error case

	b := img.Bounds()			// Get border of the image
	ymax = b.Max.Y				// Number of lines
	xmax = b.Max.X				// Number of colonnes
	imgGray := image.NewGray(b)	// Create the result gray image
	
	nbGoroutine := ymax/200		// Get the number of goroutine -> 1 goroutines threat 200 lines
								// the latest threat the rest
	fmt.Println(ymax)
	fmt.Printf("Number of goroutine : %v\n", nbGoroutine)

	var inputChannel chan lineRange	// Channel containing the range of line to threat for each goroutine
	inputChannel = make (chan lineRange, nbGoroutine+1)
	var outputChannel chan string	// Channel use to wait the end of the goroutines
	outputChannel = make (chan string, nbGoroutine+1)
	
	// Run nbGoroutine + 1 (+1 = the latest lines = ymax%200)
	for goroutine:=0 ; goroutine<nbGoroutine+1 ; goroutine++ {
		go RGBtoGray(inputChannel, outputChannel, img, imgGray)
	}

	// Add the ranges in the channel
	pushnum := 0
    for mcpt:= 0; mcpt < ymax ; mcpt+= 200{
		pushnum ++	// Count nb of channel
		fmt.Printf("goroutine %v\n", mcpt/200)
        toPush := lineRange{from: mcpt, to: mcpt+199}
		inputChannel <- toPush
		if (mcpt == nbGoroutine*200){
			toPush := lineRange{from: mcpt, to: mcpt+ymax%200}
			inputChannel <- toPush
		}
	}
	fmt.Printf("Number of channel %v\n", pushnum)
	
	for i := 0; i < pushnum; i ++{
		<- outputChannel
		fmt.Printf("goroutine %v\n", i)
	}

	outFile, err := os.Create("changed.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	png.Encode(outFile, imgGray)
   
}

func RGBtoGray(inp chan lineRange, feedback chan string, img image.Image, imgGray *image.Gray ) {
	for{
		rng := <-inp
		
		for i:= 0 ; i<xmax ; i++{
			for j:=rng.from ; j<rng.to ; j++ {
				RGBApx := img.At(i,j)
				r, g, b, _:= RGBApx.RGBA()
				gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
				grayPx := color.Gray{uint8(gray / 256)}
				imgGray.Set(i, j, grayPx)

			} 
		}
		feedback <- "FINI"
	}
}
