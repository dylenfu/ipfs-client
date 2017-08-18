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
	case "struct":
		t := &Tee{3, "hello"}
		println(t.getTeeNum())
		println(t.getTeeName())

	case "channel":
		channeldemo()
	}
}

type Tee struct {
	a  uint
	b  string
}


func (tee *Tee)  getTeeNum() uint {
	return tee.a
}

func (tee Tee) getTeeName() string {
	return tee.b
}

func channeldemo() {
	messages := make(chan string)


	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case msg := <-messages:
		println("received message", msg)
	}

}