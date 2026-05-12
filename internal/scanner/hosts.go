package scanner

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type PingResult struct {
	IP 		string
	Alive 	bool
	Latency time.Duration
	Error 	error
}



func SweepCIDR(cidr string, timeout time.Duration, workers int) ([]PingResult, error) {
	hosts, err := convertCIDR(cidr)
	if err != nil {
        return nil, err
    }
	return PingHosts(hosts, timeout, workers)
}

func convertCIDR(cidr string) ([]string, error) {
	var hosts []string
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse cidr: %v", err)
	}
	for ip := network.IP.Mask(network.Mask); network.Contains(ip); incIP(ip){ //network is 192.168.2.X, inc breaks at 192.168.3
		hosts = append(hosts, ip.String())
	} 
	return hosts, nil
}

func incIP(ip net.IP) { //192.168.1.0 would have ip[i] be 3(0) goes to 255, makes 192.168.2.0
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}


func PingHosts(hosts []string, timeout time.Duration, workers int) ([]PingResult, error) {  
	if len(hosts) < 1 {
		return nil, fmt.Errorf("no hosts entered")
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []PingResult
	slots := make(chan struct{}, workers)

	for _, host := range hosts {
		wg.Add(1)
		slots <- struct{}{}

		go func (host string)  {
			defer wg.Done()
			defer func ()  { <-slots }()
			result := pingHost(host, timeout)
			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(host)
	}

	wg.Wait()
	return results, nil
}

func pingHost(host string, timeout time.Duration) PingResult {
	result := icmpPing(host, timeout)
    if result.Alive {
        return result
    }
	for _, port := range []int{80, 443, 22, 21, 25} {  //fallback to most common ports
        result = tcpPing(host, port, timeout)
        if result.Alive {
            return result
        }
    }
	return result
}

func icmpPing(host string, timeout time.Duration) PingResult {
	result := PingResult{}
	addr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return PingResult{IP: host, Error: err}
	}
	conn, err := net.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		result.IP = addr.IP.String()
		result.Error = err
		return result
	}
	result.IP = addr.IP.String()
	defer conn.Close()
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, 
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff, 
			Seq: 1,
			Data: []byte("ping"),
		},
	}
	msgByte, err := msg.Marshal(nil)
	if err != nil {
		result.Error = err
		return result
	}
	start := time.Now()
	conn.SetDeadline(time.Now().Add(timeout))
	_, err = conn.WriteTo(msgByte, addr) 
	if err != nil {
		result.Error = err
		return result
	}
	respB := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(respB)
		if err != nil {
			result.Error = err
			return result
		}
		resp, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), respB[:n])
		if err != nil {
			result.Error = err
			return result
		}
		if resp.Type == ipv4.ICMPTypeEchoReply {
			lat := time.Since(start)
			result.Latency = lat
			result.Alive = true
			return  result
		}
	}
	
	return result
}

func tcpPing(host string, port int, timeout time.Duration) PingResult {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address ,timeout)
	if err != nil {
		return PingResult{
			IP: host,
			Alive: false,
			Error: err,
		}
	}
	defer conn.Close()
	return PingResult{
		IP: host,
		Alive: true,
		Latency: time.Since(start),
	}
}