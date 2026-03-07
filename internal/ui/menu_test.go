package ui

import "testing"

func TestParseMenuSelection(t *testing.T) {
	cases := []struct {
		in      string
		max     int
		want    int
		wantErr bool
	}{
		{"", 3, 0, false},
		{"   ", 3, 0, false},
		{"0", 3, 0, false},
		{"2", 3, 2, false},
		{"x", 3, 0, true},
		{"4", 3, 0, true},
	}

	for _, c := range cases {
		got, err := ParseMenuSelection(c.in, c.max)
		if (err != nil) != c.wantErr {
			t.Fatalf("in=%q err=%v wantErr=%v", c.in, err, c.wantErr)
		}
		if got != c.want {
			t.Fatalf("in=%q got=%d want=%d", c.in, got, c.want)
		}
	}
}
