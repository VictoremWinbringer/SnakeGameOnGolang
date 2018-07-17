package main

import (
	"fmt"

	serverAl "./serverModule/al"
	serverBll "./serverModule/bll"
	serverDal "./serverModule/dal"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	ip := ""
	port := 0
	println("Inter ip to liscen")
	_, e := fmt.Scanln(&ip)
	if e != nil {
		println(e.Error())
		fmt.Scanln()
		return
	}
	println("Inter port to liscen")
	_, e = fmt.Scanln(&port)
	if e != nil {
		println(e.Error())
		fmt.Scanln()
		return
	}
	server, err := serverAl.NewServer(port, ip, serverBll.NewSeverBllFactory(serverDal.NewServerDalFactory()))
	if err != nil {
		println(err.Error())
		fmt.Scanln()
		return
	}
	server.Start()
	println("Press Enter to exit")
	fmt.Scanln()
}
