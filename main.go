package main

import (
	cache "go-server/cache"
	sw "go-server/mylib"
	sy "go-server/system"
	"log"
	"net/http"
	"time"
)

func main() {
	router := sw.NewRouter()
	go sy.Iterator()    // enables concurrency
	go sy.CPUIterator() // enables concurrency
	cache.RedisConnect()
	test()
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func channel(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "sudhir"
}

func test() {
	ch := make(chan string, 1)
	go channel(ch)
	inp := <-ch
	for inp != "sudhir" {
		log.Println("blocking code")
	}
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
