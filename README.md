# anthro-env

English | [中文](README.zh.md)

[![Release](https://img.shields.io/github/v/release/kelaocai/anthro-env)](https://github.com/kelaocai/anthro-env/releases)
[![Homebrew Tap](https://img.shields.io/badge/Homebrew-kelaocai%2Fhomebrew--tap-blue)](https://github.com/kelaocai/homebrew-tap)

A CLI-first macOS tool for switching Claude Code / Anthropic environment profiles.

## Why this exists

Claude Code is a great CLI coding assistant.

But in practice many developers need to switch between:

- Anthropic API
- proxy gateways
- third-party providers
- multiple API keys

This usually means repeatedly editing environment variables like:

- `ANTHROPIC_BASE_URL`
- `ANTHROPIC_AUTH_TOKEN`

Manual shell edits quickly become messy.

`anthro-env` provides a simple profile-based workflow for managing these environments.

## Install

### Homebrew (recommended)

```bash
brew tap kelaocai/homebrew-tap
brew install anthro-env
```

No local Go toolchain is required for Homebrew users.

### Build from source (optional)

```bash
git clone https://github.com/kelaocai/anthro-env.git
cd anthro-env
go test ./...
go build -o ./bin/anthro-env ./cmd/anthro-env
./bin/anthro-env --help
```

## Quick Start

```bash
anthro-env add kimi
anthro-env use kimi
anthro-env current
```

Or use the menu:

```bash
anthro-env menu
```

## Concepts

- `profile`: a named environment config (base URL, model config, etc.)
- `use`: switch active profile
- `current`: show active profile
- `doctor`: quick health check

What makes this project different:

- Keychain-backed token handling
- CLI-first workflow
- Homebrew distribution
- low-friction setup for macOS users

## Storage Layout

- Profiles: `~/.config/anthropic/profiles/*.env`
- Active profile pointer: `~/.config/anthropic/current`
- Token storage: macOS Keychain (`service=anthro-env`)

## Security

- Tokens are stored in macOS Keychain.
- Profile files store metadata/config only.
- API keys are not intended to be stored in plain text profile files.

## Who It Is For

- Developers switching between multiple Claude Code / Anthropic-compatible providers
- Teams that want predictable profile switching on macOS
- Users who want Homebrew install and fast onboarding

## Roadmap

- Linux keyring support
- Windows credential manager support
- Profile export/import
- Team shared profile workflows

## Contributing

Issues and PRs are welcome.

## License

MIT
