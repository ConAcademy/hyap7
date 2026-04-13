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
