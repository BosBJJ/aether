package scanner

import (
	"fmt"
	"testing"
	"time"
)

func TestGrabBanner(t *testing.T) {
	testCases := []struct {
		name    string
		host    string
		port    int
		timeout time.Duration
	}{
		{"SSH - scanme", "scanme.nmap.org", 22, 5 * time.Second},
		{"HTTP - local", "localhost", 8080, 5 * time.Second},
		{"FTP - debian", "ftp.debian.org", 21, 5 * time.Second},
		{"SMTP - gmail", "smtp.gmail.com", 25, 5 * time.Second},
		{"HTTPS - google", "google.com", 443, 5 * time.Second},
		{"Likely Filtered", "192.168.1.1", 22, 3 * time.Second},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GrabBanner(tc.host, tc.port, tc.timeout)

			fmt.Printf("[%s] Port %d:\n", tc.name, tc.port)

			if err != nil {
				t.Logf("  Error: %v\n", err)
				return
			}
			t.Logf("  Banner  : %s", result.Banner)
			t.Logf("  Service : %s", result.Service)
			t.Logf("  Version : %s", result.Version)
			t.Logf("  OS      : %s", result.OS)
		})
	}
}