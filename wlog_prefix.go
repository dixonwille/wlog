package wlog

// PrefixUI is a wrapper for UI that prefixes all strings.
// It does add a space betweem the prefix and message.
// If no prefix is specified ("") then it does not prefix the space.
type PrefixUI struct {
	LogPrefix     string
	OutputPrefix  string
	SuccessPrefix string
	InfoPrefix    string
	ErrorPrefix   string
	WarnPrefix    string
	RunningPrefix string
	AskPrefix     string
	UI            UI
}

// Log calls UI.Log to write.
// LogPrefix is used to prefix the message.
func (ui *PrefixUI) Log(message string) {
	if ui.LogPrefix != "" {
		message = ui.LogPrefix + " " + message
	}
	ui.UI.Log(message)
}

// Output calls UI.Output to write.
// OutputPrefix is used to prefix the message.
func (ui *PrefixUI) Output(message string) {
	if ui.OutputPrefix != "" {
		message = ui.OutputPrefix + " " + message
	}
	ui.UI.Output(message)
}

// Success calls UI.Success to write.
// Useful when you want separate colors or prefixes.
// SuccessPrefix is used to prefix the message.
func (ui *PrefixUI) Success(message string) {
	if ui.SuccessPrefix != "" {
		message = ui.SuccessPrefix + " " + message
	}
	ui.UI.Success(message)
}

// Info calls UI.Info to write.
// Useful when you want separate colors or prefixes.
// InfoPrefix is used to prefix the message.
func (ui *PrefixUI) Info(message string) {
	if ui.InfoPrefix != "" {
		message = ui.InfoPrefix + " " + message
	}
	ui.UI.Info(message)
}

// Error call UI.Error to write.
// ErrorPrefix is used to prefix the message.
func (ui *PrefixUI) Error(message string) {
	if ui.ErrorPrefix != "" {
		message = ui.ErrorPrefix + " " + message
	}
	ui.UI.Error(message)
}

// Warn calls UI.Warn to write.
// Useful when you want separate colors or prefixes.
// WarnPrefix is used to prefix message.
func (ui *PrefixUI) Warn(message string) {
	if ui.WarnPrefix != "" {
		message = ui.WarnPrefix + " " + message
	}
	ui.UI.Warn(message)
}

// Running calls Output to write.
// Useful when you want separate colors or prefixes.
// RunningPrefix is used to prefix message.
func (ui *PrefixUI) Running(message string) {
	if ui.RunningPrefix != "" {
		message = ui.RunningPrefix + " " + message
	}
	ui.UI.Running(message)
}

//Ask will call UI.Ask with message then wait for UI.Ask to return a response and/or error.
//It will clean the response by removing any carriage returns and new lines that if finds.
//If a message is not used ("") then it will not prompt user before waiting on a response.
//AskPrefix is used to prefix message.
func (ui *PrefixUI) Ask(message string) (string, error) {
	if ui.AskPrefix != "" {
		message = ui.AskPrefix + " " + message
	}
	res, err := ui.UI.Ask(message)
	return res, err
}
