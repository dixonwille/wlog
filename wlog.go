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
func NewColor(successColor, warnColor, infoColor, errorColor ct.Color, ui UI) *ColorUI {
	return &ColorUI{
		LogFGColor:     ct.None,
		LogBGColor:     ct.None,
		OutputFGColor:  ct.None,
		OutputBGColor:  ct.None,
		SuccessFGColor: successColor,
		SuccessBGColor: ct.None,
		InfoFGColor:    infoColor,
		InfoBGColor:    ct.None,
		ErrorFGColor:   errorColor,
		ErrorBGColor:   ct.None,
		WarnFGColor:    warnColor,
		WarnBGColor:    ct.None,
		UI:             ui,
	}
}

// New creates a UI that is concurrent and colorfull
func New(reader io.Reader, writer, errorWriter io.Writer, successColor, warnColor, infoColor, errorColor ct.Color) UI {
	basic := NewBasic(reader, writer, errorWriter)
	color := NewColor(successColor, warnColor, infoColor, errorColor, basic)
	concurrent := NewConcurrent(color)
	return concurrent
}
