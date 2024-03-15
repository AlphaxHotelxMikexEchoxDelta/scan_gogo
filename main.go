package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
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
	fmt.Print("(_________________)\n")

}

func scan(ip string) []string {

	var opens []string

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
				opens = append(opens, "   -- "+ColorBlue+strconv.Itoa(ports.Ports[j].Number)+ColorReset+":"+ColorReset+ColorBlue+ports.Ports[j].Name+ColorReset+" --   "+colorGrey+ports.Ports[j].Description+ColorReset+"\n")
				conn.Close()
			}

		}(i)
	}
	wg.Wait()

	return opens

}

func run(addr_reseaux string, min int, max int) map[string][]string {

	var wg sync.WaitGroup
	var open []string
	result := make(map[string][]string)

	for i := min; i < max+1; i++ {

		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			open = scan(addr_reseaux + "." + strconv.Itoa(j))

			if len(open) != 0 {
				result[ColorReset+"\n  "+addr_reseaux+"."+strconv.Itoa(j)+"\t\t\t\t["+ColorGreen+"OK"+ColorReset+"]\n"] = open
			}

		}(i)
	}
	wg.Wait()

	return result
}

func main() {

	if len(os.Args[1:]) != 0 {

		banner(os.Args[1])

		plageIP := strings.Split(os.Args[1], ".")
		limitesPlage := strings.Split(plageIP[len(plageIP)-1], "-")

		ip := strings.Join(plageIP[:len(plageIP)-1], ".")

		if len(limitesPlage) != 1 {
			min, _ := strconv.Atoi(limitesPlage[0])
			max, _ := strconv.Atoi(limitesPlage[1])
			resp := run(ip, min, max)

			for key, ip := range resp {
				fmt.Print(key, ip)
			}

		} else {
			min, _ := strconv.Atoi(limitesPlage[0])
			max, _ := strconv.Atoi(limitesPlage[0])
			resp := run(ip, min, max)

			for key, ip := range resp {
				fmt.Print(key, ip)
			}

		}

	} else {

		fmt.Print(ColorRed + "\nSaisi invalide !" + ColorReset + "\n[ main.exe @ip_cible ]\n\n")

	}

}
