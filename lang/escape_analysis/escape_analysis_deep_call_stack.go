package escapeanalysis

import (
	"bytes"
	"io"
)

//go:noinline
func doRead(input *bytes.Buffer, buf []byte) {
	io.ReadFull(input, buf) // causing input, buf escape to heap
}

//go:noinline
func doRead2(input *bytes.Buffer, buf []byte) {
	// nothing, so no escape
}

func escapeAnalysisDeepCallStack() {
	b := make([]byte, 100)
	input := bytes.NewBuffer(b)
	buf := make([]byte, 10)
	doRead(input, buf)
}
