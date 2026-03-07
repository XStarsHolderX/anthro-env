# Release / Maintainer Notes

This doc is for maintainers only.

## Release flow

1. Commit changes to `main`
2. Create and push tag (example `v0.1.3-alpha`)
3. `release` workflow builds macOS binaries and uploads release assets
4. `update-homebrew-tap` updates `kelaocai/homebrew-tap` Formula

## Required secret

Repository `kelaocai/anthro-env` must have:

- `HOMEBREW_TAP_TOKEN`

The token needs write access to `kelaocai/homebrew-tap`.

## Verification checklist

- Release assets include:
  - `anthro-env_<version>_macos_arm64.tar.gz`
  - `anthro-env_<version>_macos_x86_64.tar.gz`
  - `checksums.txt`
- Tap formula is updated to latest tag and sha256 values
- `brew install anthro-env` works on a clean machine
