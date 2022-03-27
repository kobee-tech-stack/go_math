package geometry

import "math"

func CubeVolume(n int) int {
	return int(math.Pow(float64(n), 3))
}
