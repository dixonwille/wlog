package wlog

// UI is an interface for printing lines to the terminal.
// All of these will append a new line.
type UI interface {
	// Log prefixes the string with Date-Time Info.
	Log(string)
	// Output is a standard output to screen.
	Output(string)
	// Info is exactly like output.
	// Reason for having two is to allow different colors or files.
	Info(string)
	// Error will write to ErrorWriter.
	// This will not kill the program.
	Error(string)
	// Warn will write to ErrorWriter much like Error.
	// Reason for having two is to allow different colors or files.
	Warn(string)
}
