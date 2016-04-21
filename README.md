# WLog [![Build Status](https://travis-ci.org/dixonwille/wlog.svg?branch=master)](https://travis-ci.org/dixonwille/wlog) [![Go Report Card](https://goreportcard.com/badge/github.com/dixonwille/wlog)](https://goreportcard.com/report/github.com/dixonwille/wlog) [![GoDoc](https://godoc.org/github.com/dixonwille/wlog?status.svg)](https://godoc.org/github.com/dixonwille/wlog)

Package wlog creates simple to use UI structure. The UI is used to simply print
to the screen. There a wrappers that will wrap each other to create a good
looking UI. You can add color and prefixes as well as make it thread safe.

## Import
    import "github.com/dixonwille/wlog"

## Example

```go
var ui UI
ui = New(os.Stdin, os.Stdout, os.Stdout)
ui = AddPrefix("", "", Check, "", Cross, "!", "~", ui)
ui = AddConcurrent(ui)

ui.Info("Info message")
ui.Output("Output message")
ui.Running("Running message")
ui.Success("Success message")
ui.Error("Error message")
ui.Warn("Warning message")
```

Output:

```
Info message
Output message
~ Running message
✓ Success message
✗ Error message
! Warning message
```

On Windows it outputs to this (this includes color):

![winss](https://raw.githubusercontent.com/dixonwille/wlog/master/resources/winss.png)

On Mac it outputs to this (this includes color):

![macss](https://raw.githubusercontent.com/dixonwille/wlog/master/resources/macss.png)

## Usage

```go
const (

	//Check displays a checkmark
	Check = "\u2713"

	//Cross displays an x
	Cross = "\u2717"
)
```

```go
var (
	//BrightRed creates a bright red color
	BrightRed = Color{ct.Red, true}

	//BrightBlue creates a bright blue color
	BrightBlue = Color{ct.Blue, true}

	//BrightYellow creates a bright yellow color
	BrightYellow = Color{ct.Yellow, true}

	//Red creates a red color
	Red = Color{ct.Red, false}

	//Blue creaets a blue color
	Blue = Color{ct.Blue, false}

	//Yellow creates a yellow color
	Yellow = Color{ct.Yellow, false}

	//BrightGreen creates a bright green color
	BrightGreen = Color{ct.Green, true}

	//BrightCyan creates a bright cyan color
	BrightCyan = Color{ct.Cyan, true}

	//BrightMagenta creates a bright magenta color
	BrightMagenta = Color{ct.Magenta, true}

	//Green creates a green color
	Green = Color{ct.Green, false}

	//Cyan creates a cyan color
	Cyan = Color{ct.Cyan, false}

	//Magenta creates a magenta color
	Magenta = Color{ct.Magenta, false}

	//None does not change the color
	None = Color{ct.None, false}
)
```

#### type BasicUI

```go
type BasicUI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}
```

BasicUI simply writes/reads to correct input/output It is not thread safe.
Pretty simple to wrap your own functions around

#### func  New

```go
func New(reader io.Reader, writer, errorWriter io.Writer) *BasicUI
```
New creates a BasicUI. This should be the first function you call. This is not
thread safe and should only be used in serial applications.

#### func (*BasicUI) Error

```go
func (ui *BasicUI) Error(message string)
```
Error writes message to ErrorWriter.

#### func (*BasicUI) Info

```go
func (ui *BasicUI) Info(message string)
```
Info calls Output to write. Useful when you want separate colors or prefixes.

#### func (*BasicUI) Log

```go
func (ui *BasicUI) Log(message string)
```
Log prefixes to message before writing to Writer.

#### func (*BasicUI) Output

```go
func (ui *BasicUI) Output(message string)
```
Output simply writes to Writer.

#### func (*BasicUI) Running

```go
func (ui *BasicUI) Running(message string)
```
Running calls Output to write. Useful when you want separate colors or prefixes.

#### func (*BasicUI) Success

```go
func (ui *BasicUI) Success(message string)
```
Success calls Output to write. Useful when you want separate colors or prefixes.

#### func (*BasicUI) Warn

```go
func (ui *BasicUI) Warn(message string)
```
Warn calls Error to write. Useful when you want separate colors or prefixes.

#### type Color

```go
type Color struct {
	Code   ct.Color
	Bright bool
}
```

Color is a wrapper for go-colortext. Simplifies the use of this package by
assigning predefined colors that can be used.

#### type ColorUI

```go
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
```

ColorUI is a wrapper for UI that adds color.

#### func  AddColor

```go
func AddColor(logColor, outputColor, successColor, infoColor, errorColor, warnColor, runningColor Color, ui UI) *ColorUI
```
AddColor will wrap a colorful UI on top of ui. Use wlog's color variables for
the color. All background colors are not changed by this function but you are
able to change them manually. Just create this structure manually and change any
of the background colors you want.

#### func (*ColorUI) Error

```go
func (ui *ColorUI) Error(message string)
```
Error writes message to ErrorWriter. ErrorFGColor and ErrorBGColor are used for
color.

#### func (*ColorUI) Info

```go
func (ui *ColorUI) Info(message string)
```
Info calls Output to write. Useful when you want separate colors or prefixes.
InfoFGColor and InfoBGColor are used for color.

#### func (*ColorUI) Log

```go
func (ui *ColorUI) Log(message string)
```
Log prefixes to message before writing to Writer. LogFGColor and LogBGColor are
used for color.

#### func (*ColorUI) Output

```go
func (ui *ColorUI) Output(message string)
```
Output simply writes to Writer. OutputFGColor and OutputBGColor are used for
color.

#### func (*ColorUI) Running

```go
func (ui *ColorUI) Running(message string)
```
Running calls Output to write. Useful when you want separate colors or prefixes.
RunningFGColor and RunningBGColor are used for color.

#### func (*ColorUI) Success

```go
func (ui *ColorUI) Success(message string)
```
Success calls Output to write. Useful when you want separate colors or prefixes.
SuccessFGColor and SuccessBGColor are used for color.

#### func (*ColorUI) Warn

```go
func (ui *ColorUI) Warn(message string)
```
Warn calls Error to write. Useful when you want separate colors or prefixes.
WarnFGColor and WarnBGColor are used for color.

#### type ConcurrentUI

```go
type ConcurrentUI struct {
	UI UI
}
```

ConcurrentUI is a wrapper for UI that makes the UI thread safe.

#### func  AddConcurrent

```go
func AddConcurrent(ui UI) *ConcurrentUI
```
AddConcurrent will wrap a thread safe UI on top of ui. Safe to use inside of go
routines.

#### func (*ConcurrentUI) Error

```go
func (ui *ConcurrentUI) Error(message string)
```
Error writes message to ErrorWriter. This is a thread safe function.

#### func (*ConcurrentUI) Info

```go
func (ui *ConcurrentUI) Info(message string)
```
Info calls Output to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

#### func (*ConcurrentUI) Log

```go
func (ui *ConcurrentUI) Log(message string)
```
Log prefixes to message before writing to Writer. This is a thread safe
function.

#### func (*ConcurrentUI) Output

```go
func (ui *ConcurrentUI) Output(message string)
```
Output simply writes to Writer. This is a thread safe function.

#### func (*ConcurrentUI) Running

```go
func (ui *ConcurrentUI) Running(message string)
```
Running calls Output to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

#### func (*ConcurrentUI) Success

```go
func (ui *ConcurrentUI) Success(message string)
```
Success calls Output to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

#### func (*ConcurrentUI) Warn

```go
func (ui *ConcurrentUI) Warn(message string)
```
Warn calls Error to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

#### type MockUI

```go
type MockUI struct {
	Reader      io.Reader
	Writer      *bytes.Buffer
	ErrorWriter *bytes.Buffer
}
```

MockUI is a UI that is used for tests. It is exported publicly for use in
external tests if needed as well.

#### func (*MockUI) Error

```go
func (u *MockUI) Error(message string)
```
Error writes message to ErrorWriter.

#### func (*MockUI) Info

```go
func (u *MockUI) Info(message string)
```
Info calls Output to write. Useful when you want separate colors or prefixes.

#### func (*MockUI) Log

```go
func (u *MockUI) Log(message string)
```
Log prefixes to message before writing to Writer.

#### func (*MockUI) Output

```go
func (u *MockUI) Output(message string)
```
Output simply writes to Writer.

#### func (*MockUI) Running

```go
func (u *MockUI) Running(message string)
```
Running calls Output to write. Useful when you want separate colors or prefixes.

#### func (*MockUI) Success

```go
func (u *MockUI) Success(message string)
```
Success calls Output to write. Useful when you want separate colors or prefixes.

#### func (*MockUI) Warn

```go
func (u *MockUI) Warn(message string)
```
Warn calls Error to write. Useful when you want separate colors or prefixes.

#### type PrefixUI

```go
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
```

PrefixUI is a wrapper for UI that prefixes all strings. It does add a space
betweem the prefix and message. If no prefix is specified ("") then it does not
prefix the space.

#### func  AddPrefix

```go
func AddPrefix(logPre, outputPre, successPre, infoPre, errorPre, warnPre, runningPre string, ui UI) *PrefixUI
```
AddPrefix will wrap a UI that will prefix the message on top of ui. If a prefix
is set to nothing ("") then there will be no prefix for that message type.

#### func (*PrefixUI) Error

```go
func (ui *PrefixUI) Error(message string)
```
Error writes message to ErrorWriter. ErrorPrefix is used to prefix the message.

#### func (*PrefixUI) Info

```go
func (ui *PrefixUI) Info(message string)
```
Info calls Output to write. Useful when you want separate colors or prefixes.
InfoPrefix is used to prefix the message.

#### func (*PrefixUI) Log

```go
func (ui *PrefixUI) Log(message string)
```
Log prefixes to message before writing to Writer. LogPrefix is used to prefix
the message.

#### func (*PrefixUI) Output

```go
func (ui *PrefixUI) Output(message string)
```
Output simply writes to Writer. OutputPrefix is used to prefix the message.

#### func (*PrefixUI) Running

```go
func (ui *PrefixUI) Running(message string)
```
Running calls Output to write. Useful when you want separate colors or prefixes.
RunningPrefix is used to prefix message.

#### func (*PrefixUI) Success

```go
func (ui *PrefixUI) Success(message string)
```
Success calls Output to write. Useful when you want separate colors or prefixes.
SuccessPrefix is used to prefix the message.

#### func (*PrefixUI) Warn

```go
func (ui *PrefixUI) Warn(message string)
```
Warn calls Error to write. Useful when you want separate colors or prefixes.
WarnPrefix is used to prefix message.

#### type UI

```go
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
```

UI simply writes to an io.Writer with a new line appended to each call.
