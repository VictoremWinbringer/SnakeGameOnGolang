package main

import (
	"fmt"
    _ "net/http/pprof"
	serverAl "./serverModule/al"
	serverBll "./serverModule/bll"
	serverDal "./serverModule/dal"
	"flag"
	"os"
	"log"
	"runtime/pprof"
	"runtime"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

