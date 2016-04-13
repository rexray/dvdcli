package main

func initFlags() {
	initGlobalFlags()

}

var volumeDriver string
var volumeName string
var volumeOpts []string
var checkVolumeExist bool

func initGlobalFlags() {
	DvdcliCmd.PersistentFlags().StringVarP(&volumeDriver, "volumedriver", "", "",
		"Volume driver")
	DvdcliCmd.PersistentFlags().StringVarP(&volumeName, "volumename", "", "",
		"Volume name")
	createCmd.Flags().StringSliceVarP(&volumeOpts, "volumeopts", "", []string{},
		"Volume options")
	mountCmd.Flags().StringSliceVarP(&volumeOpts, "volumeopts", "", []string{},
		"Volume options")
	mountCmd.Flags().BoolVarP("checkVolumeExist", "c", false,
		"Check Volume exist or not")
	unmountCmd.Flags().BoolVarP("checkVolumeExist", "c", false,
		"Check Volume exist or not")
	pathCmd.Flags().BoolVarP("checkVolumeExist", "c", false,
		"Check Volume exist or not")
}
