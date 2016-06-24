package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MendelGusmao/go-hddtemp"
)

var address = flag.String("address", "localhost:7634", "address where the hddtemp daemon is listening")

func main() {
	disks, err := hddtemp.Fetch(*address)

	if err != nil {
		fmt.Println("main:", err)
		os.Exit(1)
	}

	for _, disk := range disks {
		fmt.Println("DeviceName: ", disk.DeviceName)
		fmt.Println("Model:      ", disk.Model)

		if disk.Status == "" {
			fmt.Printf("Temperature: %d°%s\n", disk.Temperature, disk.Unit)
		} else {
			fmt.Println("Status:     ", disk.Status)
		}
	}
}
