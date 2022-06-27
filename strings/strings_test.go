package strings_test

import (
	"testing"

	"github.com/ko-han/goutils/strings"
)

func TestContains(t *testing.T) {
	var list = []string{"1", "2", ""}
	if !strings.Contains(list, "1") {
		t.Fatal("1 should in the list")
	}
	if strings.Contains(list, "3") {
		t.Fatal("3 should not in the list")
	}
}

func TestIndex(t *testing.T) {
	var list = []string{"1", "2", ""}
	if i := strings.Index(list, "1"); i != 0 {
		t.Fatal("1 should in the list", i)
	}
	if i := strings.Index(list, "3"); i != -1 {
		t.Fatal("3 should not in the list", i)
	}
}

func TestCount(t *testing.T) {
	var list = []string{"1", "2", "3", "1"}
	if i := strings.Count(list, "1"); i != 2 {
		t.Fatal("1 should present 2 times", i)
	}
	if i := strings.Count(list, "3"); i != 1 {
		t.Fatal("3 should present 1 times", i)
	}
	if i := strings.Count(list, "4"); i != 0 {
		t.Fatal("4 should present 0 times", i)
	}
}

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
		{"", "", true},
		{"123", "", false},
	}

	for _, c := range cases {
		if got := strings.HasSeparatedSubstring(c.list, c.sub, ','); got != c.want {
			t.Errorf("list=%s, sub=%s, want=%v, got=%v", c.list, c.sub, c.want, got)
		}
	}
}
