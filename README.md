# cidrtool

This project is a sample project that houses a few functions that make working with CIDR notation a little easier.

## How to use

Install the command line tool using go's tool chain:

```bash
# optionally add go's bin path to your $PATH
export $PATH=$GOPATH/bin

# install cidrc
go install github.com/chuckha/cidrtool/cmd/cidrc

$ cidrc 192.168.1.1/16
Input: 192.168.1.1/16
Lowest IP: 192.168.0.0
Highest IP: 192.168.255.255
```
