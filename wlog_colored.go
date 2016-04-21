package wlog

import "github.com/daviddengcn/go-colortext"

// ColorUI is a wrapper for UI that adds color.
type ColorUI struct {
	LogFGColor     Color
	OutputFGColor  Color
	SuccessFGColor Color
	InfoFGColor    Color
	ErrorFGColor   Color
	WarnFGColor    Color
	RunningFGColor Color
	LogBGColor     Color
	OutputBGColor  Color
	SuccessBGColor Color
	InfoBGColor    Color
	ErrorBGColor   Color
	WarnBGColor    Color
	RunningBGColor Color
	UI             UI
}

// Log prefixes to message before writing to Writer.
// LogFGColor and LogBGColor are used for color.
func (ui *ColorUI) Log(message string) {
	ct.ChangeColor(ui.LogFGColor.Code, ui.LogFGColor.Bright, ui.LogBGColor.Code, ui.LogBGColor.Bright)
	ui.UI.Log(message)
	ct.ResetColor()
}

// Output simply writes to Writer.
// OutputFGColor and OutputBGColor are used for color.
func (ui *ColorUI) Output(message string) {
	ct.ChangeColor(ui.OutputFGColor.Code, ui.OutputFGColor.Bright, ui.OutputBGColor.Code, ui.OutputBGColor.Bright)
	ui.UI.Output(message)
	ct.ResetColor()
}

// Success calls Output to write.
// Useful when you want separate colors or prefixes.
// SuccessFGColor and SuccessBGColor are used for color.
func (ui *ColorUI) Success(message string) {
	ct.ChangeColor(ui.SuccessFGColor.Code, ui.SuccessFGColor.Bright, ui.SuccessBGColor.Code, ui.SuccessBGColor.Bright)
	ui.UI.Success(message)
	ct.ResetColor()
}

// Info calls Output to write.
// Useful when you want separate colors or prefixes.
// InfoFGColor and InfoBGColor are used for color.
func (ui *ColorUI) Info(message string) {
	ct.ChangeColor(ui.InfoFGColor.Code, ui.InfoFGColor.Bright, ui.InfoBGColor.Code, ui.InfoBGColor.Bright)
	ui.UI.Info(message)
	ct.ResetColor()
}

// Error writes message to ErrorWriter.
// ErrorFGColor and ErrorBGColor are used for color.
func (ui *ColorUI) Error(message string) {
	ct.ChangeColor(ui.ErrorFGColor.Code, ui.ErrorFGColor.Bright, ui.ErrorBGColor.Code, ui.ErrorBGColor.Bright)
	ui.UI.Error(message)
	ct.ResetColor()
}

// Warn calls Error to write.
// Useful when you want separate colors or prefixes.
// WarnFGColor and WarnBGColor are used for color.
func (ui *ColorUI) Warn(message string) {
	ct.ChangeColor(ui.WarnFGColor.Code, ui.WarnFGColor.Bright, ui.WarnBGColor.Code, ui.WarnBGColor.Bright)
	ui.UI.Warn(message)
	ct.ResetColor()
}

// Running calls Output to write.
// Useful when you want separate colors or prefixes.
// RunningFGColor and RunningBGColor are used for color.
func (ui *ColorUI) Running(message string) {
	ct.ChangeColor(ui.RunningFGColor.Code, ui.RunningFGColor.Bright, ui.RunningBGColor.Code, ui.RunningBGColor.Bright)
	ui.UI.Running(message)
	ct.ResetColor()
}
