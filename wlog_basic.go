package wlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	timeFormat = "2006-01-02T15-04-05"
)

// BasicUI implements UI.
// It is not thread safe.
// Easy to wrap own functions around.
type BasicUI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

// Log writes to screen with date-time prefixed
func (ui *BasicUI) Log(message string) {
	timeString := time.Now().Format(timeFormat)
	message = timeString + ": " + message
	ui.Output(message)
}

// Output writes to Writer in BasicUI
func (ui *BasicUI) Output(message string) {
	fmt.Fprint(ui.Writer, message)
	fmt.Fprint(ui.Writer, "\n")
}

// Info writes to Writer in BasicUI
func (ui *BasicUI) Info(message string) {
	ui.Output(message)
}

// Error writes to Error in BasicUI
func (ui *BasicUI) Error(message string) {
	if ui.ErrorWriter != nil {
		fmt.Fprint(ui.ErrorWriter, message)
		fmt.Fprint(ui.ErrorWriter, "\n")
	} else {
		fmt.Fprint(ui.Writer, message)
		fmt.Fprint(ui.Writer, "\n")
	}
}

// Fatal writes to Error in BasicUI
// Also kills the process
func (ui *BasicUI) Fatal(message string) {
	ui.Error(message)
	os.Exit(1)
}

// Warn writes to Error in BasicUI
func (ui *BasicUI) Warn(message string) {
	ui.Error(message)
}
