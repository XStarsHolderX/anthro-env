# anthro-env

[![Release](https://img.shields.io/github/v/release/kelaocai/anthro-env)](https://github.com/kelaocai/anthro-env/releases)
[![Homebrew Tap](https://img.shields.io/badge/Homebrew-kelaocai%2Fhomebrew--tap-blue)](https://github.com/kelaocai/homebrew-tap)

`anthro-env` is a macOS CLI for Claude Code / Anthropic environment profile switching.
It lets you switch between Anthropic-compatible gateways and models with one command, while storing tokens in macOS Keychain.

## Why This Tool Exists

Claude Code is great for daily coding, but many developers still face practical issues:
- official token cost
- region availability limits
- account/policy uncertainty

So people often use Anthropic-compatible third-party gateways.
The painful part is frequent environment variable switching.

`anthro-env` makes that workflow simple, safe, and fast.

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

## Usage

### Quick start

```bash
anthro-env init
source ~/.zshrc
anthro-env menu
```

### Core commands

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
