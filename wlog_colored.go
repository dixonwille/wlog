package wlog

import "github.com/daviddengcn/go-colortext"

// ColorUI is used to interface with the UI in color
type ColorUI struct {
	LogFGColor    ct.Color
	OutputFGColor ct.Color
	InfoFGColor   ct.Color
	ErrorFGColor  ct.Color
	FatalFGColor  ct.Color
	WarnFGColor   ct.Color
	LogBGColor    ct.Color
	OutputBGColor ct.Color
	InfoBGColor   ct.Color
	ErrorBGColor  ct.Color
	FatalBGColor  ct.Color
	WarnBGColor   ct.Color
	UI            UI
}

// Log prefixes timestamp and prints in color
func (ui *ColorUI) Log(message string) {
	ct.ChangeColor(ui.LogFGColor, false, ui.LogBGColor, false)
	ui.UI.Log(message)
	ct.ResetColor()
}

// Output prints to writer in color
func (ui *ColorUI) Output(message string) {
	ct.ChangeColor(ui.OutputFGColor, false, ui.OutputBGColor, false)
	ui.UI.Output(message)
	ct.ResetColor()
}

// Info prints to writer in color
func (ui *ColorUI) Info(message string) {
	ct.ChangeColor(ui.InfoFGColor, false, ui.InfoBGColor, false)
	ui.UI.Info(message)
	ct.ResetColor()
}

// Error prints to error in color
func (ui *ColorUI) Error(message string) {
	ct.ChangeColor(ui.ErrorFGColor, false, ui.ErrorBGColor, false)
	ui.UI.Error(message)
	ct.ResetColor()
}

// Fatal prints to error in color and exits the process
func (ui *ColorUI) Fatal(message string) {
	ct.ChangeColor(ui.FatalFGColor, false, ui.FatalBGColor, false)
	ui.UI.Fatal(message)
	ct.ResetColor()
}

// Warn prints to error in color
func (ui *ColorUI) Warn(message string) {
	ct.ChangeColor(ui.WarnFGColor, false, ui.WarnBGColor, false)
	ui.UI.Warn(message)
	ct.ResetColor()
}
