package wlog

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newCases = []struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}{
	{strings.NewReader("Hello Worlds\r\n"), os.Stdout, nil},
	{nil, nil, nil},
	{os.Stdin, nil, os.Stderr},
}
var addColorCases = []struct {
	logColor      Color
	outputColor   Color
	successColor  Color
	infoColor     Color
	errorColor    Color
	warnColor     Color
	runningColor  Color
	askColor      Color
	responseColor Color
}{
	{None, Blue, Green, Red, Yellow, Cyan, Magenta, White, Black},
	{None, None, None, None, None, None, None, None, None},
	{BrightBlue, BrightGreen, BrightRed, BrightYellow, BrightCyan, BrightMagenta, BrightWhite, BrightBlack, None},
}
var addPrefixCases = []struct {
	ask string
	err string
	inf string
	log string
	out string
	suc string
	run string
	war string
}{
	{"", "", "", "", "", "", "", ""},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
	{Cross, Check, "!", "~", "@", "#", "+", "="},
	{"%", "^", "&", "*", "@", ":", ",", "?"},
}

var trimCases = []struct {
	input string
	trim  string
}{
	{" 123 \r\n", " "},
	{"  123 \r\n", " "},
	{"!!123!!\r\n", "!"},
	{" !!123!! \r\n", "! "},
}

func Example() {
	var ui UI
	reader := strings.NewReader("User Input\r\n") //Simulate user typing "User Input" then pressing [enter] when reading from os.Stdin
	ui = New(reader, os.Stdout, os.Stdout)
	ui = AddPrefix("?", Cross, " ", "", "", "~", Check, "!", ui)
	ui = AddConcurrent(ui)

	ui.Ask("Ask question", " ")
	ui.Error("Error message")
	ui.Info("Info message")
	ui.Output("Output message")
	ui.Running("Running message")
	ui.Success("Success message")
	ui.Warn("Warning message")

	// Output:
	// ? Ask question
	// ✗ Error message
	//  Info message
	// Output message
	// ~ Running message
	// ✓ Success message
	// ! Warning message

}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	for _, c := range newCases {
		basic := New(c.in, c.out, c.err)
		assert.Equal(c.in, basic.Reader)
		assert.Equal(c.out, basic.Writer)
		assert.Equal(c.err, basic.ErrorWriter)
	}
}

func TestAskTrim(t *testing.T) {
	assert := assert.New(t)
	for _, c := range trimCases {
		writer, errWriter, in := initTest(c.input)
		basic := New(in, writer, errWriter)
		res, err := basic.Ask("Awesome string", c.trim)
		if err != nil {
			assert.Fail(err.Error())
		}
		out, err := writer.ReadString((byte)('\n'))
		if err != nil {
			assert.Fail(err.Error())
		}
		_, err = errWriter.ReadString((byte)('\n'))
		expectedString := "Awesome string\n"
		assert.Equal("EOF", err.Error())
		assert.Equal(expectedString, out)
		expectedAns := strings.Replace(c.input, "\r", "", -1)
		expectedAns = strings.Replace(expectedAns, "\n", "", -1)
		expectedAns = strings.Trim(expectedAns, c.trim)
		assert.Equal(expectedAns, res)
	}
}

func TestAddColor(t *testing.T) {
	assert := assert.New(t)
	basic := New(os.Stdin, os.Stdout, os.Stderr)
	for _, c := range addColorCases {
		color := AddColor(c.askColor, c.errorColor, c.infoColor, c.logColor, c.outputColor, c.responseColor, c.runningColor, c.successColor, c.warnColor, basic)
		assert.Equal(None, color.AskBGColor)
		assert.Equal(None, color.ErrorBGColor)
		assert.Equal(None, color.InfoBGColor)
		assert.Equal(None, color.LogBGColor)
		assert.Equal(None, color.OutputBGColor)
		assert.Equal(None, color.ResponseBGColor)
		assert.Equal(None, color.RunningBGColor)
		assert.Equal(None, color.SuccessBGColor)
		assert.Equal(None, color.WarnBGColor)
		assert.Equal(c.askColor, color.AskFGColor)
		assert.Equal(c.errorColor, color.ErrorFGColor)
		assert.Equal(c.infoColor, color.InfoFGColor)
		assert.Equal(c.logColor, color.LogFGColor)
		assert.Equal(c.outputColor, color.OutputFGColor)
		assert.Equal(c.responseColor, color.ResponseFGColor)
		assert.Equal(c.runningColor, color.RunningFGColor)
		assert.Equal(c.successColor, color.SuccessFGColor)
		assert.Equal(c.warnColor, color.WarnFGColor)
		assert.Equal(basic, color.UI)
	}
}

func TestAddPrefix(t *testing.T) {
	assert := assert.New(t)
	basic := New(os.Stdin, os.Stdout, os.Stderr)
	for _, c := range addPrefixCases {
		prefix := AddPrefix(c.ask, c.err, c.inf, c.log, c.out, c.run, c.suc, c.war, basic)
		assert.Equal(c.ask, prefix.AskPrefix)
		assert.Equal(c.err, prefix.ErrorPrefix)
		assert.Equal(c.inf, prefix.InfoPrefix)
		assert.Equal(c.log, prefix.LogPrefix)
		assert.Equal(c.out, prefix.OutputPrefix)
		assert.Equal(c.run, prefix.RunningPrefix)
		assert.Equal(c.suc, prefix.SuccessPrefix)
		assert.Equal(c.war, prefix.WarnPrefix)
		assert.Equal(basic, prefix.UI)
	}
}

func TestAddConcurrent(t *testing.T) {
	basic := New(os.Stdin, os.Stdout, os.Stderr)
	con := AddConcurrent(basic)
	assert.Equal(t, basic, con.UI)
}

func initTest(input string) (*bytes.Buffer, *bytes.Buffer, io.Reader) {
	var b []byte
	var e []byte
	return bytes.NewBuffer(b), bytes.NewBuffer(e), strings.NewReader(input)
}
