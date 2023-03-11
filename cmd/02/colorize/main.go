package main

import (
	"image"
	"image/png"
	"log"
	"mipt-golang-course/tasks/02/colorize"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 3 {
		fatalUsage()
	}
	switch os.Args[1] {
	case "compose":
		compose(os.Args[2])
	case "decompose":
		decompose(os.Args[2])
	default:
		fatalUsage()
	}
}

func fatalUsage() {
	log.Fatal("Usage: ./colorize <compose|decompose> <filename.png>")
}

func compose(filename string) {
	var (
		rFilename, gFilename, bFilename = toColorsFiles(filename)
		decomposed                      colorize.DecomposedImage
	)

	readGray := func(filename string) image.Image {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal("Failed to read file: ", err)
		}
		res, err := png.Decode(f)
		if err != nil {
			log.Fatal("Failed to decode image:", err)
		}
		return res
	}

	decomposed.Red = readGray(rFilename)
	decomposed.Green = readGray(gFilename)
	decomposed.Blue = readGray(bFilename)

	composed := colorize.ComposeRGB(decomposed)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open result file:", err)
	}
	if err := png.Encode(f, composed); err != nil {
		log.Fatal("Failed to save result:", err)
	}
}

func decompose(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal("Failed to decode image: %w", err)
	}

	dec := colorize.DecomposeGRB(img)
	writeGray := func(filename string, img image.Image) {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("Failed to open result file: ", err)
		}
		if err := png.Encode(f, img); err != nil {
			log.Fatal("Failed to encode image:", err)
		}
	}

	rFilename, gFilename, bFilename := toColorsFiles(filename)
	writeGray(rFilename, dec.Red)
	writeGray(gFilename, dec.Green)
	writeGray(bFilename, dec.Blue)
}

func toColorsFiles(filename string) (r string, g string, b string) {
	var (
		ext        = path.Ext(filename)
		withoutExt = filename[:len(filename)-len(ext)]
	)

	return withoutExt + "_red" + ext, withoutExt + "_green" + ext, withoutExt + "_blue" + ext
}
