package wlog

import (
	"io"

	"github.com/daviddengcn/go-colortext"
)

//New returns a BasicUI for use
func New(reader io.Reader, writer, errorWriter io.Writer) *BasicUI {
	return &BasicUI{
		Reader:      reader,
		Writer:      writer,
		ErrorWriter: errorWriter,
	}
}

// AddConcurrent returns a new ConcurrentUI
func AddConcurrent(ui UI) *ConcurrentUI {
	return &ConcurrentUI{UI: ui}
}

//AddColor returns a ColorUI for use
func AddColor(logColor, outputColor, successColor, infoColor, errorColor, warnColor, runningColor ct.Color, ui UI) *ColorUI {
	return &ColorUI{
		LogFGColor:     logColor,
		LogBGColor:     ct.None,
		OutputFGColor:  outputColor,
		OutputBGColor:  ct.None,
		SuccessFGColor: successColor,
		SuccessBGColor: ct.None,
		InfoFGColor:    infoColor,
		InfoBGColor:    ct.None,
		ErrorFGColor:   errorColor,
		ErrorBGColor:   ct.None,
		WarnFGColor:    warnColor,
		WarnBGColor:    ct.None,
		RunningFGColor: runningColor,
		RunningBGColor: ct.None,
		UI:             ui,
	}
}

// AddPrefix returns a PrefixUI for use
func AddPrefix(logPre, outputPre, successPre, infoPre, errorPre, warnPre, runningPre string, ui UI) *PrefixUI {
	return &PrefixUI{
		LogPrefix:     logPre,
		OutputPrefix:  outputPre,
		SuccessPrefix: successPre,
		InfoPrefix:    infoPre,
		ErrorPrefix:   errorPre,
		WarnPrefix:    warnPre,
		RunningPrefix: runningPre,
		UI:            ui,
	}
}
