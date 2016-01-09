package typer

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Typer struct {
	Index  int
	Speed  int
	Buffer []byte
	Out    chan []byte
}

func (t *Typer) LoadSource() {
	res, _ := http.Get(sourceUri())
	defer res.Body.Close()
	if body, err := ioutil.ReadAll(res.Body); err != nil {
		t.Buffer = []byte("Houston we have a problem...")
	} else {
		t.Buffer = body
	}
	t.Out <- []byte("goTyper:Ready")
}

func (t *Typer) Type() {
	if t.Index < len(t.Buffer) {
		res := format(t.Buffer[t.Index : t.Index+t.Speed])
		t.Out <- []byte(res)
		t.Index += t.Speed
	}
}

func format(buffer []byte) string {
	var out bytes.Buffer
	for _, n := range buffer {
		switch n {
		case 34:
			out.WriteString("&quot;")
		case 39:
			out.WriteString("&#039;")
		case 38:
			out.WriteString("&amp;")
		case 60:
			out.WriteString("&lt;")
		case 62:
			out.WriteString("&gt;")
		default:
			out.WriteString(string(n))
		}
	}
	return out.String()
}

func sourceUri() string {
	paths := []string{
		"compress/bzip2/bzip2_test",
		"compress/gzip/gzip_test",
		"encoding/csv/writer",
		"encoding/json/stream",
		"fmt/format",
		"runtime/atomic_arm",
		"runtime/gengoos",
		"reflect/all_test",
	}
	return "https://golang.org/src/" + paths[rand.Intn(len(paths))] + ".go?m=text"
}
