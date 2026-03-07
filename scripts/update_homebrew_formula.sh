#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 3 ]]; then
  echo "Usage: $0 <version> <owner> <repo>"
  exit 1
fi

VERSION="$1"
OWNER="$2"
REPO="$3"

if [[ "$VERSION" != v* ]]; then
  echo "Version must start with 'v' (example: v0.1.0)"
  exit 1
fi

URL="https://github.com/${OWNER}/${REPO}/archive/refs/tags/${VERSION}.tar.gz"
SHA256="$(curl -L --max-time 60 -s "$URL" | shasum -a 256 | awk '{print $1}')"

cat > packaging/homebrew/anthro-env.rb <<FORMULA
class AnthroEnv < Formula
  desc "macOS Anthropic environment profile manager"
  homepage "https://github.com/${OWNER}/${REPO}"
  url "${URL}"
  sha256 "${SHA256}"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/anthro-env"
  end

  test do
    assert_match "anthro-env commands", shell_output("#{bin}/anthro-env --help")
  end
end
FORMULA

echo "Updated packaging/homebrew/anthro-env.rb for ${VERSION}"
