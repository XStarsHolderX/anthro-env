class AnthroEnv < Formula
  desc "macOS Anthropic environment profile manager"
  homepage "https://github.com/kelaocai/anthro-env"
  url "https://github.com/kelaocai/anthro-env/archive/refs/tags/v0.1.0-alpha.tar.gz"
  sha256 "336026c1009b5b18a830e17aa15f4806fb9f7f4e04a852e400cbda0ef6337ed3"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/anthro-env"
  end

  test do
    assert_match "anthro-env commands", shell_output("#{bin}/anthro-env --help")
  end
end
