# claude-code-fpf

A single [Claude Code](https://docs.anthropic.com/en/docs/claude-code) skill for querying [FPF (First Principles Framework)](https://github.com/ailev/FPF) wisdom on demand.

FPF is created by [Anatoly Levenchuk](https://github.com/ailev) — a comprehensive systems thinking and engineering methodology. This project packages the FPF specification into a searchable embedded index, delivered as a single Claude Code skill with zero ceremony.

## How it works

One skill. One binary. No hooks, no mandatory chains, no context bloat.

- **`fpf.md`** — a Claude Code skill with an inline index of FPF concepts, terms, and sections. The model knows *when* to look things up without calling the binary every turn.
- **`fpf-rag`** — a static Go binary with the full FPF spec indexed via SQLite FTS5, embedded into the binary itself. Query it for detailed guidance on any FPF concept.

The agent calls `fpf-rag search "<query>"` only when it needs deep FPF context — specific definitions, procedures, principles, or templates.

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/m0n0x41d/claude-code-fpf/main/install.sh | bash
```

This places:
- `~/.claude/skills/fpf/fpf.md` — the skill file
- `~/.claude/skills/fpf/references/fpf-rag` — the search binary

## Usage

Once installed, the `/fpf` skill is available in any Claude Code session. The agent will:
1. Consult the inline index to decide if a deeper search is needed
2. Run `fpf-rag search "<query>"` for detailed FPF guidance
3. Apply the retrieved principles to the current task

You can also use the binary directly:

```bash
# Search for a concept
fpf-rag search "WLNK weak link"

# Search with context limit
fpf-rag search "ADI cycle" --limit 5

# Show version and upstream FPF commit
fpf-rag info
```

## Versioning

Releases follow `v0.YYYY.MMDD` format, tracking the upstream FPF commit SHA. Each release references the exact FPF-Spec.md version it was built from.

| This project | Upstream FPF |
|-------------|-------------|
| Release tag | Commit SHA in `FPF_VERSION` |
| [Releases](https://github.com/m0n0x41d/claude-code-fpf/releases) | [ailev/FPF](https://github.com/ailev/FPF) |

## Development

Requires: [Go 1.22+](https://go.dev/), [Task](https://taskfile.dev/)

```bash
# Pull latest FPF spec (skips if unchanged)
task sync

# Build the search index
task build-db

# Build binary for current platform
task build

# Cross-compile for all platforms
task build-all

# Full release flow: sync → build → tag → push
task release
```

## Acknowledgements

This project exists because of [Anatoly Levenchuk's](https://ailev.livejournal.com/) work on the [First Principles Framework](https://github.com/ailev/FPF) — a rigorous systems engineering methodology that brings principled reasoning to AI-assisted work.

FPF-Spec.md is the normative source. This project is a delivery mechanism, not a fork.

## License

MIT
