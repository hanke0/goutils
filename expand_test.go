package goutils

import (
	"testing"
)

func TestExpand(t *testing.T) {
	cases := []struct {
		name      string
		tpl       string
		mapping   map[string]string
		want      string
		wantError bool
	}{
		{
			name:      "empty",
			tpl:       "",
			mapping:   nil,
			want:      "",
			wantError: false,
		},
		{
			name:      "no var",
			tpl:       "我是007",
			mapping:   nil,
			want:      "我是007",
			wantError: false,
		},
		{
			name:      "var default ok",
			tpl:       "${who:=谁}是${name}!",
			mapping:   map[string]string{"name": "007"},
			want:      "谁是007!",
			wantError: false,
		},
		{
			name:      "var default chinese ok",
			tpl:       "${变量:=我是}007!",
			mapping:   nil,
			want:      "我是007!",
			wantError: false,
		},
		{
			name:      "var is empty",
			tpl:       "我是${}!",
			mapping:   nil,
			want:      "我是${}!",
			wantError: true,
		},
		{
			name:      "var is incomplete",
			tpl:       "${var007!",
			mapping:   nil,
			want:      "${var007!",
			wantError: true,
		},
		{
			name:      "var is incomplete",
			tpl:       "我是${007!",
			mapping:   nil,
			want:      "我是${007!",
			wantError: true,
		},
		{
			name:      "var is not start with brace",
			tpl:       "我是$007!",
			mapping:   nil,
			want:      "我是$007!",
			wantError: false,
		},
		{
			name:      "good var after bad var",
			tpl:       "我是${0${name:=07}!",
			mapping:   nil,
			want:      "我是${007!",
			wantError: true,
		},
		{
			name:      "var with colon",
			tpl:       "${name-::=我是}007!",
			mapping:   nil,
			want:      "我是007!",
			wantError: false,
		},
		{
			name:      "var with defult :",
			tpl:       "${a我:=:=我是}007!",
			mapping:   nil,
			want:      ":=我是007!",
			wantError: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := func(a string, def string) string {
				if c.mapping == nil {
					return def
				}
				r, ok := c.mapping[a]
				if ok {
					return r
				}
				return def
			}
			got, err := Expand(c.tpl, f)
			if (err != nil) != c.wantError {
				t.Errorf("wantError:%v,got:%v", c.wantError, err)
			}
			if got != c.want {
				t.Errorf("\nwant:%s\ngot:%s", c.want, got)
			}
		})
	}
}
