package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func scanPort(host string, port int, protocol string) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	if protocol == "tcp" {
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			return false
		}
		conn.Close()
		return true
	} else {
		conn, err := net.DialTimeout("udp", address, 1*time.Second)
		if err != nil {
			return false
		}
		conn.Close()
		return true
	}
}

func getServiceName(port int, protocol string) string {
	service, err := net.LookupPort(protocol, strconv.Itoa(port))
	if err != nil {
		return "unknown service"
	}
	return strconv.Itoa(service)
}

func main() {
	tcpScan := flag.Bool("t", false, "Perform TCP scan")
	udpScan := flag.Bool("u", false, "Perform UDP scan")

	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: tinyscanner [OPTIONS] [HOST] [PORT]")
		fmt.Println("Options:")
		fmt.Println("  -u               UDP scan")
		fmt.Println("  -t               TCP scan")
		fmt.Println("  --help           Show this message and exit.")
		return
	}

	host := flag.Arg(0)
	portRange := flag.Arg(1)
	fmt.Println(portRange)
	var protocol string

	if *udpScan {
		protocol = "udp"
	} else if *tcpScan {
		protocol = "tcp"
	}

	var ports []int
	if strings.Contains(portRange, "-") {
		portParts := strings.Split(portRange, "-")
		startPort, _ := strconv.Atoi(portParts[0])
		endPort, _ := strconv.Atoi(portParts[1])
		for i := startPort; i <= endPort; i++ {
			ports = append(ports, i)
		}
	} else {
		port, _ := strconv.Atoi(portRange)
		ports = append(ports, port)
	}

	for _, port := range ports {
		isOpen := scanPort(host, port, protocol)
		serviceName := getServiceName(port, protocol)
		status := "closed"
		if isOpen {
			status = "open"
		}
		fmt.Printf("Port %d (%s) is %s\n", port, serviceName, status)
	}
}
