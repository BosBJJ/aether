package scanner

import (
	"testing"
	"time"
)

func TestICMPPing(t *testing.T) {
    result := icmpPing("scanme.nmap.org", 5*time.Second)
    t.Logf("IP: %s", result.IP)
    t.Logf("Alive: %v", result.Alive)
    t.Logf("Latency: %v", result.Latency)
    if result.Error != nil {
        t.Logf("Error: %v", result.Error)
    }
}

func TestConvertCIDR(t *testing.T) {
    hosts, err := convertCIDR("172.27.168.240/28")
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("Host count: %d", len(hosts))
    for _, h := range hosts {
        t.Logf("  %s", h)
    }
}

func TestCIDR(t *testing.T) {
    results, err := SweepCIDR("172.27.168.240/28", 2 * time.Second, 50)
    if err != nil {
    t.Fatal(err)
}
for _, r := range results {
    if r.Alive {
        t.Logf("ALIVE  %s  latency: %.3fms", r.IP, r.Latency.Seconds()*1000)
    } else {
        t.Logf("DEAD   %s", r.IP)
    }
}
}

func TestPingLocal(t *testing.T) {
    result := icmpPing("172.27.168.240", 3*time.Second)
    t.Logf("IP: %s", result.IP)
    t.Logf("Alive: %v", result.Alive)
    t.Logf("Latency: %v", result.Latency)
    t.Logf("Error: %v", result.Error)
}