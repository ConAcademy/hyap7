package main

import "testing"

func TestMinAge(t *testing.T) {
	tests := []struct {
		age  int
		want int
	}{
		{14, 14},
		{18, 16},
		{20, 17},
		{30, 22},
		{40, 27},
		{50, 32},
		{60, 37},
		{80, 47},
		{100, 57},
	}
	for _, tt := range tests {
		got := MinAge(tt.age)
		if got != tt.want {
			t.Errorf("MinAge(%d) = %d, want %d", tt.age, got, tt.want)
		}
	}
}

func TestMaxAge(t *testing.T) {
	tests := []struct {
		age  int
		want int
	}{
		{14, 14},
		{15, 16},
		{18, 22},
		{20, 26},
		{30, 46},
		{40, 66},
		{50, 86},
		{60, 106},
	}
	for _, tt := range tests {
		got := MaxAge(tt.age)
		if got != tt.want {
			t.Errorf("MaxAge(%d) = %d, want %d", tt.age, got, tt.want)
		}
	}
}

func TestRangeWidth(t *testing.T) {
	tests := []struct {
		age  int
		want int
	}{
		{14, 0},
		{16, 3},
		{20, 9},
		{30, 24},
		{40, 39},
		{50, 54},
	}
	for _, tt := range tests {
		got := RangeWidth(tt.age)
		if got != tt.want {
			t.Errorf("RangeWidth(%d) = %d, want %d", tt.age, got, tt.want)
		}
	}
}

func TestCumulativeRange(t *testing.T) {
	tests := []struct {
		age  int
		want float64
	}{
		{14, 0},
		{20, 27},    // 3·400/4 - 21·20 + 147 = 300 - 420 + 147
		{30, 192},   // 3·900/4 - 21·30 + 147 = 675 - 630 + 147
		{40, 507},   // 3·1600/4 - 21·40 + 147 = 1200 - 840 + 147
	}
	for _, tt := range tests {
		got := CumulativeRange(tt.age)
		if got != tt.want {
			t.Errorf("CumulativeRange(%d) = %v, want %v", tt.age, got, tt.want)
		}
	}
}

func TestInRange(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{30, 22, true},  // lower bound
		{30, 46, true},  // upper bound
		{30, 21, false}, // too young
		{30, 47, false}, // too old
		{30, 30, true},  // same age
		{20, 17, true},  // 20's lower bound
		{20, 16, false}, // just below
		{50, 32, true},  // 50's lower bound
		{50, 86, true},  // 50's upper bound
	}
	for _, tt := range tests {
		got := InRange(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("InRange(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
