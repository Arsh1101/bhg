package main

import (
	"fmt"
	"net"
	"sort"
)

// TCP Range: 1 to 65535
const starPort, endPort, channelSize int = 1, 1024, 100

// Alternative test destination > "127.0.0.1"
const destination, protocol string = "scanme.nmap.org", "tcp"

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf(destination+":%d", p)
		conn, err := net.Dial(protocol, address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, channelSize)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := starPort; i <= endPort; i++ {
			ports <- i
		}
	}()

	for i := 0; i < endPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
