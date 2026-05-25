package scanner

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)





type BannerResult struct {
	Port 		int			`json:"port"`
	Banner 		string		`json:"banner,omitempty"`
	Service 	string		`json:"service"`
	Version 	string		`json:"version,omitempty"`
	OS			string		`json:"os,omitempty"`
}

func GrabBanner(host string, port int, timeout time.Duration) (BannerResult, error) {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	conn, err := net.DialTimeout("tcp", address ,timeout)
	if err != nil {
		return BannerResult{Port: port}, fmt.Errorf("connection failed to %s: %v", address, err)
	}
	defer conn.Close()
	err = conn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		return BannerResult{Port: port}, fmt.Errorf("error :%v", err)
	}
	banner := ""
	if port == 80 || port == 443 || port == 8080 {
		request := fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n", host)
		_, err = conn.Write([]byte(request))
		if err != nil {
			return BannerResult{Port: port}, fmt.Errorf("failed to send HTTP request: %v", err)
		}
	}
	reader := bufio.NewReader(conn)
	var bannerLines []string
	for i := 0; i < 10; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		bannerLines = append(bannerLines, line)

		if strings.HasPrefix(strings.ToLower(line), "server:") {
			break
		}
	}
	banner = strings.Join(bannerLines, "\n")
	result := BannerResult{Port: port, Banner: banner}
	return ParseBanner(result), nil
}
//Took top 6 most commonly encountered services, might add on to it later
func ParseBanner(result BannerResult) BannerResult {
	version := ""
	os := ""
	banner := result.Banner
	switch {
	case strings.Contains(banner, "SSH"): 			//SSH
		parts := strings.SplitN(banner, "-", 3)
		if len(parts) < 3 {
			break
		}
		fields := strings.Fields(parts[2])
		if len(fields) < 1 {
			break
		}
		if len(fields) < 2 {
			version = fields[0]
			break
		}
		version = fields[0]
		os = fields[1]
		result.Service = "SSH"
		result.Version = version
		result.OS = os
	case strings.HasPrefix(banner, "+OK"): 			//POP3
		fields := strings.Fields(banner)
		result.Service = fields[1]
	case strings.HasPrefix(banner, "* OK"): 		//IMAP
		fields := strings.Fields(banner)
		result.Service = fields[2]
	case strings.Contains(banner, "ESMTP"): 		//SMTP
		fields := strings.Fields(banner)
		result.Service = fields[3]
	case strings.HasPrefix(banner, "220"):  		//FTP
		fields := strings.Fields(banner)
		result.Service = fields[2]
	case strings.Contains(banner, "Server:"):		//HTTP
		for _, line := range strings.Split(banner, "\n") {
			if strings.Contains(line, "Server:"){
				parts := strings.SplitN(line, ":", 2)
				trimmed := strings.TrimSpace(parts[1]) //Parts has a space after `Server: `
				split := strings.SplitN(trimmed, "/", 2) // SimpleHTTP/0.6 Python/3.12.3 -> SimpleHTTP and 0.6 Python/3.12.3
				result.Service = split[0]
				if len(split) > 1 && len(strings.Fields(split[1])) > 0 {
					result.Version = strings.Fields(split[1])[0]
				}
			}
		}
	default:
		result.Service = banner
		result.Version = ""
}
return result
}