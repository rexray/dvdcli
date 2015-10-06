package main

import (
	"testing"
)

func TestCreateCmd(t *testing.T) {
	createCmd.Execute()
}

func TestRemoveCmd(t *testing.T) {
	removeCmd.Execute()
}

func TestMountCmd(t *testing.T) {
	mountCmd.Execute()
}

func TestUnmountCmd(t *testing.T) {
	unmountCmd.Execute()
}

func TestPathCmd(t *testing.T) {
	pathCmd.Execute()
}

func TestVersionCmd(t *testing.T) {
	versionCmd.Execute()
}
