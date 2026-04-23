# hyap7

`hyap7` is a TUI for working with the quantitative ethical question "Half your age plus 7"

It is a Golang application using [BubbleTea v2](https://charm.land/bubbletea) for TUI and [NTCharts v2](https://github.com/NimbleMarkets/ntcharts) for charting, with [Lip Gloss v2](https://charm.land/lipgloss) for styling.

## Build & Run

Native binary:

```
go build -o hyap7 .
./hyap7
```

Or directly: `go run .`

Web (WASM) build for GitHub Pages:

```
go tool booba-wasm-build -o web/app.wasm .
go tool booba-assets web/
```

Serve locally: `task serve` (requires `npx serve`)

## Test

```
go test ./...
```

## File Layout

| File | Purpose |
|------|---------|
| `main.go` | Entrypoint, uses `booba.Run` for native + WASM compatibility |
| `model.go` | BubbleTea model: state, Update (key handling), View (rendering) |
| `calc.go` | Pure functions: `MinAge()`, `MaxAge()`, `InRange()` |
| `calc_test.go` | Table-driven tests for calc functions |
| `chart.go` | NTCharts line chart builder for the acceptable-range visualization |
| `web/index.html` | Static page for GitHub Pages (customize freely) |
| `.github/workflows/pages.yml` | Build WASM + deploy to GitHub Pages |

## Test

```
go test ./...
```

## File Layout

| File | Purpose |
|------|---------|
| `main.go` | Entrypoint, initializes BubbleTea program |
| `model.go` | BubbleTea model: state, Update (key handling), View (rendering) |
| `calc.go` | Pure functions: `MinAge()`, `MaxAge()`, `InRange()` |
| `calc_test.go` | Table-driven tests for calc functions |
| `chart.go` | NTCharts line chart builder for the acceptable-range visualization |

## Conventions

- All code is in `package main` (single binary)
- v2 import paths: `charm.land/bubbletea/v2`, `charm.land/lipgloss/v2`, `github.com/NimbleMarkets/ntcharts/v2`
- Calculation logic is separated from TUI logic for testability
- The chart redraws on every `View()` call based on current model state
