package main

import (
	"SMSApp/pkg/router"
	"fmt"
)

func main() {
	listenAddress := "localhost:8090"
	router.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}
