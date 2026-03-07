package core

import "testing"

func TestParseEnv(t *testing.T) {
	input := `
# comment
export ANTHROPIC_BASE_URL='https://example.com'
ANTHROPIC_MODEL="kimi-k2.5"
INVALID_LINE
`
	m := ParseEnv(input)
	if m["ANTHROPIC_BASE_URL"] != "https://example.com" {
		t.Fatalf("unexpected base url: %q", m["ANTHROPIC_BASE_URL"])
	}
	if m["ANTHROPIC_MODEL"] != "kimi-k2.5" {
		t.Fatalf("unexpected model: %q", m["ANTHROPIC_MODEL"])
	}
}

func TestValidProfileName(t *testing.T) {
	if !ValidProfileName("bailian-kimi-k2_5") {
		t.Fatal("expected valid profile name")
	}
	if ValidProfileName("bad name") {
		t.Fatal("expected invalid profile name")
	}
}
