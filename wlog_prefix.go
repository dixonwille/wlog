package wlog

// PrefixUI adds a prefix before printing to screen
type PrefixUI struct {
	LogPrefix     string
	OutputPrefix  string
	SuccessPrefix string
	InfoPrefix    string
	ErrorPrefix   string
	WarnPrefix    string
	RunningPrefix string
	UI            UI
}

// Log prefixs and writes to log
func (ui *PrefixUI) Log(message string) {
	if ui.LogPrefix != "" {
		message = ui.LogPrefix + " " + message
	}
	ui.UI.Log(message)
}

// Output prefixs and writes to output
func (ui *PrefixUI) Output(message string) {
	if ui.OutputPrefix != "" {
		message = ui.OutputPrefix + " " + message
	}
	ui.UI.Output(message)
}

// Success prefixs and writes to success
func (ui *PrefixUI) Success(message string) {
	if ui.SuccessPrefix != "" {
		message = ui.SuccessPrefix + " " + message
	}
	ui.UI.Output(message)
}

// Info prefixs and writes to info
func (ui *PrefixUI) Info(message string) {
	if ui.InfoPrefix != "" {
		message = ui.InfoPrefix + " " + message
	}
	ui.UI.Output(message)
}

// Error prefixs and writes to error
func (ui *PrefixUI) Error(message string) {
	if ui.ErrorPrefix != "" {
		message = ui.ErrorPrefix + " " + message
	}
	ui.UI.Error(message)
}

// Warn prefixs and writes to warn
func (ui *PrefixUI) Warn(message string) {
	if ui.WarnPrefix != "" {
		message = ui.WarnPrefix + " " + message
	}
	ui.UI.Warn(message)
}

// Running prefixs and writes to running
func (ui *PrefixUI) Running(message string) {
	if ui.RunningPrefix != "" {
		message = ui.RunningPrefix + " " + message
	}
	ui.UI.Running(message)
}
