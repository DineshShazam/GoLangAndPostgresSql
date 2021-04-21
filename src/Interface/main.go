package main

import (
	"fmt"
)

type systemDetails struct {
	machineName string
	osType      string
	// machineIp   string
}

// type systemConfiguration struct {

// 	// core        string
// }

// this has to be used for both logging and api log
type machineDetails interface {
	machineInfo() string
	coreInfo() string
}

// if we gonna pass values we use (s sytemDetails)
// we only need the struct type here

func (m systemDetails) machineInfo() string {
	return m.machineName
}

func (c systemDetails) coreInfo() string {
	return c.osType
}

func main() {
	var machineDetails machineDetails
	machineDetails = systemDetails{
		machineName: "Shazam",
		osType:      "Linux",
	}

	fmt.Println("Machine Name: " + machineDetails.machineInfo())
	fmt.Println("Machine Info: " + machineDetails.coreInfo())
}
