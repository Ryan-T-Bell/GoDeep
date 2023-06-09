package trigger

import (
	"net"
	"testing"
)

func TestParseTriggerInput(t *testing.T) {
	testCases := []struct {
		input         string
		expectedIP    *net.IPAddr
		expectedPorts []int
	}{
		{
			input: "192.168.1.1 80,443",
			expectedIP: &net.IPAddr{
				IP: net.ParseIP("192.168.1.1"),
			},
			expectedPorts: []int{80, 443},
		},
		{
			input: "127.0.0.1 80",
			expectedIP: &net.IPAddr{
				IP: net.ParseIP("127.0.0.1"),
			},
			expectedPorts: []int{80},
		},
	}

	for _, testCase := range testCases {
		ip, ports := parseTriggerInput(testCase.input)

		if ip.String() != testCase.expectedIP.String() {
			t.Errorf("Expected IP %s, got %s", testCase.expectedIP, ip)
		}

		if len(ports) != len(testCase.expectedPorts) {
			t.Errorf("Expected %d ports, got %d", len(testCase.expectedPorts), len(ports))
		}

		for i, port := range ports {
			if port != testCase.expectedPorts[i] {
				t.Errorf("Expected port %d, got %d", testCase.expectedPorts[i], port)
			}
		}
	}
}
