package triangle

import (
	"math"
	"sort"
)

const testVersion = 2

// KindFromSides return the triangle type
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(sides))

	switch {
	case math.IsNaN(sides[0]) || sides[0] <= 0 || sides[2] == math.Inf(1) || sides[2] > sides[0]+sides[1]:
		return NaT
	case sides[0] == sides[1] && sides[0] == sides[2] && sides[1] == sides[2]:
		return Equ
	case sides[0] == sides[1] || sides[1] == sides[2] || sides[0] == sides[2]:
		return Iso
	}

	return Sca
}

// Kind type
type Kind int

// Triangle's type
const (
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
	NaT = iota // not a triangle
)
