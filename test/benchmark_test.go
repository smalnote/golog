package test

import (
	"bytes"
	"net/http"
	"testing"
	"text/template"
)

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

func BenchmarkGetNodeURLParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Get("http://127.0.0.1:8080/nodeURL")
			if err != nil {
				continue
			}
			defer resp.Body.Close()
		}
	})
}
