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

	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/errors"
)

type AsciiImageView struct {
	Image *widgets.Paragraph
}

func NewAsciiImageView(path string) (*AsciiImageView, error) {
	ascii, err := imageFileToASCII(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert image file to ASCII")
	}

	asciiWidget := widgets.NewParagraph()
	asciiWidget.Text = ascii
	asciiWidget.Border = false

	return &AsciiImageView{
		Image: asciiWidget,
	}, nil
}

func convertPixel(pixel color.Color) byte {
	pixels := []byte("Ñ@#W$9876543210?!abc;:+=-,._                    ")

	r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()

	intensity := (r + g + b) * a / 255
	precision := float64(255 * 3 / (len(pixels) - 1))
	rawChar := pixels[roundValue(float64(intensity)/precision)]

	return rawChar
}

func roundValue(value float64) int {
	return int(math.Floor(value + 0.5))
}

func imageToASCII(img image.Image) string {
	var ascii string
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			pixel := img.At(x, y)
			ascii += fmt.Sprintf("%c", convertPixel(pixel))
		}
		ascii += "\n"
	}

	return ascii
}

func imageFileToASCII(path string) (string, error) {
	img, err := openImage(path)
	if err != nil {
		return "", errors.Wrap(err, "Failed to open image file")
	}

	return imageToASCII(img), nil
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
