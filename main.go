package main

import (
	"fmt"
	"log"
)

func main() {
	ipInfo, err := getIpInfo()
	if err != nil {
		log.Panic("We can't obtain the IP info.")
	}
	fmt.Printf("This is your public IP address: %s\nIt belongs to: %s.\nYou are located at %s, %s", ipInfo.IpAddr, ipInfo.OrgName, ipInfo.RegionName, ipInfo.CountryCode)
}
