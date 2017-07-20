# WLog [![Build Status](https://travis-ci.org/dixonwille/wlog.svg?branch=master)](https://travis-ci.org/dixonwille/wlog) [![Go Report Card](https://goreportcard.com/badge/github.com/dixonwille/wlog)](https://goreportcard.com/report/github.com/dixonwille/wlog) [![GoDoc](https://godoc.org/github.com/dixonwille/wlog?status.svg)](https://godoc.org/github.com/dixonwille/wlog) [![Coverage Status](https://coveralls.io/repos/github/dixonwille/wmenu/badge.svg?branch=master)](https://coveralls.io/github/dixonwille/wmenu?branch=master)

Package wlog creates simple to use UI structure. The UI is used to simply print
to the screen. There a wrappers that will wrap each other to create a good
looking UI. You can add color and prefixes as well as make it thread safe.

## Import

It is recommended to use `govendor` or any vendoring tool allowing you to specify which version of this package you would like to use.

```go
import "github.com/dixonwille/wlog"
```

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
ui = AddPrefix("?", Cross, " ", "", "", "~", Check, "!", ui)
ui = AddConcurrent(ui)

ui.Ask("Ask question", "")
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

Please use the [Documentaion](https://godoc.org/github.com/dixonwille/wlog) to read more into how to use this package.
