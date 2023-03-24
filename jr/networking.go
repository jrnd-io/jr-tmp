package jr

import (
	"net"
	"time"
)

func unixTimeStamp(days int) int64 {
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	first := now.AddDate(0, 0, -days).Sub(unixEpoch).Seconds()
	last := now.Sub(unixEpoch).Seconds()
	return Random.Int63n(int64(last-first)) + int64(first)
}

func ipKnownPorts() string {
	ports := []string{"80", "81", "443", "22", "631"}
	return ports[Random.Intn(len(ports))]
}

func ipKnownProtocols() string {
	protocols := []string{"TCP", "UDP", "ICMP", "FTP", "HTTP", "SFTP"}
	return protocols[Random.Intn(len(protocols))]
}

func httpMethod() string {
	method := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	return method[Random.Intn(len(method))]
}

func ipv6() string {
	ip := make(net.IP, net.IPv6len)
	for i := 0; i < net.IPv6len; i++ {
		ip[i] = byte(Random.Intn(256))
	}
	ip[0] &= 0xfe // Set the "locally administered" flag
	ip[0] |= 0x02 // Set the "unicast" flag
	return ip.String()
}

func mac() string {
	mac := make(net.HardwareAddr, 6)
	Random.Read(mac)
	mac[0] &= 0xfe // Set the "locally administered" flag
	mac[0] |= 0x02 // Set the "unicast" flag
	return mac.String()
}

func ip(cidr string) string {

GENERATE:

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "0.0.0.0"
	}

	ones, _ := ipnet.Mask.Size()
	quotient := ones / 8
	remainder := ones % 8

	r := make([]byte, 4)
	Random.Read(r)

	for i := 0; i <= quotient; i++ {
		if i == quotient {
			shifted := (r[i]) >> remainder
			r[i] = ^ipnet.IP[i] & shifted
		} else {
			r[i] = ipnet.IP[i]
		}
	}
	ip = net.IPv4(r[0], r[1], r[2], r[3])

	if ip.Equal(ipnet.IP) /*|| ip.Equal(broadcast) */ {
		goto GENERATE
	}
	return ip.String()
}

func password(length int, memorable bool, prefix string, postfix string) string {

	const (
		// Define the set of vowels and consonants that can be used to generate the password.
		vowels     = "aeiouyAEIOUY"
		consonants = "bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ"
	)

	// Generate a memorable password of the specified length.
	password := make([]byte, length)
	if memorable {
		// Generate a memorable password.
		for i := range password {
			if i%2 == 0 {
				// Use a vowel.
				char := vowels[Random.Intn(len(vowels))]
				password[i] = char
			} else {
				// Use a consonant.
				char := consonants[Random.Intn(len(consonants))]
				password[i] = char
			}
		}
	} else {
		// Generate a random password using the full charset.
		charset := vowels + consonants + "0123456789!@#$%^&*()_+{}:\"<>?,./;'[]\\-=`~"
		for i := range password {
			char := charset[Random.Intn(len(charset))]
			password[i] = char
		}
	}

	return prefix + string(password) + postfix
}
