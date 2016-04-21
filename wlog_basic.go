package wlog

import (
	"fmt"
	"io"
	"time"
)

const (
	timeFormat = "2006-01-02T15-04-05"
)

// BasicUI simply writes/reads to correct input/output
// It is not thread safe.
// Pretty simple to wrap your own functions around
type BasicUI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

// Log prefixes to message before writing to Writer.
func (ui *BasicUI) Log(message string) {
	timeString := time.Now().Format(timeFormat)
	message = timeString + ": " + message
	ui.Output(message)
}

// Output simply writes to Writer.
func (ui *BasicUI) Output(message string) {
	fmt.Fprint(ui.Writer, message)
	fmt.Fprint(ui.Writer, "\n")
}

// Success calls Output to write.
// Useful when you want seperate colors or prefixes.
func (ui *BasicUI) Success(message string) {
	ui.Output(message)
}

// Info calls Output to write.
// Useful when you want seperate colors or prefixes.
func (ui *BasicUI) Info(message string) {
	ui.Output(message)
}

// Error writes message to ErrorWriter.
func (ui *BasicUI) Error(message string) {
	if ui.ErrorWriter != nil {
		fmt.Fprint(ui.ErrorWriter, message)
		fmt.Fprint(ui.ErrorWriter, "\n")
	} else {
		fmt.Fprint(ui.Writer, message)
		fmt.Fprint(ui.Writer, "\n")
	}
}

// Warn calls Error to write.
// Useful when you want seperate colors or prefixes.
func (ui *BasicUI) Warn(message string) {
	ui.Error(message)
}

// Running calls Output to write.
// Useful when you want seperate colors or prefixes.
func (ui *BasicUI) Running(message string) {
	ui.Output(message)
}
