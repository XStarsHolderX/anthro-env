#!/usr/bin/env bash
set -euo pipefail

VERSION="${1:-}"
if [[ -z "$VERSION" ]]; then
  echo "Usage: scripts/release.sh <version>"
  exit 1
fi

echo "Tagging version: $VERSION"
git tag "$VERSION"
git push origin "$VERSION"
