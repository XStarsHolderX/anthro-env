package ui

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMenuSelection(input string, max int) (int, error) {
	in := strings.TrimSpace(input)
	if in == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(in)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}
	if n < 0 || n > max {
		return 0, fmt.Errorf("out of range")
	}
	return n, nil
}
