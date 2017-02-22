package wlog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestColoredLog(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		t := time.Now()
		ui.Log("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		expectedString = t.Format(timeFormat) + ": " + expectedString
		assert.Equal(expectedString, out)
	}
}

func TestColoredOutput(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Output("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}

func TestColoredSuccess(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Success("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}

func TestColoredInfo(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Info("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}

func TestColoredRunning(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Running("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}

func TestColoredAsk(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("abc\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		res, err := ui.Ask("Awesome string", "")
		if err != nil {
			assert.Fail(err.Error())
		}
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
		assert.Equal("abc", res)
	}
}

func TestColoredError(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Error("Awesome string")
		out, err := errWriter.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = writer.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}

func TestColoredWarn(t *testing.T) {
	for _, c := range addColorCases {
		writer, errWriter, reader := initTest("\r\n")
		assert := assert.New(t)
		var ui UI
		ui = New(reader, writer, errWriter)
		ui = AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, ui)
		ui.Warn("Awesome string")
		out, err := errWriter.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = writer.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		assert.Equal(expectedString, out)
	}
}
