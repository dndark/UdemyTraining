package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of tracing evnts throughout code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}
type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprintln(t.out, a...)
}
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
