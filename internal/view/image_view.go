package internal

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/errors"
)

type ImageView struct {
	Image       *widgets.Image
	imagePath   string
	imageBase64 string
}

type ImageOptions func(*ImageView)

func WithImagePath(imgPath string) ImageOptions {
	return func(imgView *ImageView) {
		imgView.imagePath = imgPath
	}
}

func WithImageBase64(imgBase64 string) ImageOptions {
	return func(imgView *ImageView) {
		imgView.imageBase64 = imgBase64
	}
}

func NewImageView(otps ...ImageOptions) (*ImageView, error) {
	imgView := &ImageView{}

	for _, opt := range otps {
		opt(imgView)
	}

	var img *widgets.Image
	if imgView.imagePath != "" {
		imageBase64, err := readImageBase64(imgView.imagePath)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read image file")
		}

		image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageBase64)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load %v image: %v \n", imgView.imagePath, err)
			spew.Dump(imageBase64)
			os.Exit(1)
		}

		img = widgets.NewImage(image)
	}

	if imgView.imageBase64 != "" {
		image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgView.imageBase64)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load gopher image: %v \n", err)
			os.Exit(1)
		}

		img = widgets.NewImage(image)
	}

	imgView.Image = img
	return imgView, nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func readImageBase64(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read image file")
	}

	var base64Encoding string
	// mimeType := http.DetectContentType(bytes)

	// switch mimeType {
	// case "image/jpeg":
	// 	base64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	base64Encoding += "data:image/png;base64,"
	// }

	base64Encoding += toBase64(bytes)

	return base64Encoding, nil
}
