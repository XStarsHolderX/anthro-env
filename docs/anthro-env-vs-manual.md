# anthro-env vs Manual Environment Variable Switching

## Manual switching

- Repeated edits in shell files
- Easy to mix old/new values
- Hard to keep token handling safe
- Troubleshooting is scattered

## anthro-env

- Profile-based switching with one command
- Stable state tracking (`current` profile)
- Token kept in macOS Keychain
- Built-in `doctor` for quick diagnostics
- Homebrew install for fast onboarding

## Bottom line

If you only have one fixed provider, manual may be enough.
If you switch providers/models regularly, `anthro-env` is faster and safer.
