package wlog

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// MockUI is a UI that is used for tests.
// It is exported publicly for use in external tests if needed as well.
type MockUI struct {
	Reader      io.Reader
	Writer      *bytes.Buffer
	ErrorWriter *bytes.Buffer

	once sync.Once
}

// Error writes message to ErrorWriter.
func (u *MockUI) Error(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.ErrorWriter, message)
	fmt.Fprint(u.ErrorWriter, "\n")
}

// Info calls Output to write.
// Useful when you want seperate colors or prefixes.
func (u *MockUI) Info(message string) {
	u.Output(message)
}

// Output simply writes to Writer.
func (u *MockUI) Output(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

// Running calls Output to write.
// Useful when you want seperate colors or prefixes.
func (u *MockUI) Running(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

// Success calls Output to write.
// Useful when you want seperate colors or prefixes.
func (u *MockUI) Success(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

// Warn calls Error to write.
// Useful when you want seperate colors or prefixes.
func (u *MockUI) Warn(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.ErrorWriter, message)
	fmt.Fprint(u.ErrorWriter, "\n")
}

// Log prefixes to message before writing to Writer.
func (u *MockUI) Log(message string) {
	u.once.Do(u.init)
	timeString := time.Now().Format(timeFormat)
	message = timeString + ": " + message
	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

func (u *MockUI) init() {
	u.ErrorWriter = new(bytes.Buffer)
	u.Writer = new(bytes.Buffer)
}
