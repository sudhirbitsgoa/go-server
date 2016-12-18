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

	cache "server/cache"
	sw "server/mylib"
	sy "server/system"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	sy.Iterator()
	sy.CPUIterator()
	cache.RedisConnect()
	log.Fatal(http.ListenAndServe(":8080", router))
}
