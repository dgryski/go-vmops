package vmops

import (
	"testing"
)

// ab*c
var abcVM = VM{
	{OpEOF, 0, ActionRET, -1},
	{OpNE, 'a', ActionGOTO, 0},
	{OpEOF, 0, ActionRET, -1},
	{OpLE, '`', ActionGOTO, 0},
	{OpLE, 'b', ActionGOTO, 2},
	{OpNE, 'c', ActionGOTO, 0},
	{OpALWAYS, '\x00', ActionRET, 2},
}

func TestMatchABC(t *testing.T) {
	tests := []struct {
		s     string
		match bool
	}{
		{"a", false},
		{"aa", false},
		{"ab", false},
		{"ac", true},
		{"abc", true},
		{"abbbc", true},
		{"Qabc", true},
	}

	for _, tt := range tests {
		got := abcVM.Match([]byte(tt.s)) != -1
		if got != tt.match {
			t.Errorf("abcVM.Match(%q)=%v, want %v\n", tt.s, got, tt.match)

		}

		// -1 on failure
		got = abcVM.MatchString(tt.s) != -1
		if got != tt.match {
			t.Errorf("abcVM.MatchString(%q)=%v, want %v\n", tt.s, got, tt.match)

		}
	}
}
