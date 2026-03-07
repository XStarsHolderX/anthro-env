# anthro-env

[![Release](https://img.shields.io/github/v/release/kelaocai/anthro-env)](https://github.com/kelaocai/anthro-env/releases)
[![Homebrew Tap](https://img.shields.io/badge/Homebrew-kelaocai%2Fhomebrew--tap-blue)](https://github.com/kelaocai/homebrew-tap)

`anthro-env` is a **macOS CLI for Claude Code / Anthropic environment profile switching**.
It helps developers switch between Anthropic-compatible gateways and models with one command, while keeping tokens in macOS Keychain.

Common search intents this project solves:
- `claude code environment variable manager`
- `anthropic profile switcher mac`
- `claude code homebrew tool`
- `switch anthropic base url and token quickly`

## Quick Links

- 中文 README: [README.zh-CN.md](README.zh-CN.md)
- Quick Start (中文): [docs/快速上手.md](docs/快速上手.md)
- Full Guide (中文): [docs/项目中文手册.md](docs/项目中文手册.md)
- FAQ (EN): [docs/faq.md](docs/faq.md)
- Comparison (EN): [docs/anthro-env-vs-manual.md](docs/anthro-env-vs-manual.md)

## Why This Project Exists

Claude Code is excellent for daily coding, but many developers face practical issues:
- official token cost
- region availability limits
- account/policy uncertainty

So people often use Anthropic-compatible third-party gateways.
The painful part is frequent environment variable switching.

`anthro-env` exists to make that workflow simple, safe, and fast.

## Install (Homebrew)

```bash
brew tap kelaocai/homebrew-tap
brew install anthro-env
```

No local Go toolchain is required for users during Homebrew install.

## Build from Source (for power users)

If you prefer local compilation or want to contribute:

```bash
git clone https://github.com/kelaocai/anthro-env.git
cd anthro-env
go test ./...
go build -o ./bin/anthro-env ./cmd/anthro-env
./bin/anthro-env --help
```

## 30-Second Start

```bash
anthro-env init
source ~/.zshrc
anthro-env menu
```

## Core Commands

```bash
anthro-env init
anthro-env menu
anthro-env add <name>
anthro-env use <name>
anthro-env ls
anthro-env current
anthro-env rm <name>
anthro-env doctor
```

## Key Features

- One-command profile switching for Claude Code / Anthropic env.
- Interactive menu (`Enter` defaults to Exit).
- Token in macOS Keychain (not plain text profile file).
- Shell hook support for zsh/bash and auto-sync in current shell.
- Built-in doctor command for fast troubleshooting.

## Storage Layout

- Profiles: `~/.config/anthropic/profiles/*.env`
- Current profile: `~/.config/anthropic/current`
- Token storage: macOS Keychain (`service=anthro-env`)

## Who It Is For

- Developers using Claude Code with multiple providers/gateways
- Teams needing stable env profile switching on macOS
- Users who want Homebrew installation and low-friction setup

## Contributing

Issues and PRs are welcome.

## License

MIT

## Maintainer Notes

- On tag push (`v*`), GitHub Actions updates `Formula/anthro-env.rb` in `kelaocai/homebrew-tap`.
- Required repository secret: `HOMEBREW_TAP_TOKEN` (PAT with repo write access to `kelaocai/homebrew-tap`).
- If sync fails, workflow prints explicit reasons (missing secret, no tap access, push denied).
