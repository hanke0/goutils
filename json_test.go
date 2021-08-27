package goutils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestToJSON(t *testing.T) {
	type sj struct {
		I int    `json:"i"`
		S string `json:"s"`
	}
	for i := 0; i < 1000; i++ {
		ic := i
		t.Run(fmt.Sprintf("TestToJSON-%d", ic), func(t *testing.T) {
			t.Parallel()
			s := sj{ic, fmt.Sprintf("我&<%d>你", ic)}
			got := ToJSON(s)
			want := fmt.Sprintf(`{"i":%[1]d,"s":"我&<%[1]d>你"}`, ic)
			if got != want {
				t.Errorf("\nwant:`%s`\ngot :`%s`", want, got)
			}
		})
	}
}

func TestToPrettyJSON(t *testing.T) {
	type sj struct {
		I int    `json:"i"`
		S string `json:"s"`
	}
	for i := 0; i < 1000; i++ {
		ic := i
		t.Run(fmt.Sprintf("TestToJSON-%d", ic), func(t *testing.T) {
			t.Parallel()
			s := sj{ic, fmt.Sprintf("我&<%d>你", ic)}
			got := ToPrettyJson(s)
			want := fmt.Sprintf("{\n  \"i\": %[1]d,\n  \"s\": \"我&<%[1]d>你\"\n}", ic)
			if got != want {
				t.Errorf("\nwant:`%s`\ngot :`%s`", want, got)
			}
		})
	}
}

func BenchmarkToJSON(b *testing.B) {
	type sj struct {
		I int    `json:"i"`
		S string `json:"s"`
	}
	for i := 0; i < b.N; i++ {
		s := sj{i, fmt.Sprintf("我&<%d>你", i)}
		_ = ToJSON(s)
	}
}

func BenchmarkRawJSONMarshal(b *testing.B) {
	type sj struct {
		I int    `json:"i"`
		S string `json:"s"`
	}
	for i := 0; i < b.N; i++ {
		s := sj{i, fmt.Sprintf("我&<%d>你", i)}
		e, _ := json.Marshal(s)
		a := string(e)
		_ = a
	}
}
