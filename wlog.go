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
func AddColor(successColor, warnColor, infoColor, errorColor ct.Color, ui UI) *ColorUI {
	return &ColorUI{
		LogFGColor:     ct.None,
		LogBGColor:     ct.None,
		OutputFGColor:  ct.None,
		OutputBGColor:  ct.None,
		SuccessFGColor: successColor,
		SuccessBGColor: ct.None,
		InfoFGColor:    infoColor,
		InfoBGColor:    ct.None,
		ErrorFGColor:   errorColor,
		ErrorBGColor:   ct.None,
		WarnFGColor:    warnColor,
		WarnBGColor:    ct.None,
		UI:             ui,
	}
}

// AddPrefix returns a PrefixUI for use
func AddPrefix(logPre, outputPre, successPre, warnPre, infoPre, errorPre string, ui UI) *PrefixUI {
	return &PrefixUI{
		LogPrefix:     logPre,
		OutputPrefix:  outputPre,
		SuccessPrefix: successPre,
		WarnPrefix:    warnPre,
		InfoPrefix:    infoPre,
		ErrorPrefix:   errorPre,
		UI:            ui,
	}
}
