package main

import (
	"errors"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/plugins"
	"github.com/docker/docker/volume"
	"github.com/docker/docker/volume/drivers"
)

func init() {
	// updateLogLevel()
	initCommands()
	initFlags()
	// initUsageTemplates()
}

const version = "0.1"

func updateLogLevel() {
	log.SetLevel(log.DebugLevel)
}

func getVolumeDriver(name string) (volume.Driver, error) {
	if name == "" {
		return nil, errors.New("Missing volume driver name")
	}
	return volumedrivers.Lookup(name)
}

var driver volume.Driver
var err error

func initDvd() {
	driver, err = getVolumeDriver(volumeDriver)
	if err != nil {
		log.WithField("volumeDriver", volumeDriver).Info(err)
		os.Exit(1)
	}
}

func initPlugin() {
	volPlugin, err = plugins.Get(volumeDriver, "VolumeDriver")
	volProxy = volumeDriverProxy{volPlugin.Client}
	if err != nil {
		log.WithField("volumeDriver", volumeDriver).Info(err)
	}
}

func main() {
	DvdcliCmd.Execute()
}
