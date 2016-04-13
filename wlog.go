package wlog

import (
	"io"

	"github.com/daviddengcn/go-colortext"
)

// NewConcurrent returns a new ConcurrentUI
func NewConcurrent(ui UI) *ConcurrentUI {
	return &ConcurrentUI{UI: ui}
}

//NewBasic returns a BasicUI for use
func NewBasic(reader io.Reader, writer, errorWriter io.Writer) *BasicUI {
	return &BasicUI{
		Reader:      reader,
		Writer:      writer,
		ErrorWriter: errorWriter,
	}
}

//NewColor returns a ColorUI for use
func NewColor(outputColor ct.Color, errorColor ct.Color, ui UI) *ColorUI {
	return &ColorUI{
		LogFGColor:    outputColor,
		LogBGColor:    ct.None,
		OutputFGColor: outputColor,
		OutputBGColor: ct.None,
		InfoFGColor:   outputColor,
		InfoBGColor:   ct.None,
		ErrorFGColor:  errorColor,
		ErrorBGColor:  ct.None,
		FatalFGColor:  errorColor,
		FatalBGColor:  ct.None,
		WarnFGColor:   errorColor,
		WarnBGColor:   ct.None,
		UI:            ui,
	}
}

// New creates a UI that is concurrent and colorfull
func New(reader io.Reader, writer, errorWriter io.Writer, outputColor ct.Color, errorColor ct.Color) UI {
	basic := NewBasic(reader, writer, errorWriter)
	color := NewColor(outputColor, errorColor, basic)
	concurrent := NewConcurrent(color)
	return concurrent
}
