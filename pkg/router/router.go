package router

import (
	"fmt"
	"net/http"
)

func Start(listenAddress string) {
	router := GetRouter()
	addRoutes(router)
	go func() {
		err := http.ListenAndServe(listenAddress, router)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Listening on port:", listenAddress)
	}()
}
