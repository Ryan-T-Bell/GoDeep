package trigger

import (
	"godeep/server/utils"
	"net"
	"strconv"
	"strings"
)

func isValidPort(port int) bool {
	return port > 0 && port < 65536
}

// Craft raw packet
func sendPortKnock(ip *net.IPAddr, port int) {
	// TODO
}

// Send port knock sequence
func sendPortKnockSequence(ip *net.IPAddr, ports []int) {
	for _, port := range ports {
		if isValidPort(port) {
			sendPortKnock(ip, port)
		}
	}
}

func craftTCP_Packet(ip string, port int) {

}

// parse input to pull out IP / ports.
// TODO: Refactor to use db-saved triggers
func parseTriggerInput(input string) (*net.IPAddr, []int) {
	words := strings.Split(input, " ")
	strPorts := strings.Split(words[2], ",")

	ip, err := net.ResolveIPAddr("ip", words[1])
	utils.PrintError("Trigger IP input error: \n", err)

	var ports []int

	for _, port := range strPorts {
		portInt, err := strconv.Atoi(port)
		ports = append(ports, portInt)
		utils.PrintError("Trigger port input error: \n", err)
	}

	return ip, ports
}

// Exported function to handle trigger input
func HandleTrigger(input string) {
	ip, ports := parseTriggerInput(input)
	sendPortKnockSequence(ip, ports)
}
