package main

import (
	"flag"
	"log"
)

func parseArgs() (string, int) {
	// not sure if the following has sense...
	// out there should be libs that allow to define short and long flags

	// I'm not going to test this function

	hostPtrLong := flag.String("host", "localhost", "address \n")
	hostPtr := flag.String("h", "", "-host")
	portPtrLong := flag.Int("port", 9090, "port to listen\n")
	portPtr := flag.Int("p", 0, "-port")

	flag.Parse()

	hostShort := *hostPtr
	hostLong := *hostPtrLong
	portShort := *portPtr
	portLong := *portPtrLong

	var host string
	var port int

	if hostShort == "" {
		host = hostLong
	} else {
		host = hostShort
		if hostLong != "localhost" {
			log.Panicln("--host and -h defined. Use only one value")
		}
	}

	if portShort == 0 {
		port = portLong
	} else {
		port = portShort
		if portLong != 9090 {
			log.Panicln("--port and -p defined. Use only one value")
		}
	}

	return host, port
}
