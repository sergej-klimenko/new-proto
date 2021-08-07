package main

import (
	"fmt"
	"log"
	"new-proto/api"
)

func main() {

	if err := api.LoadConfiguration(); err != nil {
		log.Fatalf("\n%+v\n", err)
	}

	svr := api.New()

	fmt.Println("Listening on localhost:8888")

	err := svr.ListenAndServe()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
