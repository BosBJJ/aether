package scanner

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

const (
	minPort = 1
	maxPort = 65535
)

type ScanResult struct {
	Port 	int			`json:"port"`
	State 	string		`json:"state"`
}

func ScanPorts(host string, ports []int, timeout time.Duration, workers int) ([]ScanResult, error) {
	if len(ports) == 0 {
		return nil, fmt.Errorf("no ports provided to scan")
	}
	if workers < 1 {
		return nil, fmt.Errorf("workers must be at least 1")
	}
	var results []ScanResult
	var mu sync.Mutex
	var wg sync.WaitGroup

	slots := make(chan struct{}, workers)

	for _, port := range ports {
		if port < minPort || port > maxPort {
			return nil, fmt.Errorf("invalid Port: %d, (must be between %d and %d)", port, minPort, maxPort)
		}

		wg.Add(1)
		slots <- struct{}{}

		go func (port int)  {
			defer wg.Done()
			defer func ()  { <- slots }()
			scan := scanSinglePort(host, port, timeout)
			mu.Lock()
			results = append(results, scan)
			mu.Unlock()
		}(port)
	}
	wg.Wait()
	sortResultsByPort(results)
	return results, nil
}

func scanSinglePort(host string, port int, timeout time.Duration) ScanResult {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	
	conn, err := net.DialTimeout("tcp", address ,timeout)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return ScanResult{Port: port, State: "Filtered"}
		}
		return ScanResult{Port: port, State: "Closed"}
	}
	conn.Close()
	
	return ScanResult{Port: port, State: "Open"}
}

func sortResultsByPort(results []ScanResult) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})
}