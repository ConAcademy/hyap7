package main

// MinAge returns the youngest age you should date given your age.
// Based on the rule: (age / 2) + 7
func MinAge(age int) int {
	min := age/2 + 7
	if min > age {
		return age
	}
	if min < 14 {
		return 14
	}
	return min
}

// MaxAge returns the oldest age that should date you given your age.
// Based on the inverse: (age - 7) * 2
func MaxAge(age int) int {
	max := (age - 7) * 2
	if max < age {
		return age
	}
	return max
}

// InRange returns true if person of ageA can date person of ageB
// according to the half-your-age-plus-7 rule (checked both ways).
func InRange(ageA, ageB int) bool {
	return ageB >= MinAge(ageA) && ageB <= MaxAge(ageA)
}

// RangeWidth returns the dating range width at a given age: 3*age/2 - 21.
// This is MaxAge(age) - MinAge(age) in the continuous form.
func RangeWidth(age int) int {
	w := 3*age/2 - 21
	if w < 0 {
		return 0
	}
	return w
}

// CumulativeRange returns the integral of range width from age 14 to the given age.
// Represents cumulative person-years of dating eligibility.
// Formula: 3a²/4 - 21a + 147
func CumulativeRange(age int) float64 {
	if age <= 14 {
		return 0
	}
	a := float64(age)
	return 3*a*a/4 - 21*a + 147
}
