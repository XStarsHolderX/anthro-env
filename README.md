# anthro-env

[![Release](https://img.shields.io/github/v/release/kelaocai/anthro-env)](https://github.com/kelaocai/anthro-env/releases)
[![Homebrew Tap](https://img.shields.io/badge/Homebrew-kelaocai%2Fhomebrew--tap-blue)](https://github.com/kelaocai/homebrew-tap)

- 快速上手（中文）: [docs/快速上手.md](docs/快速上手.md)
- 完整中文手册: [docs/项目中文手册.md](docs/项目中文手册.md)
- English README: this file

A macOS-first profile manager for Claude Code / Anthropic environment variables.

## Why This Project Exists

Claude Code is one of the best coding CLI assistants today, and many people use it as their daily coding partner.

But in real life, using official Anthropic access is often painful:

- official tokens can be expensive
- some regions are not supported
- account risk and policy limits can suddenly block usage

So many developers use Anthropic-compatible third-party gateways and models for better availability and lower cost.
The problem is: switching environment variables all day is messy and annoying.

`anthro-env` was built to solve exactly that everyday pain, with a simple profile switch flow.

## Features

- Switch Anthropic profiles with one command.
- Interactive menu with safe defaults (`Enter` = Exit).
- Token stored in macOS Keychain (not in profile files).
- Auto-sync env vars in your current shell via hook.

## Quickstart

```bash
brew install anthro-env/tap/anthro-env
anthro-env init
source ~/.zshrc
anthro-env menu
```

## Commands

```bash
anthro-env init
anthro-env menu
anthro-env profile add <name>
anthro-env profile use <name>
anthro-env profile ls
anthro-env profile current
anthro-env profile rm <name>
anthro-env doctor
```

## Storage Layout

- Profiles: `~/.config/anthropic/profiles/*.env`
- Current profile: `~/.config/anthropic/current`
- Token: macOS Keychain (`service=anthro-env`)

## License

MIT

## Maintainer Notes

- On tag push (`v*`), GitHub Actions updates `Formula/anthro-env.rb` in `kelaocai/homebrew-tap`.
- Required repository secret: `HOMEBREW_TAP_TOKEN` (a PAT with repo write access to `kelaocai/homebrew-tap`).
