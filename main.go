package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
)

type Ports struct {
	Ports []Data `json:"ports"`
}

type Data struct {
	Number      int    `json:"number"`
	Tcp         bool   `json:"tcp"`
	Udp         bool   `json:"udp"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func scan(ip string, p int) bool {
	addr := ip + ":" + strconv.Itoa(p)
	_, err_scan := net.Dial("tcp", addr)

	if err_scan == nil {
		return true
	}
	return false
}

func main() {

	for i := 0; i < 36000; i++ {
		resp := scan("localhost", i)

		if resp {

			content, err := ioutil.ReadFile("ppt.json")
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}

			var ports Ports
			json.Unmarshal(content, &ports)

			for u := 0; u < len(ports.Ports); u++ {

				if ports.Ports[u].Number == i {

					fmt.Println("------- " + strconv.Itoa(i) + " -------")
					fmt.Println("[nom]\t", ports.Ports[u].Name)
					fmt.Println("[description]\t", ports.Ports[u].Description)
					fmt.Println("-------------------\n")

					break
				}
			}
		}

	}

}
