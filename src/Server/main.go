package main

import (
	"fmt"

	serverAl "./serverModule/al"
	serverBll "./serverModule/bll"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	dispatcher := serverBll.NewSeverBllFactory().CreateDispatcher()
	server, err := serverAl.NewServer(7788, "127.0.0.1", dispatcher)
	if err != nil {
		println(err.Error())
		fmt.Scanln()
		return
	}
	server.Start()
	fmt.Scanln()
}
