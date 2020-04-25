package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanInfanOut() {
	toInt := func(
		done <-chan interface{}, valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	repeatFn := func(
		done <-chan interface{}, fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	take := func(
		done <-chan interface{}, valueStream <-chan interface{}, num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer fmt.Println("exit take")
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	fanIn := func(
		done <-chan interface{}, channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})
		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		// Select from all the channels
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	primeFinder := func(
		done <-chan interface{}, valueStream <-chan int,
	) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer fmt.Println("exit prime")
			defer close(primeStream)

			for val := range valueStream {
				// fmt.Println(val)
				half := val / 2

				for i := half; i >= 2; i-- {
					if val%i == 0 {
						val = 0
						break
					}
				}

				select {
				case <-done:
					return
				case primeStream <- val:

				}

			}
		}()
		return primeStream
	}

	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	// _, _ = fanIn, finders
	for prime := range take(done, fanIn(done, finders...), 1000) {
		if prime != 0 {
			fmt.Printf("\t%d\n", prime)
		}
	}
	fmt.Printf("Search took: %v", time.Since(start))

}
