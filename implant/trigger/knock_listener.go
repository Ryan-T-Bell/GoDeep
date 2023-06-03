package trigger

import (
	"fmt"
	"net"
)

func StartKnockListener() {
	knockSequence := []int{8000, 9000, 7000} // Replace with your desired port knock sequence

	fmt.Println("Listening for port knock sequence...")

	listenAddr := ":8080" // Replace with the port you want to listen on

	ln, err := net.Listen("tcp", listenAddr)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
		ip := remoteAddr.IP.String()
		port := remoteAddr.Port

		fmt.Printf("Received connection from %s:%d\n", ip, port)

		if isPortKnockSequenceValid(knockSequence, port) {
			fmt.Printf("Valid port knock sequence received from %s:%d\n", ip, port)

			// Reset the knock sequence
			knockSequence = []int{8000, 9000, 7000}
		} else {
			fmt.Printf("Invalid port knock sequence received from %s:%d\n", ip, port)
		}

		conn.Close()
	}
}

func isPortKnockSequenceValid(sequence []int, port int) bool {
	if len(sequence) == 0 {
		return false
	}

	if sequence[0] == port {
		// Remove the first port from the sequence
		sequence = sequence[1:]

		if len(sequence) == 0 {
			return true
		}
	}

	return false
}
