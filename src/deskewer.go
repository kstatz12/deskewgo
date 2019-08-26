package main

import (
	"github.com/google/uuid"
	"gocv.io/x/gocv"
	"image"
	"log"
	"os"
)

func detectAngle(file * os.File) (angle float64) {
	img := gocv.IMRead(file.Name(), 0)
	gocv.BitwiseNot(img, &img)
	gocv.Threshold(img, &img, 0, 255, gocv.ThresholdBinary)
	element := gocv.GetStructuringElement(gocv.MorphRect, image.Point{X: 5, Y: 3})
	gocv.Erode(img, &img, element)
	var points []image.Point
	for x := 0; x < img.Cols(); x++{
		for y := 0; y < img.Rows(); y++{
			if img.GetUCharAt(y, x) > 0 {
				point := image.Point{
					X: x,
					Y: y,
				}
				points = append(points, point)
			}
		}
	}
	box := gocv.MinAreaRect(points)
	angle = box.Angle
	if angle < -45 {
		angle += 90
	}
	return angle
}

func deskew(file * os.File, angle float64) (newFile * os.File) {
	img := gocv.IMRead(file.Name(), 0)

	size := img.Size()

	sz := image.Point{
		X: size[0],
		Y: size[1],
	}

	center := image.Point{
		X: size[0] / 2,
		Y: size[1] / 2,
	}

	m := gocv.GetRotationMatrix2D(center, angle, 1.0)
	rotated := gocv.NewMat()
	gocv.WarpAffine(img, &rotated, m, sz)

	newFile, err := os.Create("/tmp/deskew/"+newFileName())
	if err != nil{
		log.Fatal(err)
	}

	newFile.Write(b)
	return newFile
}


func newFileName() (name string) {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	name = u.String() + ".png"
	return name
}
