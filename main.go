package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

type Ports struct {
	Ports []Data `json:"ports"`
}

type Data struct {
	Number int `json:"number"`
	//Tcp         bool   `json:"tcp"`
	//Udp         bool   `json:"udp"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func banner(ip string) {

	fmt.Println("     _________")
	fmt.Println("    / ======= \\")
	fmt.Println("   / __________\\")
	fmt.Println("  | ___________ |")
	fmt.Println("  | | -       | |    @ip_target: " + ip)
	fmt.Println("  | | scan_go | |    by @BigBurgerBoy")
	fmt.Println("  | |_________| |")
	fmt.Println("  \\=____________/")
	fmt.Println("  / ''''''''''''\\")
	fmt.Println(" / ::::::::::::: \\")
	fmt.Println("(_________________)\n")

}

func scan(ip string, p int) bool {
	addr := ip + ":" + strconv.Itoa(p)
	_, err_scan := net.Dial("tcp", addr)

	if err_scan == nil {
		return true
	}
	return false
}

func scanner(ip string) {

	content, _ := ioutil.ReadFile("ppt.json")
	var ports Ports
	json.Unmarshal(content, &ports)

	for i := range ports.Ports {

		resp := scan(ip, ports.Ports[i].Number)

		if resp {

			fmt.Println("\n[ " + strconv.Itoa(ports.Ports[i].Number) + " ]\t" + ports.Ports[i].Name)
			fmt.Println("--> ", ports.Ports[i].Description)

			/*
				if ports.Ports[i].Tcp {
					fmt.Print(" [  TCP  ]")
				}
				if ports.Ports[i].Udp {
					fmt.Print(" [  UDP  ]")
				}
			*/

		}

	}
}

func main() {

	if len(os.Args) != 1 {

		ip := os.Args[1]

		banner(ip)
		scanner(ip)

	} else {

		fmt.Println("Veuillez indiquez une ip")

	}

}
