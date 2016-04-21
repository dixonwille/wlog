package wlog

// UI simply writes to an io.Writer with a new line appended to each call.
type UI interface {
	// Log prefixes to message before writing to Writer.
	Log(string)

	// Output simply writes to Writer.
	Output(string)

	// Success calls Output to write.
	// Useful when you want separate colors or prefixes.
	Success(string)

	// Info calls Output to write.
	// Useful when you want separate colors or prefixes.
	Info(string)

	// Error writes message to ErrorWriter.
	Error(string)

	// Warn calls Error to write.
	// Useful when you want separate colors or prefixes.
	Warn(string)

	// Running calls Output to write.
	// Useful when you want separate colors or prefixes.
	Running(string)
}
