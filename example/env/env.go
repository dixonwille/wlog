package env

import (
	"os"
	"runtime"

	"github.com/dixonwille/wlog"
)

//UI is used to print to Screen
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
