package wlog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrefixConcurrentLog(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui = AddConcurrent(ui)
		t := time.Now()
		ui.Log("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.log == " " {
			expectedString = t.Format(timeFormat) + ": " + c.log + expectedString
		} else if c.log == "" {
			expectedString = t.Format(timeFormat) + ": " + expectedString
		} else {
			expectedString = t.Format(timeFormat) + ": " + c.log + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixOutput(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Output("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.out == " " {
			expectedString = c.out + expectedString
		} else if c.out != "" {
			expectedString = c.out + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixSuccess(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Success("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.suc == " " {
			expectedString = c.suc + expectedString
		} else if c.suc != "" {
			expectedString = c.suc + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixInfo(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Info("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.inf == " " {
			expectedString = c.inf + expectedString
		} else if c.inf != "" {
			expectedString = c.inf + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixRunning(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Running("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.run == " " {
			expectedString = c.run + expectedString
		} else if c.run != "" {
			expectedString = c.run + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixError(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Error("Awesome string")
		out, err := errWriter.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = writer.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.err == " " {
			expectedString = c.err + expectedString
		} else if c.err != "" {
			expectedString = c.err + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixWarn(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Warn("Awesome string")
		out, err := errWriter.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = writer.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.war == " " {
			expectedString = c.war + expectedString
		} else if c.war != "" {
			expectedString = c.war + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixWarnNoErrorWriter(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("\r\n")
		ui = New(reader, writer, nil)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		ui.Warn("Awesome string")
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.war == " " {
			expectedString = c.war + expectedString
		} else if c.war != "" {
			expectedString = c.war + " " + expectedString
		}
		assert.Equal(expectedString, out)
	}
}

func TestPrefixAsk(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("345\r\n")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
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
		if c.ask == " " {
			expectedString = c.ask + expectedString
		} else if c.ask != "" {
			expectedString = c.ask + " " + expectedString
		}
		assert.Equal(expectedString, out)
		assert.Equal("345", res)
	}
}

func TestPrefixAskError(t *testing.T) {
	for _, c := range addPrefixCases {
		assert := assert.New(t)
		var ui UI
		writer, errWriter, reader := initTest("")
		ui = New(reader, writer, errWriter)
		ui = AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, ui)
		res, err := ui.Ask("Awesome string", "")
		if err != nil {
			assert.Equal("EOF", err.Error())
		}
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		assert.Equal("EOF", err.Error())
		expectedString := "Awesome string\n"
		if c.ask == " " {
			expectedString = c.ask + expectedString
		} else if c.ask != "" {
			expectedString = c.ask + " " + expectedString
		}
		assert.Equal(expectedString, out)
		assert.Equal("", res)
	}
}
