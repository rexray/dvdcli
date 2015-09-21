package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func initCommands() {
	DvdcliCmd.AddCommand(createCmd)
	DvdcliCmd.AddCommand(removeCmd)
	DvdcliCmd.AddCommand(mountCmd)
	DvdcliCmd.AddCommand(unmountCmd)
	DvdcliCmd.AddCommand(pathCmd)
}

//DvdcliCmd
var DvdcliCmd = &cobra.Command{
	Use: "dvdcli",
	Short: "dvdcli:\n" +
		"  A Docker Volume Driver CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

//createCmd
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a volume",
	Run: func(cmd *cobra.Command, args []string) {
		initDvd()
		vol, err := driver.Create(volumeName, nil)
		if err != nil {
			log.WithField("volumeName", volumeName).Error(err)
			os.Exit(1)
		}
		log.Info(fmt.Sprintf("%s", vol.Name()))
	},
}

//removeCmd
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a volume",
	Run: func(cmd *cobra.Command, args []string) {
		initDvd()

		vol := fakeVolume{volumeName}
		err := driver.Remove(vol)
		if err != nil {
			log.WithField("volumeName", volumeName).Error(err)
			os.Exit(1)
		}
	},
}

//mountCmd
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount a volume",
	Run: func(cmd *cobra.Command, args []string) {
		initDvd()
		vol, err := driver.Create(volumeName, nil)
		if err != nil {
			log.WithField("volumeName", volumeName).Error(err)
			os.Exit(1)
		}

		path, err := vol.Mount()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Info(fmt.Sprintf("%s", path))
		fmt.Println(path)
	},
}

//unmountCmd
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "Unmount a volume",
	Run: func(cmd *cobra.Command, args []string) {
		initDvd()
		vol, err := driver.Create(volumeName, nil)
		if err != nil {
			log.WithField("volumeName", volumeName).Error(err)
			os.Exit(1)
		}

		if err = vol.Unmount(); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	},
}

//pathCmd
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Path of a volume",
	Run: func(cmd *cobra.Command, args []string) {
		initDvd()
		vol, err := driver.Create(volumeName, nil)
		if err != nil {
			log.WithField("volumeName", volumeName).Error(err)
			os.Exit(1)
		}

		path := vol.Path()
		if path != "" {
			log.Info(fmt.Sprintf("%s", path))
			fmt.Println(path)
		}
	},
}