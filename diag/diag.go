package diag

import (
	"bufio"
	"fmt"
	"io"
)

type Tracer struct {
	traceTo *bufio.Writer
	format  string
}

func (target *Tracer) TraceLine(line string) {
	_, _ = target.traceTo.WriteString(line + "\n")
}
func (target *Tracer) TraceData(data ...any) {
	_, _ = target.traceTo.WriteString(fmt.Sprintf(target.format, data...) + "\n")
}
func (target *Tracer) Format(format string) {
	target.format = format
}

func NewTracer(target io.Writer) *Tracer {
	return &Tracer{traceTo: bufio.NewWriter(target), format: ""}
}
