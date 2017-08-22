package main

import (
	"flag"
	"ipfs-client/listener"
	"ipfs-client/base"
	"fmt"
)

var (
	testcase = flag.String("testcase", "hi", "chose test case to execute")
)

func main() {
	flag.Parse()

	switch *testcase {
	case "ipfs-sub":
		listener.Listen()
	}

}
