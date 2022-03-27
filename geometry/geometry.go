package geometry

import (
	"fmt"
	"math"
)

func CubeVolume(n int) int {
	fmt.Println("Function Executed")
	return int(math.Pow(float64(n), 3))
}
