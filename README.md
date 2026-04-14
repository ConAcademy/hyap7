# hyap7 — Half Your Age Plus 7

<p>
    <a href="https://github.com/ConAcademy/hyap7/tags"><img src="https://img.shields.io/github/tag/ConAcademy/hyap7.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/ConAcademy/hyap7?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
</p>

`hyap7` is a terminal UI for exploring the ["Half Your Age Plus 7"](https://en.wikipedia.org/wiki/Age_disparity_in_sexual_relationships#The_%22half-your-age-plus-seven%22_rule) dating age rule, allowing exploration of quantitative ethics in the terminal.

<p><img src="etc/demo.gif" width="600" alt="hyap7 demo"></p>

The folk rule says the youngest person you should date is **(your age / 2) + 7**.  The inverse says the oldest person who should date you is **(your age × 2 - 14)**.

`hyap7` lets you explore this interactively with a live chart of the acceptable dating zone.

## Features

 * Adjustable age input with real-time range calculation
 * Forward and inverse formula display
 * Two-person compatibility check with pass/fail verdict
 * Braille line chart of the acceptable dating range zone using [NTCharts](https://github.com/NimbleMarkets/ntcharts)
 * Vim-style keyboard controls

## Installation

 * Download from the [`hyap7` releases page](https://github.com/ConAcademy/hyap7/releases)
 * Or install with [Homebrew](https://brew.sh):
 ```
 brew install ConAcademy/homebrew-tap/hyap7`
 ```
 * Or [build from source](#building-from-source)

## Usage

```
hyap7 [my-age [other-age]]
```

| Key | Action |
|-----|--------|
| `←` `→` / `h` `l` | Adjust age by 1 |
| `Shift+←` `Shift+→` / `H` `L` | Adjust age by 5 |
| `Tab` | Switch between your age and partner age |
| `Backspace` | Clear partner age |
| `q` / `Ctrl+C` | Quit |

## Building from Source

Requires [Go](https://golang.org/) 1.26+.

```sh
go build -o hyap7 .
```

Or with [Task](https://taskfile.dev):

```sh
task build
```

### Available Tasks

| Task | Description |
|------|-------------|
| `task build` | Build the binary (runs tests first) |
| `task test` | Run tests |
| `task run` | Run directly via `go run` |
| `task tidy` | Run `go mod tidy` |
| `task clean` | Remove build artifacts |

## The Math

The **half-your-age-plus-7** rule defines an acceptable dating age range:

 * **Minimum partner age:** `age / 2 + 7`
 * **Maximum partner age:** `age × 2 - 14` (the inverse)

This creates an asymmetric range that widens with age.  At 20 your range is 17–26; at 50 it's 32–86.

## Dependencies

 * [BubbleTea v2](https://charm.land/bubbletea) — TUI framework
 * [Lip Gloss v2](https://charm.land/lipgloss) — Terminal styling
 * [NTCharts v2](https://github.com/NimbleMarkets/ntcharts) — Terminal charts

## Open Collaboration

We welcome contributions and feedback.

 * [GitHub Issues](https://github.com/ConAcademy/hyap7/issues)
 * [GitHub Pull Requests](https://github.com/ConAcademy/hyap7/pulls)

## License

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE.txt](./LICENSE.txt).

Copyright (c) 2026 [Neomantra BV](https://www.neomantra.com).

----
Made with :heart: and :fire: by the team at [ConAcademy](https://github.com/ConAcademy).
