package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/chuckha/cidrtool"
)

func main() {
	flag.Parse()
	input := flag.Arg(0)
	ip, cidr := inputToIPAndCIDR(input)
	ipint, err := cidrtool.IPToInt(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println("Input:", input)
	fmt.Println("Lowest IP:", cidrtool.IPToString(cidrtool.Lower(ipint, cidr)))
	fmt.Println("Highest IP:", cidrtool.IPToString(cidrtool.Upper(ipint, cidr)))
	// TODO number of IPs
}

func inputToIPAndCIDR(input string) (string, int) {
	split := strings.Split(input, "/")
	cidr, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return split[0], cidr
}
