package trigger

import (
	"fmt"
	"godeep/server/utils"
	"net"
	"strconv"
	"strings"
)

// Valididators
func isValidIP(ip string) bool {
	_, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return false
	}
	return true
}

func isValidPort(port int) bool {
	return port > 0 && port < 65536
}

// Craft raw packet
func sendPortKnock(ip string, port int) {
	conn, err := net.Dial("tcp", ip+":"+string(port))
	if err != nil {
		return
	}
	defer conn.Close()
}

// Send port knock sequence
func sendPortKnockSequence(ip string, ports []int) {
	if isValidIP(ip) {
		for _, port := range ports {
			if isValidPort(port) {
				sendPortKnock(ip, port)
			}
		}
	}
}

// parse input to pull out IP / ports.
// Refactor to use db-saved triggers
func parseTriggerInput(input string) (string, []int) {
	words := strings.Split(input, " ")
	strPorts := strings.Split(words[2], ",")

	ip := words[1]
	var ports []int

	for _, port := range strPorts {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			fmt.Printf("%sTrigger input processing error: \n %s", utils.WARN, err)
		}

		ports = append(ports, portInt)
	}
	return ip, ports
}

// Exported function to handle trigger input
func HandleTrigger(input string) {
	ip, ports := parseTriggerInput(input)
	sendPortKnockSequence(ip, ports)
}
