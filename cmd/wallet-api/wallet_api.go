package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/udayangaac/wallet-api/api/restapi"
	"gopkg.in/alecthomas/kingpin.v2"
)

var port = kingpin.Flag("port", "Http server port.").Short('p').Default("8081").Int()

func main() {

	os.Setenv("TZ", "UTC")
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	restapi.ConfigureAPI(*port)

	<-osSignal
}
