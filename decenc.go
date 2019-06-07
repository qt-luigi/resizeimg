package resizeimg

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func decgif(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != err {
		return nil, err
	}
	defer f.Close()

	return gif.Decode(f)
}

func encgif(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != err {
		return err
	}
	defer f.Close()

	return gif.Encode(f, img, nil)
}

func decjpg(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != err {
		return nil, err
	}
	defer f.Close()

	return jpeg.Decode(f)
}

func encjpg(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != err {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}

func decpng(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != err {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}

func encpng(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != err {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
