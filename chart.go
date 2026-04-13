package main

import (
	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/linechart"
)

var (
	chartAxisStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	chartLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	minLineStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("82"))  // green
	maxLineStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // orange
	markerStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("205")) // pink
	partnerMarker   = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))  // cyan
)

const (
	ageMin = 14.0
	ageMax = 80.0
)

// renderChart draws the acceptable-range chart with the user's age highlighted.
func renderChart(width, height, yourAge, partnerAge int) string {
	// Reserve some space for padding
	w := width - 4
	if w < 20 {
		w = 20
	}

	lc := linechart.New(
		w, height,
		ageMin, ageMax, // X range: ages
		ageMin, ageMax+40, // Y range: acceptable partner ages
		linechart.WithXYSteps(2, 2),
		linechart.WithStyles(chartAxisStyle, chartLabelStyle, chartAxisStyle),
	)

	lc.DrawXYAxisAndLabel()

	// Draw the min line: y = x/2 + 7
	for age := ageMin; age <= ageMax; age += 0.5 {
		next := age + 0.5
		y1 := age/2.0 + 7.0
		y2 := next/2.0 + 7.0
		p1 := canvas.Float64Point{X: age, Y: y1}
		p2 := canvas.Float64Point{X: next, Y: y2}
		lc.DrawBrailleLineWithStyle(p1, p2, minLineStyle)
	}

	// Draw the max line: y = (x - 7) * 2
	for age := ageMin; age <= ageMax; age += 0.5 {
		next := age + 0.5
		y1 := (age - 7.0) * 2.0
		y2 := (next - 7.0) * 2.0
		p1 := canvas.Float64Point{X: age, Y: y1}
		p2 := canvas.Float64Point{X: next, Y: y2}
		lc.DrawBrailleLineWithStyle(p1, p2, maxLineStyle)
	}

	// Draw the "you = you" diagonal (identity line) as reference
	for age := ageMin; age <= ageMax; age += 0.5 {
		next := age + 0.5
		p1 := canvas.Float64Point{X: age, Y: age}
		p2 := canvas.Float64Point{X: next, Y: next}
		lc.DrawBrailleLineWithStyle(p1, p2, chartAxisStyle)
	}

	// Draw vertical marker at your age
	ya := float64(yourAge)
	if ya >= ageMin && ya <= ageMax {
		minY := ya/2.0 + 7.0
		maxY := (ya - 7.0) * 2.0
		bottom := canvas.Float64Point{X: ya, Y: minY}
		top := canvas.Float64Point{X: ya, Y: maxY}
		lc.DrawBrailleLineWithStyle(bottom, top, markerStyle)
	}

	// Draw partner marker if set
	if partnerAge > 0 {
		pa := float64(partnerAge)
		lc.DrawRuneWithStyle(
			canvas.Float64Point{X: ya, Y: pa},
			'●',
			partnerMarker,
		)
	}

	return lc.View()
}
