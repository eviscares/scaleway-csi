package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scaleway/scaleway-csi/driver"
	"k8s.io/klog/v2"
)

var (
	endpoint     = flag.String("endpoint", "unix:/tmp/csi.sock", "CSI endpoint")
	prefix       = flag.String("prefix", "", "Prefix to add in block volume name")
	version      = flag.Bool("version", false, "Print the version and exit")
	mode         = flag.String("mode", string(driver.AllMode), "The mode in which the CSI driver will be run (all, node, controller)")
	volumeNumber = flag.Int("volumes", 16, "Number of PVCs per node. Be aware that anything higher than default(16) can lead to decreased performance")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	if *version {
		info := driver.GetVersion()

		fmt.Printf("%+v", info)
		os.Exit(0)
	}

	scwDriver, err := driver.NewDriver(&driver.DriverConfig{
		Endpoint: *endpoint,
		Mode:     driver.Mode(*mode),
		Prefix:   *prefix,
	}, &driver.NodeConfig{
		VolumeNumber: *volumeNumber,
	})
	if err != nil {
		klog.Fatalln(err)
	}

	if err := scwDriver.Run(); err != nil {
		klog.Fatalln(err)
	}
}
