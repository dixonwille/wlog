package wlog

import (
	"bytes"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	w := new(bytes.Buffer)
	e := new(bytes.Buffer)
	basic := New(os.Stdin, w, e)
	if basic.Reader != os.Stdin {
		t.Error("Making a new BasicUI failed to create reader")
	}
	if basic.Writer != w {
		t.Error("Making a new BasicUI failed to create Writer")
	}
	if basic.ErrorWriter != e {
		t.Error("Making a new BasicUI failed to create ErrorWriter")
	}
}

func Example() {
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
	// Output:
	// Info message
	// Output message
	// ~ Running message
	// ✓ Success message
	// ✗ Error message
	// ! Warning message
}

func TestAddColor(t *testing.T) {
	mock := &MockUI{
		Reader:      os.Stdin,
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	color := AddColor(None, Cyan, BrightBlue, Red, Green, Yellow, BrightMagenta, mock)

	if color.ErrorBGColor != None {
		t.Error("ErrorBGColor is not correct")
	}
	if color.InfoBGColor != None {
		t.Error("InfoBGColor is not correct")
	}
	if color.LogBGColor != None {
		t.Error("LogBGColor is not correct")
	}
	if color.OutputBGColor != None {
		t.Error("OutputBGColor is not correct")
	}
	if color.RunningBGColor != None {
		t.Error("RunningBGColor is not correct")
	}
	if color.SuccessBGColor != None {
		t.Error("SuccessBGColor is not correct")
	}
	if color.ErrorFGColor != Green {
		t.Error("ErrorFGColor is not correct")
	}
	if color.InfoFGColor != Red {
		t.Error("InfoFGColor is not correct")
	}
	if color.LogFGColor != None {
		t.Error("LogFGColor is not correct")
	}
	if color.OutputFGColor != Cyan {
		t.Error("OutputFGColor is not correct")
	}
	if color.RunningFGColor != BrightMagenta {
		t.Error("RunningFGColor is not correct")
	}
	if color.SuccessFGColor != BrightBlue {
		t.Error("SuccessFGColor is not correct")
	}
	if color.UI != mock {
		t.Error("UI was not set correctly")
	}
}

func TestAddPrefix(t *testing.T) {
	mock := &MockUI{
		Reader:      os.Stdin,
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	prefix := AddPrefix("*", "@", Check, "", Cross, "!", "~", mock)
	if prefix.ErrorPrefix != Cross {
		t.Error("ErrorPrefix was not set correctly")
	}
	if prefix.InfoPrefix != "" {
		t.Error("InfoPrefix was not set correctly")
	}
	if prefix.LogPrefix != "*" {
		t.Error("LogPrefix was not set correctly")
	}
	if prefix.OutputPrefix != "@" {
		t.Error("OutputPrefix was not set correctly")
	}
	if prefix.RunningPrefix != "~" {
		t.Error("RunningPrefix was not set correctly")
	}
	if prefix.SuccessPrefix != Check {
		t.Error("SuccessPrefix was not set correctly")
	}
	if prefix.WarnPrefix != "!" {
		t.Error("WarnPrefix was not set correctly")
	}
	if prefix.UI != mock {
		t.Error("UI was not set correctly")
	}
}

func TestAddConcurrent(t *testing.T) {
	mock := &MockUI{
		Reader:      os.Stdin,
		Writer:      new(bytes.Buffer),
		ErrorWriter: new(bytes.Buffer),
	}
	con := AddConcurrent(mock)
	if con.UI != mock {
		t.Error("UI was not set correctly")
	}
}
