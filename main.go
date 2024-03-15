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

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	colorGrey   = "\033[90m"
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

	fmt.Println("       _________")
	fmt.Println("      / ======= \\")
	fmt.Println("     / __________\\")
	fmt.Println("    | ___________ |")
	fmt.Println("    | | -       | |    @ip_target: " + ColorRed + ip + ColorReset)
	fmt.Println("    | | " + ColorRed + "scan_go" + ColorReset + " | |    by @BigBurgerBoy")
	fmt.Println("    | |_________| |")
	fmt.Println("    \\=____________/")
	fmt.Println("    / ''''''''''''\\")
	fmt.Println("   / ::::::::::::: \\")
	fmt.Print("  (_________________)\n")

}

func scan(ip string) []string {

	var opens []string

	content, _ := ioutil.ReadFile("ppt.json")
	var ports Ports
	json.Unmarshal(content, &ports)

	var wg sync.WaitGroup

	for _, port := range ports.Ports {

		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			conn, err_scan := net.Dial("tcp", ip+":"+strconv.Itoa(port.Number))

			if err_scan == nil {
				opens = append(opens, "   -- "+ColorBlue+strconv.Itoa(port.Number)+ColorReset+":"+ColorReset+ColorBlue+port.Name+ColorReset+" --   "+colorGrey+port.Description+ColorReset+"\n")
				conn.Close()
			}

		}(port.Number)
	}
	wg.Wait()

	return opens

}

func run(ip string) map[string][]string {

	var open []string
	result := make(map[string][]string)

	open = scan(ip)

	if len(open) != 0 {
		result[ColorReset+"\n  "+ip+"\t\t\t\t["+ColorGreen+"UP"+ColorReset+"]\n"] = open[:]
	}

	return result
}

func main() {

	if len(os.Args[1:]) != 0 {

		banner(os.Args[1])

		resp := run(os.Args[1])

		for ip, ports := range resp {
			fmt.Print(ip)
			for port := range ports {
				fmt.Print(ports[port])
			}
		}

	} else {

		fmt.Print(ColorRed + "\nSaisi invalide !" + ColorReset + "\n[ main.exe @ip_cible ]\n\n")

	}

}
