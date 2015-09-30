package main

type fakeVolume struct {
	name string
}

func (vol fakeVolume) Name() string           { return vol.name }
func (vol fakeVolume) DriverName() string     { return "fake" }
func (vol fakeVolume) Path() string           { return "path" }
func (vol fakeVolume) Mount() (string, error) { return "mount", nil }
func (vol fakeVolume) Unmount() error         { return nil }
