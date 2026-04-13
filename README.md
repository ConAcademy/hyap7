# hya-7

A terminal UI for exploring the "Half Your Age Plus 7" dating age rule.

The folk rule says the youngest person you should date is **(your age / 2) + 7**, and the oldest person who should date you is **(your age - 7) x 2**.

## Usage

```
go run .
```

### Controls

| Key | Action |
|-----|--------|
| `←` `→` / `h` `l` | Adjust age by 1 |
| `Shift+←` `Shift+→` / `H` `L` | Adjust age by 5 |
| `Tab` | Switch between your age and partner age |
| `Backspace` | Clear partner age |
| `q` / `Ctrl+C` | Quit |

### What You See

- **Your age** — adjustable with arrow keys
- **Partner age** — optional, toggle with Tab
- **Acceptable range** — calculated from the rule
- **Verdict** — pass/fail check when a partner age is set
- **Chart** — the acceptable dating zone across all ages, with your position highlighted

The chart shows two curves (min and max acceptable age), a diagonal identity line, and a vertical marker at your current age. If a partner age is set, it appears as a dot on the chart.

## Build

```
go build -o hya-7 .
```

## Test

```
go test ./...
```

## Dependencies

- [BubbleTea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Styling
- [NTCharts](https://github.com/NimbleMarkets/ntcharts) — Terminal charts
