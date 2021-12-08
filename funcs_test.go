package goutils

import "testing"

func TestShortcutUTF8(t *testing.T) {
	cases := []struct {
		s      string
		max    int
		suffix string
		want   string
	}{
		{
			s:      "foobar",
			max:    3,
			suffix: "...",
			want:   "foo...",
		},
		{
			s:      "foobar",
			max:    6,
			suffix: "...",
			want:   "foobar",
		},
		{
			s:      "中文",
			max:    2,
			suffix: "...",
			want:   "中文",
		},
		{
			s:      "中文",
			max:    1,
			suffix: "...",
			want:   "中...",
		},
		{
			s:      "foobar",
			max:    0,
			suffix: "...",
			want:   "...",
		},
		{
			s:      "foobar",
			max:    0,
			suffix: "...",
			want:   "...",
		},
	}
	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			got := ShortcutUTF8(c.s, c.max, c.suffix)
			if got != c.want {
				t.Errorf("\nwant:%s\ngot:%s\n", c.want, got)
			}
		})
	}
}
