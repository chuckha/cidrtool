package cidrtool_test

import (
	"testing"

	"github.com/chuckha/cidrtool"
)

func TestEverything(t *testing.T) {
	testcases := []struct {
		name          string
		inputIP       string
		expectedIP    int
		cidr          int
		lower         int
		upper         int
		lowerReadable string
		upperReadable string
	}{
		{
			name:          "exact match",
			inputIP:       "19.83.113.32",
			expectedIP:    324235552,
			cidr:          32,
			lower:         324235552,
			upper:         324235552,
			lowerReadable: "19.83.113.32",
			upperReadable: "19.83.113.32",
		},
		{
			name:          "full wildcard",
			inputIP:       "13.246.33.122",
			expectedIP:    234234234,
			cidr:          0,
			lower:         0,
			upper:         (1 << 32) - 1,
			lowerReadable: "0.0.0.0",
			upperReadable: "255.255.255.255",
		},
		{
			name:          "something in the middle",
			inputIP:       "71.183.107.48",
			expectedIP:    1203202864,
			cidr:          16,
			lower:         1203175424,
			upper:         1203240959,
			lowerReadable: "71.183.0.0",
			upperReadable: "71.183.255.255",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ip, err := cidrtool.IPToInt(tc.inputIP)
			if err != nil {
				t.Fatalf("could not convert ip to int: %v", err)
			}
			if ip != tc.expectedIP {
				t.Errorf("[ip to int] expected %v got %v", tc.expectedIP, ip)
			}
			if tc.lower != cidrtool.Lower(ip, tc.cidr) {
				t.Errorf("[lower] expected %v got %v", tc.lower, cidrtool.Lower(ip, tc.cidr))
			}
			if tc.upper != cidrtool.Upper(ip, tc.cidr) {
				t.Errorf("[upper] expected %v got %v", tc.upper, cidrtool.Upper(ip, tc.cidr))
			}
			if cidrtool.IPToString(cidrtool.Upper(ip, tc.cidr)) != tc.upperReadable {
				t.Errorf("[toString] expected %v got %v", tc.upperReadable, cidrtool.IPToString(cidrtool.Upper(ip, tc.cidr)))
			}
		})
	}
}

func TestErrorCase(t *testing.T) {
	_, err := cidrtool.IPToInt("hello")
	if err == nil {
		t.Fatal("expected a failure but did not get one")
	}
}
