package wlog

// UI simply writes to an io.Writer with a new line appended to each call.
// It also has the ability to ask a question and return a response.
type UI interface {
	// Log writes a timestamped message to the writer
	Log(string)
	// Output writes a message to the writer
	Output(string)
	// Success writes a message indicating an success message
	Success(string)
	// Info writes a message indicating an informational message
	Info(string)
	// Error writes a message indicating an error
	Error(string)
	// Warn writes a message indicating a warning
	Warn(string)
	// Running writes a message indicating a process is running
	Running(string)
	// Ask writes a message to the writer and reads the user's input
	Ask(string, string) (string, error)
}
