// A package for exploring various ways to work with ipv4 ip addresses.
// All references to IP is implied ipv4.
package cidrtool

import (
	"fmt"
	"strconv"
	"strings"
)

// IPToInt converts an IP address (x.x.x.x) to a decimal based integer.
func IPToInt(ip string) (int, error) {
	octets := strings.Split(ip, ".")
	sum := 0
	for i := 0; i < 4; i++ {
		val, err := strconv.Atoi(octets[i])
		if err != nil {
			return 0, err
		}
		sum += val << uint(24-8*i)
	}
	return sum, nil
}

// Upper returns the highest address in an ipv4 block.
func Upper(ip, cidr int) int {
	// low range | wildcard is the highest
	// how to get wild card?
	// 32 1s &^ with cidr == wildcard
	high := ((1 << 32) - 1) &^ (((2 << uint(cidr)) - 1) << uint(32-cidr))
	return Lower(ip, cidr) | high
}

// Lower finds the lowest address in an ipv4 block.
func Lower(ip, cidr int) int {
	// the minus one is to flip 10 to 1 or 100 to 11 to create a mask
	// the second left shift is to put the mask in the right place
	// example:
	// cidr of 8: 2 << 8            == 100000000
	//           (2 << 8) - 1       == 11111111
	//           (2 << 8) - 1 << 24 == 11111111000000000000000000000000
	return ip & (((2 << uint(cidr)) - 1) << uint(32-cidr))
}

// IPToString converts an integer to a regular looking IP address.
func IPToString(ip int) string {
	msk := 255
	out := make([]string, 4)
	for i := range out {
		out[i] = fmt.Sprintf("%d", ip>>uint(24-8*i)&msk)
	}
	return strings.Join(out, ".")
}
