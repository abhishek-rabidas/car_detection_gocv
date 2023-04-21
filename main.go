package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Print("Enter the video footage filename: ")
	var filename string
	fmt.Scanf("%s", &filename)

	footage, err := gocv.VideoCaptureFile("assets/" + filename)

	if err != nil {
		fmt.Println("Error in reading video footage")
		fmt.Println(err)
		return
	}

	window := gocv.NewWindow("SYN")

	mat := gocv.NewMat()

	var flag bool = true

	blue := color.RGBA{0, 0, 255, 0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("data/cars.xml") {
		fmt.Println("Error reading cascade file")
		return
	}

	for flag {
		isTrue := footage.Read(&mat)

		if isTrue {

			rects := classifier.DetectMultiScale(mat)
			for _, r := range rects {
				gocv.Rectangle(&mat, r, blue, 3)
			}
			window.IMShow(mat)
			window.WaitKey(1)
		}

	}

	footage.Close()
	window.Close()

}
