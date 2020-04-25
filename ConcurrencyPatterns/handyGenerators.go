package main

import (
	"fmt"
	"math/rand"
	"time"
)

func handyGenerators() {
	repeat := func(
		done <-chan interface{}, values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer fmt.Println("exit repeat")
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
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

	// personal note: before repeat exits main exits
	done := make(chan interface{})
	// defer close(done)
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println("Loop done")
	close(done)
	time.Sleep(3 * time.Second)

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

	done = make(chan interface{})
	defer close(done)
	rand := func() interface{} {
		return rand.Int()
	}
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}
