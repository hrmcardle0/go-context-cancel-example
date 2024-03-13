package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	myCtx := context.Background()            // master context
	ctx, cancel := context.WithCancel(myCtx) // parent context
	childCtx, _ := context.WithCancel(ctx)   // child context
	var wg sync.WaitGroup

	// call parent cancel after 3 seconds
	go func(cancel func()) {
		log.Println("goroutine started")
		select {
		case <-time.After(3 * time.Second):
			log.Println("3 second timer up!")
			cancel()

		}
		log.Println("goroutine ended")

	}(cancel)
	wg.Add(2)

	//
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("MASTER: started checker goroutine")
		select {
		case <-ctx.Done():
			log.Println("MASTER: channel done called!")
		case <-time.After(8 * time.Second):
			log.Println("MASTER: done waiting, done never called")
			return
		}
	}(myCtx, &wg)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("PARENT: started checker goroutine")
		select {
		case <-ctx.Done():
			log.Println("PARENT: channel done called")
		case <-time.After(6 * time.Second):
			log.Println("PARENT: done waiting, done never called")
			return
		}
	}(childCtx, &wg)
loop:
	for {
		log.Println("Looping")
		select {
		case <-ctx.Done():
			log.Println("Context done() called")
			break loop
		}

	}

	log.Println("Waiting for goroutine to finish")
	wg.Wait()
	log.Println("All Done. Exiting")

}
