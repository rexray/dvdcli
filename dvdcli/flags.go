package main

func initFlags() {
	initGlobalFlags()

}

var volumeDriver string
var volumeName string
var volumeOpts []string
var explicitCreate bool

func initGlobalFlags() {
	DvdcliCmd.PersistentFlags().StringVarP(&volumeDriver, "volumedriver", "", "",
		"Volume driver")
	DvdcliCmd.PersistentFlags().StringVarP(&volumeName, "volumename", "", "",
		"Volume name")
	createCmd.Flags().StringSliceVarP(&volumeOpts, "volumeopts", "", []string{},
		"Volume options")
	mountCmd.Flags().StringSliceVarP(&volumeOpts, "volumeopts", "", []string{},
		"Volume options")
	mountCmd.Flags().BoolVarP(&explicitCreate, "explicitcreate", "", false,
		"Explicit create flag to disable creating volumes that don't exist")
}
