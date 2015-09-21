package main

func initFlags() {
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
