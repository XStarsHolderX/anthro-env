# anthro-env FAQ

## What is anthro-env?

`anthro-env` is a macOS CLI tool for switching Claude Code / Anthropic environment profiles.

## Why not switch env vars manually?

Manual switching is error-prone and slow when you use multiple gateways or models.
`anthro-env` gives you profile-based switching with one command.

## Does anthro-env store my API token in plain text?

By default, no. Tokens are stored in macOS Keychain.

## Is this only for Anthropic official API?

No. It works with Anthropic-compatible gateways and model providers.

## Does it support Homebrew installation?

Yes.

```bash
brew tap kelaocai/homebrew-tap
brew install anthro-env
```

## Which shells are supported?

Currently zsh-first, with bash hook support.

## How do I check if setup is healthy?

Run:

```bash
anthro-env doctor
```

## I switched profile but current terminal didn't update. What now?

```bash
source ~/.zshrc
anthro-env doctor
```

## Is there a GUI app?

Not yet in v1. CLI is the current focus. GUI is planned for a later stage.
