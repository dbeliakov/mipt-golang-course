package main

import (
	"log"
	"os"
	"time"
	"image/color"

	"github.com/dbeliakov/mipt-golang-course/tasks/02/timepng"
)

func main() {
	file, err := os.Create("time.png")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	timepng.TimePNG(file, time.Now(), color.RGBA{
		R: 100,
		G: 100,
		B: 255,
		A:255,
	}, 10)
}
