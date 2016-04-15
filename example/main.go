package main

import (
	"os"
	"runtime"

	"github.com/dixonwille/wlog"
)

//UI is used to print to screen
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

func main() {
	UI.Log("Log Statement")
	UI.Output("Output Statement")
	UI.Running("Running Statement")
	UI.Success("Success Statement")
	UI.Info("Info Statement")
	UI.Warn("Warning Statement")
	UI.Error("Error statement")
}
