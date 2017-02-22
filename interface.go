package wlog

// UI simply writes to an io.Writer with a new line appended to each call.
// It also has the ability to ask a question and return a response.
type UI interface {
	Log(string)
	Output(string)
	Success(string)
	Info(string)
	Error(string)
	Warn(string)
	Running(string)
	Ask(string, string) (string, error)
}
