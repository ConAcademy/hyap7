package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type field int

const (
	fieldYourAge field = iota
	fieldPartnerAge
)

type model struct {
	yourAge     int
	partnerAge  int
	activeField field
	width       int
	height      int
	quitting    bool
}

func initialModel() model {
	return model{
		yourAge:     25,
		partnerAge:  0, // 0 means not set
		activeField: fieldYourAge,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "tab", "shift+tab":
			if m.activeField == fieldYourAge {
				m.activeField = fieldPartnerAge
				if m.partnerAge == 0 {
					m.partnerAge = m.yourAge
				}
			} else {
				m.activeField = fieldYourAge
			}

		case "right", "l", "up", "k":
			m = m.incrementAge(1)
		case "left", "h", "down", "j":
			m = m.incrementAge(-1)
		case "shift+right", "L", "K", "shift+up":
			m = m.incrementAge(5)
		case "shift+left", "H", "J", "shift+down":
			m = m.incrementAge(-5)

		case "backspace", "delete":
			if m.activeField == fieldPartnerAge {
				m.partnerAge = 0
				m.activeField = fieldYourAge
			}
		}
	}
	return m, nil
}

func (m model) incrementAge(delta int) model {
	switch m.activeField {
	case fieldYourAge:
		m.yourAge = clamp(m.yourAge+delta, 14, 120)
	case fieldPartnerAge:
		m.partnerAge = clamp(m.partnerAge+delta, 14, 120)
	}
	return m
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Padding(0, 1)

	activeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86"))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	rangeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("214"))

	passStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("82"))

	failStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Half Your Age Plus 7"))
	b.WriteString("\n\n")

	// Your age input
	yourLabel := "  Your age:    "
	yourValue := fmt.Sprintf("[%d]", m.yourAge)
	if m.activeField == fieldYourAge {
		b.WriteString(activeStyle.Render(yourLabel))
		b.WriteString(activeStyle.Render(yourValue))
		b.WriteString(dimStyle.Render("  ◀ ▶"))
	} else {
		b.WriteString(dimStyle.Render(yourLabel))
		b.WriteString(dimStyle.Render(yourValue))
	}
	b.WriteString("\n")

	// Partner age input
	partnerLabel := "  Partner age: "
	if m.partnerAge > 0 {
		partnerValue := fmt.Sprintf("[%d]", m.partnerAge)
		if m.activeField == fieldPartnerAge {
			b.WriteString(activeStyle.Render(partnerLabel))
			b.WriteString(activeStyle.Render(partnerValue))
			b.WriteString(dimStyle.Render("  ◀ ▶  (backspace to clear)"))
		} else {
			b.WriteString(dimStyle.Render(partnerLabel))
			b.WriteString(dimStyle.Render(partnerValue))
		}
	} else {
		b.WriteString(dimStyle.Render(partnerLabel + "(tab to set)"))
	}
	b.WriteString("\n\n")

	// Range display with both formulas
	minA := MinAge(m.yourAge)
	maxA := MaxAge(m.yourAge)
	minFormula := fmt.Sprintf("%d / 2 +  7", m.yourAge)
	maxFormula := fmt.Sprintf("%d × 2 - 14", m.yourAge)
	b.WriteString(dimStyle.Render(fmt.Sprintf("  Youngest you should date:   %10s = ", minFormula)))
	b.WriteString(rangeStyle.Render(fmt.Sprintf("%d", minA)))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render(fmt.Sprintf("  Oldest who should date you: %10s = ", maxFormula)))
	b.WriteString(rangeStyle.Render(fmt.Sprintf("%d", maxA)))
	b.WriteString("\n")

	// Verdict (if partner age is set)
	if m.partnerAge > 0 {
		if InRange(m.yourAge, m.partnerAge) {
			b.WriteString(passStyle.Render(fmt.Sprintf("  ✓ %d is in range!", m.partnerAge)))
		} else {
			diff := 0
			if m.partnerAge < minA {
				diff = minA - m.partnerAge
			} else {
				diff = m.partnerAge - maxA
			}
			b.WriteString(failStyle.Render(fmt.Sprintf("  ✗ %d is out of range by %d year(s)", m.partnerAge, diff)))
		}
		b.WriteString("\n")
	}
	b.WriteString("\n")

	// Chart
	chartWidth := m.width
	if chartWidth <= 0 {
		chartWidth = 70
	}
	chartHeight := m.height - 14
	if chartHeight < 8 {
		chartHeight = 8
	}
	if chartHeight > 30 {
		chartHeight = 30
	}
	b.WriteString(renderChart(chartWidth, chartHeight, m.yourAge, m.partnerAge))
	b.WriteString("\n")

	// Help
	b.WriteString(helpStyle.Render("  ←→/hljk: ±1 age  shift: ±5  tab: switch field  backspace: clear partner  q: quit"))
	b.WriteString("\n")

	v := tea.NewView(b.String())
	v.AltScreen = true
	return v
}
