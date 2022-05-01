package internal

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"reflect"

	"golang.org/x/image/draw"

	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/errors"
)

const (
	downscaledX = 45
	downscaledY = 20
)

type AsciiImageView struct {
	Image      *widgets.Paragraph
	pixelArray []byte
	reversed   bool
	path       string
}

type AsciiImageOptions func(*AsciiImageView)

func WithReversed(reversed bool) AsciiImageOptions {
	return func(i *AsciiImageView) {
		i.reversed = reversed
	}
}

func NewAsciiImageView(path string, opts ...AsciiImageOptions) (*AsciiImageView, error) {
	imgView := &AsciiImageView{}
	for _, opt := range opts {
		opt(imgView)
	}

	imgView.path = path

	imgView.pixelArray = []byte("@80GCLft1i;:,.      ")
	if imgView.reversed {
		imgView.pixelArray = revereByteSlice(imgView.pixelArray)
	}

	ascii, err := imgView.imageFileToASCII()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert image file to ASCII")
	}

	asciiWidget := widgets.NewParagraph()
	asciiWidget.Text = ascii
	asciiWidget.Border = false

	imgView.Image = asciiWidget

	return imgView, nil
}

func (i *AsciiImageView) convertPixel(pixel color.Color) byte {
	r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()

	intensity := (r + g + b) * a / 255
	precision := float64(255 * 3 / (len(i.pixelArray) - 1))
	rawChar := i.pixelArray[roundValue(float64(intensity)/precision)]

	return rawChar
}

func (i *AsciiImageView) imageToASCII(img image.Image) string {
	var ascii string
	downscaledImage := downscaledImage(img, downscaledX, downscaledY)

	for y := 0; y < downscaledImage.Bounds().Max.Y; y++ {
		for x := 0; x < downscaledImage.Bounds().Max.X; x++ {
			pixel := downscaledImage.At(x, y)
			ascii += fmt.Sprintf("%c", i.convertPixel(pixel))
		}
		ascii += "\n"
	}

	return ascii
}

func (i *AsciiImageView) imageFileToASCII() (string, error) {
	img, err := openImage(i.path)
	if err != nil {
		return "", errors.Wrap(err, "Failed to open image file")
	}

	return i.imageToASCII(img), nil
}

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open image file")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode image file")
	}

	return img, nil
}

func downscaledImage(img image.Image, width, height int) image.Image {
	downscaledImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(downscaledImage, downscaledImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return downscaledImage
}

func roundValue(value float64) int {
	return int(math.Floor(value + 0.5))
}

func revereByteSlice(slice []byte) []byte {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
