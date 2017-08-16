package main

import (
	"flag"
	"ipfs-client/listener"
)

var (
	testcase = flag.String("testcase", "hi", "chose test case to execute")
)


func main() {
	flag.Parse()

	switch *testcase {
	case "hi":
		println("Hello IPFS!")
	case "sub":
		listener.Listen()
	}
}
