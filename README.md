# WLog
This is a simple package to allow a user to print to their terminal with style.

## How I got the idea
I used Mitchellh's [CLI](https://github.com/mitchellh/cli) structure and wrapping for the different structures. It was a clean look and feel and pretty simple to use. But I didn't want all the other cli stuff that came with the package so I created this.

For color I use DavidDenGCN's [Go-ColorText](https://github.com/daviddengcn/go-colortext). He claims that it colors the text and is cross platform (only tested under windows). I made a wrapper with all possible color combinations with his package. So you only have to import this package (one less line).

## How To Use
Here is an example I used. I usually put this in a package outside of my main package called env (you can use any package name). For examples I will be acting like it is in the `env` package.
```
var UI wlog.UI
func init() {
	var prefixed wlog.UI
	basic := wlog.New(os.Stdin, os.Stdout, os.Stderr)
	if runtime.GOOS != "windows" {
		prefixed = wlog.AddPrefix("", "", "["+wlog.Check+"]", "", "["+wlog.Cross+"]", "[!]", "[~]", basic)
	} else {
		prefixed = wlog.AddPrefix("", "", "[$]", "", "[x]", "[!]", "[~]", basic)
	}
	colored := wlog.AddColor(wlog.None, wlog.None, wlog.BrightGreen, wlog.BrightCyan, wlog.BrightRed, wlog.BrightYellow, wlog.BrightMagenta, prefixed)
	UI = wlog.AddConcurrent(colored)
}
```

1. I first create a BasicUI using `New(in io.Writer, out io.Writer, err io.Writer)`
2. I can add Prefixes to the different commands by using `AddPrefix(logPre, outputPre, successPre, infoPre, errorPre, warnPre, runningPre string, ui UI)` (the UI is the basic one we created)
3. Adding Color is just as simple `AddColor(logColor, outputColor, successColor, infoColor, errorColor, warnColor, runningColor Color, ui UI)` (Replace UI with the recent UI created in this case PrefixUI)
4. Added Concurrency to allow go routines to use `AddConcurrent(ui UI)` (used ColorUI)
5. Stored the final ui to a global variable to use across the program.

In the example above I use different prefixes depending on the OS that is running. Windows does not support Unicode characters in terminal so I have two different sets of prefixes. I did include constants for a few (two) useful prefixes: A Check mark and an X or Cross.

To print to screen you can use one of the following:

```
func main() {
	env.UI.Log("Log Statement")
	env.UI.Output("Output Statement")
	env.UI.Running("Running Statement")
	env.UI.Success("Success Statement")
	env.UI.Info("Info Statement")
	env.UI.Warn("Warning Statement")
	env.UI.Error("Error statement")
}
```
On Windows it outputs to this:

![winss](https://gihub.com/dixonwille/wlog/raw/master/resources/winss.png)

On Mac it outputs to this (notice the prefix change):

![winss](https://gihub.com/dixonwille/wlog/raw/master/resources/macss.png)

## TODO
* Add a table UI to allow a user to create a table. Use [this](https://github.com/olekukonko/tablewriter).
