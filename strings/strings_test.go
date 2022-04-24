package strings

import (
	"testing"
)

func TestListContains(t *testing.T) {
	var cases = []struct {
		list string
		sub  string
		want bool
	}{
		{"124,456", "124", true},
		{"124,456", "456", true},
		{"124,456", "121", false},
		{"124,456", "12", false},
		{"124,456", "56", false},
		{"124,456", "45", false},
		{"124,456", "24", false},
		{"124", "124", true},
		{"124", "12", false},
		{"124", "24", false},
		{"124", "456", false},
		{",124,", "124", true},
		{",124,", "12", false},
		{",124,", "24", false},
		{",124,", "134", false},
		{"", "124", false},
	}

	for _, c := range cases {
		if got := ListContains(c.list, c.sub, ','); got != c.want {
			t.Errorf("list=%s, sub=%s, want=%v, got=%v", c.list, c.sub, c.want, got)
		}
	}
}
