package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

/*
Identical to P5.js' map()
*/
func Map(n, start1, stop1 int, start2, stop2 float32) float32 {
	fn, fstart1, fstop1, fstart2, fstop2 := float32(n),
		float32(start1),
		float32(stop1),
		float32(start2),
		float32(stop2)
	return ((fn-fstart1)/(fstop1-fstart1))*(fstop2-fstart2) + fstart2
}

type Image struct {
	img           *image.NRGBA
	Width, Height int
}

func NewImage(size_x, size_y int) *Image {
	return &Image{
		image.NewNRGBA(image.Rect(0, 0, size_x, size_y)),
		size_x,
		size_y,
	}
}

func (i *Image) Encode(outfile string, overrideIfExists bool) error {
	if _, err := os.Stat(outfile); err == nil && !overrideIfExists {
		return fmt.Errorf("file %s exists", outfile)
	}

	f, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer f.Close()

	png.Encode(f, i.img)

	return nil
}

func (i *Image) SetPixel(x, y int, r, g, b, a uint8) {
	i.img.Set(x, y, color.NRGBA{r, g, b, a})
}
