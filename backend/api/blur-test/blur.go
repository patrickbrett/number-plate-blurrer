package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
)

func blur_portion(left int, top int, width int, height int, blurAmount float64) {
	right := left + width
	bottom := top + height

	imgFile1, err := os.Open("testimage.jpg")
	if err != nil {
		fmt.Println(err)
	}

	img1, _, err := image.Decode(imgFile1)
	if err != nil {
		fmt.Println(err)
	}

	// Create new blank image in memory and draw the loaded image onto it
	dstImage := image.NewRGBA(img1.Bounds())
	draw.Draw(dstImage, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)

	// Create blurred version of image separately in memory
	blurred := image.NewRGBA(img1.Bounds())
	graphics.Blur(blurred, img1, &graphics.BlurOptions{StdDev: blurAmount})

	// Cut the desired portion of the blurred image and paste it on top of the copy of the original
	blurBounds := image.Rect(left, top, right, bottom)
	draw.Draw(dstImage, blurBounds, blurred, image.Point{left, top}, draw.Src)

	// Create file to save
	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
	}

	// Save as JPEG file
	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(out, dstImage, &opt)
}

func main() {
	blur_portion(729, 304, 575, 216, 18.0)
}
