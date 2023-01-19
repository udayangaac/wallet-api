package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/udayangaac/wallet-api/api/restapi"
	"gopkg.in/alecthomas/kingpin.v2"
)

var port = kingpin.Flag("port", "Http server port.").Short('p').Default("8081").Int()

func main() {

	kingpin.Parse()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	err := restapi.ConfigureAPI(*port)
	if err != nil {
		log.Printf("Unable to start Wallet API: %s", err)
		os.Exit(2)
	}

	<-osSignal
}
