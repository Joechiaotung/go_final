package model

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
)

// Tells if the embedded images are to be used. If false, images from files will be loaded.
const useEmbeddedImages = true

// Image of the board
var BoardImg *image.RGBA
var BoardImgs [1000]*image.RGBA

// Image of the wall block
var WallImg *image.RGBA

// Image of the empty block

// Image of the empty block
var TargetImg *image.RGBA

// Image of a door, this is the exit sign
var ExitImg *image.RGBA

// Image of a congratulation
var WonImg *image.RGBA

func init() {
	// WallImg = loadImg("wall.png", true)
	// WonImg = loadImg("won.png", false)
}

// loadImg loads a PNG image from the specified file, and converts it to image.RGBA and makes sure image has zero Min point.
// This function only used during development as the result contains the images embedded.
// blockSize tells if the image must be of the size of a block (else panics).
func loadImg(name string, blockSize bool) *image.RGBA {
	var data []byte
	var err error

	if useEmbeddedImages {
		data, err = base64.StdEncoding.DecodeString(base64Imgs[name])
	} else {
		data, err = ioutil.ReadFile(name)
	}
	if err != nil {
		panic(err)
	}
	return decodeImg(data, blockSize)
}

// decodeImg decodes an image from the specified data which must be of PNG format.
// blockSize tells if the image must be of the size of a block (else panics).
func decodeImg(data []byte, blockSize bool) *image.RGBA {
	src, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	// Convert to image.RGBA, also make sure result image has zero Min point
	b := src.Bounds()
	if blockSize && (b.Dx() != BlockSize || b.Dy() != BlockSize) {
		panic("Invalid image size!")
	}

	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img, src.Bounds(), src, b.Min, draw.Src)

	return img
}

// printBase64Imgs prints the Base64 encoded strings of the images.
// The printed text is a valid go source format created a map with file names mapped to their base64 encoded contents.
// Used only during development to include those Base64 strings here in the source file
// in order to embed them in the executable native binary.
func printBase64Imgs() {
	var names []string

	names = append(names, "won.png")

	// Generate output
	fmt.Print("var base64Imgs = map[string]string{")
	for i, name := range names {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			panic(err)
		}

		if i > 0 {
			fmt.Print(",")
		}

		fmt.Printf("\n\t\"%s\": \"%s\"", name, base64.StdEncoding.EncodeToString(data))
	}
	fmt.Print("}")
}
