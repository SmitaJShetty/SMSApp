package main

import (
	"SMSApp/pkg/router"
)

func main() {
	listenAddress := "localhost:8090"
	router.Start(listenAddress)
	select {}
}
