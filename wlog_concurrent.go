package wlog

import "sync"

// ConcurrentUI Wraps arround UI interface.
// Makes the UI concurrent.
type ConcurrentUI struct {
	UI UI
	l  sync.Mutex
}

// Log prefixes timestamp then writes to writer
func (ui *ConcurrentUI) Log(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Log(message)
}

// Output writes to writer
func (ui *ConcurrentUI) Output(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Output(message)
}

// Info writes to writer
func (ui *ConcurrentUI) Info(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Info(message)
}

// Error writes to error
func (ui *ConcurrentUI) Error(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Error(message)
}

// Warn writes to error
func (ui *ConcurrentUI) Warn(message string) {
	ui.l.Lock()
	defer ui.l.Unlock()
	ui.UI.Warn(message)
}
