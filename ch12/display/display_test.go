package display

import (
	"os"
	"reflect"
	"testing"
)

func TestDisplayStderr(t *testing.T) {
	Display("os.Stderr", os.Stderr)
}

func TestDisplayStderrValue(t *testing.T) {
	Display("rV", reflect.ValueOf(os.Stderr))
}
