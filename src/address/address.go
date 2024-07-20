package address

import (
	"fmt"
	"net/netip"
)

type VersionNumber byte

const (
	maxPortNum = 65535
	minPortNum = 0

	Version4 VersionNumber = iota
	Version6
)

// Address is a struct which better represents the structure of an IP address.
type Address struct {
	IP      string
	Port    int
	Version VersionNumber
}

// CreateNew creates a new Address. The IP and port are checked to be valid.
func CreateNew(ip string, port int) (Address, error) {

	if port < minPortNum || port > maxPortNum {
		return Address{}, fmt.Errorf("port number is out of range: %d", port)
	}

	addr, err := netip.ParseAddr(ip)

	if err != nil {
		return Address{}, err
	}

	toRet := Address{IP: ip, Port: port}

	if addr.Is4() {
		toRet.Version = Version4
	} else {
		toRet.Version = Version6
	}

	return toRet, nil
}

// Is4 checks if the IP is IPv4.
func (a *Address) Is4() bool {
	return a.Version == Version4
}

// Is6 checks if the IP is IPv6.
func (a *Address) Is6() bool {
	return a.Version == Version6
}
