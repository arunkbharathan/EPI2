package main

import "fmt"

func pipelines() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		fmt.Println("l")
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			fmt.Println("a")
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
					fmt.Println("b")
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{}, intStream <-chan int, multiplier int,
	) <-chan int {
		fmt.Println("m")
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			fmt.Println("c")
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
					fmt.Println("d")
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{}, intStream <-chan int, additive int,
	) <-chan int {
		fmt.Println("n")
		addedStream := make(chan int)
		go func() {
			fmt.Println("e")
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
					fmt.Println("f")
				}
			}
		}()
		return addedStream
	}
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	fmt.Println("g")
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	fmt.Println("h")
	for v := range pipeline {
		fmt.Println("i")
		fmt.Println(v)
		fmt.Println("j")
	}
	fmt.Println("k")
}
