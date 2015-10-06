package version_info

import (
	"testing"
)

func TestVersionInfo(t *testing.T) {
	t.Logf("Semver: %s", SemVer)
	t.Logf("ShaLong: %s", ShaLong)
	t.Logf("Epoch: %s", Epoch)
	t.Logf("Branch: %s", Branch)
	t.Logf("Arch: %s", Arch)
}

func TestEpochToRfc1123(t *testing.T) {
	t.Logf("EpochToRfc1123: %s", EpochToRfc1123())
}
