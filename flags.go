package main

import "github.com/spf13/cobra"

func initFlags() {
	cobra.HelpFlagShorthand = "?"
	cobra.HelpFlagUsageFormatString = "Help for %s"

	initGlobalFlags()

}

var volumeDriver string
var volumeName string

func initGlobalFlags() {
	DvdcliCmd.PersistentFlags().StringVarP(&volumeDriver, "volumedriver", "", "",
		"Volume driver")
	DvdcliCmd.PersistentFlags().StringVarP(&volumeName, "volumename", "", "",
		"Volume name")
}
