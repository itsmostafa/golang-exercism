package triangle

import "math"

// Kind is a triangle type.
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = iota
	// Equ is an equilateral
	Equ
	// Iso is an isosceles
	Iso
	// Sca is a scalene
	Sca
)

// IsTriangle returns whether its a triangle or not.
func IsTriangle(a, b, c float64) bool {
	if a > 0 && b > 0 && c > 0 {
		if a != math.Inf(1) && b != math.Inf(1) && c != math.Inf(1) {
			if (a <= b+c) && (b <= a+c) && (c <= a+b) {
				return true
			}
		}
	}
	return false
}

// IsEqu checks if a triangle is equilateral.
func IsEqu(a, b, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

// IsIso checks if a triangle is an Isosceles.
func IsIso(a, b, c float64) bool {
	if a == b || b == c || c == a && !IsEqu(a, b, c) {
		return true
	}
	return false
}

// IsSca checks if a triangle is Scalene.
func IsSca(a, b, c float64) bool {
	if a != b && a != c && b != c {
		return true
	}
	return false
}

// KindFromSides returns the triangle type based on the sides of its length.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if IsTriangle(a, b, c) == false {
		return NaT
	}

	if IsEqu(a, b, c) == true {
		k = Equ
	} else if IsIso(a, b, c) == true {
		k = Iso
	} else if IsSca(a, b, c) == true {
		k = Sca
	}
	return k
}
