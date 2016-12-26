package main

import (
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

// func makeChan() {
// 	done := make(chan string, 1)
// 	go func(done chan string) {
// 		fmt.Print("working...")
// 		time.Sleep(time.Second)
// 		fmt.Println("done")
// 		// Send a value to notify that we’re done.
// 		done <- "true"
// 	}(done)
// 	c := <-done
// 	print(c)
// }
