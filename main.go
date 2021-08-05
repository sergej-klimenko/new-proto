package main

import (
	"cloud-native/api"
	"fmt"
	"log"
)

func main() {

	if err := api.LoadConfiguration(); err != nil {
		log.Fatalf("%+v\n", err)
	}

	conn, err := api.NewMongoConnection()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer conn.Disconnect()

	svr := api.New(api.Dependencies{MongoClient: conn.Client})

	fmt.Println("Listening on localhost:8888")

	err = svr.ListenAndServe()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
