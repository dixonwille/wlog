package wlog

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// MockUI is a mock UI that is used for tests and is exported publicly for use in external tests if needed as well.
type MockUI struct {
	Reader      io.Reader
	Writer      *bytes.Buffer
	ErrorWriter *bytes.Buffer

	once sync.Once
}

//Error prints to bytes.Buffer for testing
func (u *MockUI) Error(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.ErrorWriter, message)
	fmt.Fprint(u.ErrorWriter, "\n")
}

//Info prints to bytes.Buffer for testing
func (u *MockUI) Info(message string) {
	u.Output(message)
}

//Output prints to bytes.Buffer for testing
func (u *MockUI) Output(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

//Running prints to bytes.Buffer for testing
func (u *MockUI) Running(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

//Success prints to bytes.Buffer for testing
func (u *MockUI) Success(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.Writer, message)
	fmt.Fprint(u.Writer, "\n")
}

//Warn prints to bytes.Buffer for testing
func (u *MockUI) Warn(message string) {
	u.once.Do(u.init)

	fmt.Fprint(u.ErrorWriter, message)
	fmt.Fprint(u.ErrorWriter, "\n")
}

//Log prefixes time then prints to bytes.Buffer for testing
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
