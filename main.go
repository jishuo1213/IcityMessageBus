package main

import (
	"IcityMessageBus/requester"
	"IcityMessageBus/server"
	config2 "IcityMessageBus/config"
	"log"
)

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
	if err != nil {
		log.Print("parse config file failed")
		return
	}

	requester.Start()
	server.Start("127.0.0.1", 1214)
}
