# cronscope

A terminal UI for visualizing, validating, and dry-running cron expressions with next-run previews.

---

## Installation

```bash
go install github.com/yourusername/cronscope@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/cronscope.git
cd cronscope
go build -o cronscope .
```

---

## Usage

Launch the interactive TUI:

```bash
cronscope
```

You can also pass a cron expression directly to get a quick preview:

```bash
cronscope "*/15 9-17 * * 1-5"
```

Inside the TUI, type or paste any cron expression into the input field. cronscope will:

- **Validate** the expression in real time and highlight any errors
- **Display** a human-readable description of the schedule
- **Preview** the next 10 upcoming run times

Use `Tab` to switch between fields, `Enter` to confirm, and `q` / `Ctrl+C` to quit.

---

## Supported Formats

- Standard 5-field cron (`* * * * *`)
- Extended 6-field with seconds (`* * * * * *`)
- Common macros (`@daily`, `@hourly`, `@weekly`, etc.)

---

## Requirements

- Go 1.21 or later

---

## License

MIT © 2024 yourusername