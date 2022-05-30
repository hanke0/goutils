package goutils

import (
	"bytes"
	"encoding/json"
	"sync"
)

var jsonEncoderPool sync.Pool

type jsonencoder struct {
	bs bytes.Buffer
	j  json.Encoder
}

func (j *jsonencoder) setpretty() {
	j.j.SetIndent("", "  ")
}

func (j *jsonencoder) reset() {
	j.bs.Reset()
	j.j.SetIndent("", "")
}

func getjsonencoder() *jsonencoder {
	e, ok := jsonEncoderPool.Get().(*jsonencoder)
	if !ok || e == nil {
		e = new(jsonencoder)
		e.j = *json.NewEncoder(&e.bs)
		e.j.SetEscapeHTML(false)
		return e
	}
	return e
}

func putjsonencoder(e *jsonencoder) {
	e.reset()
	jsonEncoderPool.Put(e)
}

// ToJSON return the object json marshal not HTML escaped string.
// An empty string returned if there is a marshal error.
func ToJSON(o interface{}) string {
	e := getjsonencoder()
	defer putjsonencoder(e)
	err := e.j.Encode(o)
	if err != nil {
		panic(err)
	}
	if e.bs.Len() > 1 {
		e.bs.Truncate(e.bs.Len() - 1)
	}
	return e.bs.String()
}

// ToPrettyJson like ToJSON with pretty output
func ToPrettyJson(o interface{}) string { // nolint: revive
	e := getjsonencoder()
	defer putjsonencoder(e)
	e.setpretty()
	err := e.j.Encode(o)
	if err != nil {
		panic(err)
	}
	if e.bs.Len() > 1 {
		e.bs.Truncate(e.bs.Len() - 1)
	}
	return e.bs.String()
}

// ToPrettyJSON likes ToJSON with pretty output
var ToPrettyJSON = ToPrettyJson
