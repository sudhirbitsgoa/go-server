package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//

	"log"
	"net/http"

	cache "go-server/cache"
	sw "go-server/mylib"
	sy "go-server/system"
)

func main() {
	router := sw.NewRouter()
	go sy.Iterator()    // enables concurrency
	go sy.CPUIterator() // enables concurrency
	cache.RedisConnect()
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
