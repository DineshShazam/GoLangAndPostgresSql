package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	allTech := genTechForWeek()
	fmt.Println("BackEnd Technologies")
	backendTech := allTech.listBackend()
	backendTech.writeFile("BackEndStack " + currentTime.Format(`2006-01-02`))
	fmt.Println("Devops Technologies")
	devops := allTech.listDevops()
	devops.writeFile("Devops Stack " + currentTime.Format(`2006-01-02`))

}
