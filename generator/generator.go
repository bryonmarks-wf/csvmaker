package generator
import (
	"fmt"
	"bytes"
)

type Generator struct {
	numRows int
	numColumns int
	buffer bytes.Buffer
}

// helper method to write a specific line to the buffer
func (cr *Generator) writeLine(lineIdx int, numCols int) {
	for i := 0; i < numCols; i++ {
		cr.buffer.WriteString(fmt.Sprintf("%d", lineIdx))
		if (i != numCols - 1) {
			cr.buffer.WriteString(", ")
		}
	}
}

// writes all lines to the buffer
func (cr *Generator) writeLines(numLines int, numCols int) {
	for i := 0; i < numLines; i++ {
		cr.writeLine(i, numCols)
		cr.buffer.WriteString("\n")
	}
}

// Returns a new Generator
func New(numRows int, numColumns int) *Generator {
	var buffer bytes.Buffer
	cr := Generator{numRows, numColumns, buffer}
	cr.writeLines(numRows, numColumns)
	return &cr
}

// Read from the Generator
func (r *Generator) Read(p []byte) (n int, err error) {
	return r.buffer.Read(p)
}

