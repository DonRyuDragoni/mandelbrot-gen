package image

import (
	"sync"
)

type MandelbrotSetImage struct {
	Img *Image
}

/*
Checks if a given point is tending twards infinity.
*/
func pointTendsToinfinity(x, y, w, h, maxItersPerPixel int) (n int) {
	z := complex(Map(x, 0, w, -2.5, 1.), Map(y, 0, h, -1., 1.))

	rr := real(z) * real(z)
	ii := imag(z) * imag(z)

	for zi := z; n < maxItersPerPixel && rr+ii < 4; n++ {
		zi = zi*zi + z

		rr, ii = real(zi)*real(zi), imag(zi)*imag(zi)
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

/*
Split the interval [0, `endPoint`] into `length` equaly-spaced points.
*/
func splitCols(endPoint, length int) []int {
	grouplist := make([]int, length)

	for start, i := 0, 0; i < length; i++ {
		grouplist[i] = start
		start += endPoint / length
	}

	return grouplist
}

/*
Basically the same as splitCols, but returns a slice of `[start, end]`
representing the initial and the final points present in each ith interval.
*/
func pointList(endPoint, length int) [][2]int {
	points := make([][2]int, length)
	s := append(splitCols(endPoint, length), endPoint)

	for i := 0; i < len(s)-1; i++ {
		if i == 0 {
			points[i][0] = s[i]
		} else {
			points[i][0] = s[i] + 1
		}

		points[i][1] = s[i+1]
	}

	return points
}

type ColorFn func(int, int, int, int) [4]uint8

func NewMandelbrot(size_x, size_y, maxItersPerPixel int,
	colorer ColorFn, numThreads int) (m *MandelbrotSetImage) {
	m = new(MandelbrotSetImage)
	m.Img = NewImage(size_x, size_y)

	threadPoints := pointList(m.Img.Width, numThreads)

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for _, p := range threadPoints {
		go func(startx, endx int) {
			for x := startx; x <= endx; x++ {
				for y := 0; y <= m.Img.Height; y++ {
					n := pointTendsToinfinity(x, y, m.Img.Width, m.Img.Height,
						maxItersPerPixel)

					color := colorer(x, y, n, maxItersPerPixel)

					m.Img.SetPixel(x, y, color[0], color[1], color[2], color[3])
				}
			}

			wg.Done()
		}(p[0], p[1])
	}

	wg.Wait()

	return
}
