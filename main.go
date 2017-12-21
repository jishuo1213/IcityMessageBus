package main

import (
	config2 "IcityMessageBus/config"
	"IcityMessageBus/cmsp"
	"log"
	"IcityMessageBus/requester"
	"IcityMessageBus/server"
	"runtime"
)

func init() {
	err := cmsp.GetQueue(config2.REQUEST_QUEUE_NAME)
	if err != nil {
		panic(err)
	}
}

func main() {
	//err := cmsp.PutMessageIntoQueue("test", []byte("aaaaaaaaaaaaa"))
	//if err == nil {
	//	fmt.Println("put success")
	//} else {
	//	fmt.Println(err)
	//}
	//msg, err := cmsp.ReadOneMessageFromQueue("test")
	//if err == nil {
	//	fmt.Println(string(msg))
	//}
	//config.REQUESTER_NUM = 8

	err := config2.InitConfig()

	log.Print("go procs:", config2.Config.MaxProcs)
	runtime.GOMAXPROCS(config2.Config.MaxProcs)

	if err != nil {
		log.Print("parse config file failed")
		return
	}

	requester.Start()
	server.Start("0.0.0.0", 1214)

	//fmt.Println(utils.DigestMessage([]byte("asdasdafsadf")))
}
