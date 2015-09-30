package util

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"

	"github.com/bugsnag/osext"
	version "github.com/emccode/dvdcli/version_info"
)

var thisExeAbsPath string

func init() {
	_, _, thisExeAbsPath = GetThisPathParts()
}

func GetPathParts(path string) (dirPath, fileName, absPath string) {
	lookup, lookupErr := exec.LookPath(path)
	if lookupErr == nil {
		path = lookup
	}
	absPath, _ = filepath.Abs(path)
	dirPath = filepath.Dir(absPath)
	fileName = filepath.Base(absPath)
	return
}

func GetThisPathParts() (dirPath, fileName, absPath string) {
	exeFile, err := osext.Executable()
	if err != nil {
		panic(err)
	}
	return GetPathParts(exeFile)
}

func PrintVersion(out io.Writer) {
	fmt.Fprintf(out, "Binary: %s\n", thisExeAbsPath)
	fmt.Fprintf(out, "SemVer: %s\n", version.SemVer)
	fmt.Fprintf(out, "OsArch: %s\n", version.Arch)
	fmt.Fprintf(out, "Branch: %s\n", version.Branch)
	fmt.Fprintf(out, "Commit: %s\n", version.ShaLong)
	fmt.Fprintf(out, "Formed: %s\n", version.EpochToRfc1123())
}
