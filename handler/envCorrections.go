package handler

import (
	"os"
	"strings"
)

func getAllowBrowsing() bool {
	return os.Getenv("ALLOW_BROWSING") == "true"
}

func getTrustedProxy() bool {
	return os.Getenv("TRUSTED_PROXY") == "true"
}

func getTrustedProxies() []string {
	ips := os.Getenv("TRUSTED_PROXIES")
	// split ips by comma and remove spaces
	ipsArray := strings.Split(ips, ",")
	for i, ip := range ipsArray {
		ipsArray[i] = strings.TrimSpace(ip)
	}
	return ipsArray
}
