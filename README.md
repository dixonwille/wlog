# WLog [![Build Status](https://travis-ci.org/dixonwille/wlog.svg?branch=master)](https://travis-ci.org/dixonwille/wlog) [![Go Report Card](https://goreportcard.com/badge/github.com/dixonwille/wlog)](https://goreportcard.com/report/github.com/dixonwille/wlog) [![GoDoc](https://godoc.org/github.com/dixonwille/wlog?status.svg)](https://godoc.org/github.com/dixonwille/wlog)

Package wlog creates simple to use UI structure. The UI is used to simply print
to the screen. There a wrappers that will wrap each other to create a good
looking UI. You can add color and prefixes as well as make it thread safe.

## Import
    import "github.com/dixonwille/wlog"

## Idea Behind WLog

I used Mitchellh's [CLI](https://github.com/mitchellh/cli) structure and
 wrapping for the different structures. It was a clean look and feel. Plus it
 was pretty simple to use. But I didn't want all the other cli stuff that came
 with the package so I created this.

For color I use DavidDenGCN's
[Go-ColorText](https://github.com/daviddengcn/go-colortext). His color package
allows for color that is available cross-platforms. I made a wrapper with all
possible color combinations with his package. So you only have to import this
package (one less line).


## Example

```go
var ui UI
reader := strings.NewReader("User Input\r\n") //Simulate user typing "User Input" then pressing [enter] when reading from os.Stdin
ui = New(reader, os.Stdout, os.Stdout)
ui = AddPrefix("?", Cross, "", "", "", "~", Check, "!", ui)
ui = AddConcurrent(ui)

ui.Ask("Ask question")
ui.Error("Error message")
ui.Info("Info message")
ui.Output("Output message")
ui.Running("Running message")
ui.Success("Success message")
ui.Warn("Warning message")
```

Output:

```
? Ask question
✗ Error message
Info message
Output message
~ Running message
✓ Success message
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

	//White creates a white color
	White = Color{ct.White, false}

	//BrightWhite creates a bright white color
	BrightWhite = Color{ct.White, true}

	//Black creates a black color
	Black = Color{ct.Black, false}

	//BrightBlack creates a bright black color
	BrightBlack = Color{ct.Black, true}

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

#### func (*BasicUI) Ask

```go
func (ui *BasicUI) Ask(message string) (string, error)
```
Ask will call output with message then wait for Reader to print newline (\n). If
Reader is os.Stdin then that is when ever a user presses [enter]. It will clean
the response by removing any carriage returns and new lines that if finds. If a
message is not used ("") then it will not prompt user before waiting on a
response.

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
	LogFGColor      Color
	OutputFGColor   Color
	SuccessFGColor  Color
	InfoFGColor     Color
	ErrorFGColor    Color
	WarnFGColor     Color
	RunningFGColor  Color
	AskFGColor      Color
	ResponseFGColor Color
	LogBGColor      Color
	OutputBGColor   Color
	SuccessBGColor  Color
	InfoBGColor     Color
	ErrorBGColor    Color
	WarnBGColor     Color
	RunningBGColor  Color
	AskBGColor      Color
	ResponseBGColor Color
	UI              UI
}
```

ColorUI is a wrapper for UI that adds color.

#### func  AddColor

```go
func AddColor(askColor, errorColor, infoColor, logColor, outputColor, responseColor, runningColor, successColor, warnColor Color, ui UI) *ColorUI
```
AddColor will wrap a colorful UI on top of ui. Use wlog's color variables for
the color. All background colors are not changed by this function but you are
able to change them manually. Just create this structure manually and change any
of the background colors you want. Arguments are in alphabetical order.

#### func (*ColorUI) Ask

```go
func (ui *ColorUI) Ask(message string) (string, error)
```
Ask will call UI.Output with message then wait for UI.Ask to return a response
and/or error. It will clean the response by removing any carriage returns and
new lines that if finds. If a message is not used ("") then it will not prompt
user before waiting on a response. AskFGColor and AskBGColor are used for
message color. ResponseFGColor and ResponseBGColor are used for response color.

#### func (*ColorUI) Error

```go
func (ui *ColorUI) Error(message string)
```
Error calls UI.Error to write. ErrorFGColor and ErrorBGColor are used for color.

#### func (*ColorUI) Info

```go
func (ui *ColorUI) Info(message string)
```
Info calls UI.Info to write. Useful when you want separate colors or prefixes.
InfoFGColor and InfoBGColor are used for color.

#### func (*ColorUI) Log

```go
func (ui *ColorUI) Log(message string)
```
Log calls UI.Log to write. LogFGColor and LogBGColor are used for color.

#### func (*ColorUI) Output

```go
func (ui *ColorUI) Output(message string)
```
Output calls UI.Output to write. OutputFGColor and OutputBGColor are used for
color.

#### func (*ColorUI) Running

```go
func (ui *ColorUI) Running(message string)
```
Running calls UI.Running to write. Useful when you want separate colors or
prefixes. RunningFGColor and RunningBGColor are used for color.

#### func (*ColorUI) Success

```go
func (ui *ColorUI) Success(message string)
```
Success calls UI.Success to write. Useful when you want separate colors or
prefixes. SuccessFGColor and SuccessBGColor are used for color.

#### func (*ColorUI) Warn

```go
func (ui *ColorUI) Warn(message string)
```
Warn calls UI.Warn to write. Useful when you want separate colors or prefixes.
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

#### func (*ConcurrentUI) Ask

```go
func (ui *ConcurrentUI) Ask(message string) (string, error)
```
Ask will call UI.Ask with message then wait for UI.Ask to return a response
and/or error. It will clean the response by removing any carriage returns and
new lines that if finds. If a message is not used ("") then it will not prompt
user before waiting on a response. This is a thread safe function.

#### func (*ConcurrentUI) Error

```go
func (ui *ConcurrentUI) Error(message string)
```
Error calls UI.Error to write. This is a thread safe function.

#### func (*ConcurrentUI) Info

```go
func (ui *ConcurrentUI) Info(message string)
```
Info calls UI.Info to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

#### func (*ConcurrentUI) Log

```go
func (ui *ConcurrentUI) Log(message string)
```
Log calls UI.Log to write. This is a thread safe function.

#### func (*ConcurrentUI) Output

```go
func (ui *ConcurrentUI) Output(message string)
```
Output calls UI.Output to write. This is a thread safe function.

#### func (*ConcurrentUI) Running

```go
func (ui *ConcurrentUI) Running(message string)
```
Running calls UI.Running to write. Useful when you want separate colors or
prefixes. This is a thread safe function.

#### func (*ConcurrentUI) Success

```go
func (ui *ConcurrentUI) Success(message string)
```
Success calls UI.Success to write. Useful when you want separate colors or
prefixes. This is a thread safe function.

#### func (*ConcurrentUI) Warn

```go
func (ui *ConcurrentUI) Warn(message string)
```
Warn calls UI.Warn to write. Useful when you want separate colors or prefixes.
This is a thread safe function.

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
	AskPrefix     string
	UI            UI
}
```

PrefixUI is a wrapper for UI that prefixes all strings. It does add a space
betweem the prefix and message. If no prefix is specified ("") then it does not
prefix the space.

#### func  AddPrefix

```go
func AddPrefix(askPre, errorPre, infoPre, logPre, outputPre, runningPre, successPre, warnPre string, ui UI) *PrefixUI
```
AddPrefix will wrap a UI that will prefix the message on top of ui. If a prefix
is set to nothing ("") then there will be no prefix for that message type.
Arguments are in alphabetical order.

#### func (*PrefixUI) Ask

```go
func (ui *PrefixUI) Ask(message string) (string, error)
```
Ask will call UI.Ask with message then wait for UI.Ask to return a response
and/or error. It will clean the response by removing any carriage returns and
new lines that if finds. If a message is not used ("") then it will not prompt
user before waiting on a response. AskPrefix is used to prefix message.

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
	Log(string)
	Output(string)
	Success(string)
	Info(string)
	Error(string)
	Warn(string)
	Running(string)
	Ask(string) (string, error)
}
```

UI simply writes to an io.Writer with a new line appended to each call. It also
has the ability to ask a question and return a response.
