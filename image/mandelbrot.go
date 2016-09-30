package image

type MandelbrotSetImage struct {
	Img *Image
}

/*
Returns the absolute value of a number.
*/
func abs(n float32) float32 {
	if n < float32(0) {
		return -n
	} else {
		return n
	}
}

/*
Checks if a given point is tending twards infinity.
*/
func pointTendsToinfinity(x, y, w, h, maxItersPerPixel int) (n int) {
	maxval := float32(1.4)
	minval := -maxval

	z := complex(Map(x, 0, w, minval, maxval), Map(y, 0, h, minval, maxval))

	zi := z

	for ; n < maxItersPerPixel; n++ {
		zi = zi*zi + z

		if abs(real(zi)+imag(zi)) > 20 {
			break
		}
	}

	return
}

/*
Default function used to generate the color of each pixel.

x                -> horizontal coordinate of the pixel
y                -> vertical coordinate of the pixel
itersDone        -> number of iterations
maxItersPerPixel -> maximum allowed number of iterations per pixel

colors -> array containing the RGBA values for the given pixel
*/
func DefaultColorFn(x, y, itersDone, maxItersPerPixel int) (colors [4]uint8) {
	if itersDone != maxItersPerPixel {
		brightness := uint8(Map(itersDone, 0, maxItersPerPixel, 0, 255))
		colors[0] = brightness
		colors[1] = brightness
		colors[2] = brightness
		colors[3] = 255
	} else {
		colors[0] = 0
		colors[1] = 0
		colors[2] = 0
		colors[3] = 255
	}

	return
}

func NewMandelbrot(size_x, size_y, maxItersPerPixel int, colorFn func(int, int,
	int, int) [4]uint8) (m *MandelbrotSetImage) {
	m = new(MandelbrotSetImage)
	m.Img = NewImage(size_x, size_y)

	for x := 0; x <= m.Img.Width; x++ {
		for y := 0; y <= m.Img.Height; y++ {
			n := pointTendsToinfinity(x, y, m.Img.Width, m.Img.Height,
				maxItersPerPixel)

			color := colorFn(x, y, n, maxItersPerPixel)

			m.Img.SetPixel(x, y, color[0], color[1], color[2], color[3])
		}
	}

	return
}
