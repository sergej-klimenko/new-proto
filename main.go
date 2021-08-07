package main

import (
	"fmt"
	"log"
	"new-proto/api"
	"new-proto/api/config"
)

func main() {

	if err := config.Load(); err != nil {
		log.Fatalf("\n%+v\n", err)
	}

	svr := api.New()

	fmt.Println("Listening on localhost:8888")

	err := svr.ListenAndServe()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
