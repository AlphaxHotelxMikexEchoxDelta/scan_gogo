package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"sync"
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

func run(ip string) {

	content, _ := ioutil.ReadFile("ppt.json")
	var ports Ports
	json.Unmarshal(content, &ports)

	var wg sync.WaitGroup

	for i := range ports.Ports {

		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			conn, err_scan := net.Dial("tcp", ip+":"+strconv.Itoa(ports.Ports[j].Number))

			if err_scan == nil {
				fmt.Println("\n[ "+strconv.Itoa(ports.Ports[j].Number)+" ] "+ports.Ports[j].Name+"\n--> ", ports.Ports[j].Description)
				conn.Close()
			}

		}(i)
	}
	wg.Wait()

}

func main() {

	if len(os.Args) != 1 {
		ip := os.Args[1]
		banner(ip)
		run(ip)
	} else {
		fmt.Println("\nSaisi invalide !\n[ main.exe @ip_cible ]\n")
	}

}
