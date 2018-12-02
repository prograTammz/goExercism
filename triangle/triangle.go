// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind uint

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	//checking if the triangle is valid
	//if sum of two sides is not bigger than third, it's unvalid
	if a+b < c {
		k = NaT
		return k
	}
	//condition if the all sides are equal
	if a == b && b == c {
		k = Equ
		return k
	}
	// if there is two equal sides of the triangle
	if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
