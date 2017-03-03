package util

import (
	"os"
	"strings"
	"testing"
)

func TestGetPathPaths(t *testing.T) {
	dirPath, fileName, absPath := GetPathParts("/bin/sh")
	if dirPath != "/bin" {
		t.Fail()
	}
	if fileName != "sh" {
		t.Fail()
	}
	if absPath != "/bin/sh" {
		t.Fail()
	}
}

func TestGetThisPathParts(t *testing.T) {
	dirPath, fileName, absPath := GetThisPathParts()
	if !strings.Contains(dirPath,
		"github.com/codedellemc/dvdcli/util/_test") {
		t.Fail()
	}
	if fileName != "util.test" {
		t.Fail()
	}
	if !strings.Contains(absPath,
		"github.com/codedellemc/dvdcli/util/_test/util.test") {
		t.Fail()
	}
}

func TestPrintVersion(t *testing.T) {
	PrintVersion(os.Stdout)
}
