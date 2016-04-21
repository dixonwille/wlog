package wlog

import "sync"

// ConcurrentUI is a wrapper for UI that makes the UI thread safe.
type ConcurrentUI struct {
	UI UI
	l  sync.Mutex
}

// Log prefixes to message before writing to Writer.
// This is a thread safe function.
func (ui *ConcurrentUI) Log(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Log(message)
}

// Output simply writes to Writer.
// This is a thread safe function.
func (ui *ConcurrentUI) Output(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Output(message)
}

// Success calls Output to write.
// Useful when you want seperate colors or prefixes.
// This is a thread safe function.
func (ui *ConcurrentUI) Success(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Success(message)
}

// Info calls Output to write.
// Useful when you want seperate colors or prefixes.
// This is a thread safe function.
func (ui *ConcurrentUI) Info(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Info(message)
}

// Error writes message to ErrorWriter.
// This is a thread safe function.
func (ui *ConcurrentUI) Error(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Error(message)
}

// Warn calls Error to write.
// Useful when you want seperate colors or prefixes.
// This is a thread safe function.
func (ui *ConcurrentUI) Warn(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Warn(message)
}

// Running calls Output to write.
// Useful when you want seperate colors or prefixes.
// This is a thread safe function.
func (ui *ConcurrentUI) Running(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Running(message)
}
