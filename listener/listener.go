package listener


import (
	"github.com/ipfs/go-ipfs-api"
	"strconv"
)

func Listen() {
	sh := shell.NewLocalShell()
	scribe, _:= sh.PubSubSubscribe("test_topic")
	for {
		record,_:=scribe.Next()
		peerId := record.From().String()
		topics := record.TopicIDs()
		seqNum := record.SeqNo()
		data := string(record.Data())


		println("peer id is " + peerId)
		println("topics are " + topics[0])
		println("sequence num is " + strconv.FormatInt(seqNum, 10))
		println("data is " + data)
		println("===========================================")
	}
}
