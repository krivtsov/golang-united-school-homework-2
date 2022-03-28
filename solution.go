package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SideName int

type Brand int

const (
	SidesCircle   SideName = 0
	SidesSquare   SideName = 4
	SidesTriangle SideName = 3
)

func CalcSquare(sideLen float64, sidesNum SideName) float64 {
	switch sidesNum {
	case 4:
		return sideLen * sideLen
	case 3:
		return sideLen * sideLen * math.Sin(math.Pi/3) / 2
	case 0:
		return math.Pi * sideLen * sideLen
	}
	return 0
}
