package main

import (
	"encoding/json"
	"fmt"
	czmq "github.com/zeromq/goczmq"
	"log"
)

type Vehicle struct {
	Action string `json:"action"`
	Value  int16  `json:"value"`
}

//https://github.com/zeromq/goczmq
/**


需要安装配置
libsodium
libzmq
czmq
*/

func main() {
	endpoint := "tcp://127.0.0.1:5556"
	repSock, err := czmq.NewRep(endpoint)
	if err != nil {
		panic(err)
	}

	defer repSock.Destroy()
	p := &Vehicle{}
	for {

		msg, _, err := repSock.RecvFrame()
		//fmt.Println(msg)

		json.Unmarshal(msg, p)

		//fmt.Println(msg)
		fmt.Println(p)
		if err != nil {
			log.Fatalf("Failed to receive message: %s", err)
		}





		err = repSock.SendFrame(msg, 0)
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}
	}
}

////**
//
//// Send messages and read replies.
//for i := 0; i != *roundtripCount; i++ {
//err := reqSock.SendMessage(msg)
//if err != nil {
//log.Fatalf("Failed to send message: %s", err)
//}
//
//reply, err := reqSock.RecvMessage()
//if err != nil {
//log.Fatalf("Failed to receive message: %s", err)
//}
//
//if len(reply) != 1 {
//log.Fatalf("Message of incorrect size received: %d", len(reply))
//}
//
//if len(reply[0]) != *messageSize {
//log.Fatalf("Message of incorrect size received: %d", len(reply[0]))
//}
//}
