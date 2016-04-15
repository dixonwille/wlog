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
