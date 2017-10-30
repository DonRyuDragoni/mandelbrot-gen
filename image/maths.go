package image

/*
Identical to P5.js' map()
*/
func Remap(n, start1, stop1, start2, stop2 float32) float32 {
	return ((n-start1)/(stop1-start1))*(stop2-start2) + start2
}

/*
Checks if a given point is tending twards infinity.
*/
func pointTendsToinfinity(x, y, w, h, maxItersPerPixel int) (n int) {
	z := complex(Remap(
		float32(x),
		0, float32(w),
		-2.5, 1.,
	), Remap(
		float32(y),
		0, float32(h),
		-1., 1.,
	))

	for zi, zz := z, z*z; n < maxItersPerPixel && real(zz)+imag(zz) < 4; n++ {
		zi = zi*zi + z
		zz = zi * zi
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
