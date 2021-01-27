package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/fatih/color"
	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var errChecking = errors.New("error occurred")

func TestCheckErrorText(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "text")
	CheckError(errChecking, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(`error occurred`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestCheckErrorNilText(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "text")
	CheckError(nil, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(``)
	assert.True(t, re.Match(b.Bytes()))
}

func TestCheckErrorDefault(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "dummy")
	CheckError(errChecking, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(`unknown type format dummy. Hint: use --output json|text`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestCheckErrorNilDefault(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "dummy")
	CheckError(nil, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(``)
	assert.True(t, re.Match(b.Bytes()))
}

func TestCheckErrorJSON(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "json")
	CheckError(errChecking, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(`{"error":{},"detail":"error occurred"}`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestCheckErrorNilJSON(t *testing.T) {
	defer func(a func()) { ErrAction = a }(ErrAction)
	defer func(a io.Writer) { color.Output = a }(color.Output)

	var b bytes.Buffer
	defer func(b bytes.Buffer) { b.Reset() }(b)
	w := bufio.NewWriter(&b)
	color.Output = w
	ErrAction = func() {}

	viper.Set(config.ArgOutput, "json")
	CheckError(nil, w)
	err := w.Flush()
	assert.NoError(t, err)

	re := regexp.MustCompile(``)
	assert.True(t, re.Match(b.Bytes()))
}

func TestNewRequiredFlagErr(t *testing.T) {
	requiredFlag := NewRequiredFlagErr("dummy")
	assert.True(t, requiredFlag.Error() == "dummy required flag is not set\n")
}

func TestVarErrAction(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		ErrAction()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestVarErrAction")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
