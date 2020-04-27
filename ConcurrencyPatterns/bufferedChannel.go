package main

// func buffereChannel() {
// 	take := func(
// 		done <-chan interface{}, valueStream <-chan interface{}, num int,
// 	) <-chan interface{} {
// 		takeStream := make(chan interface{})
// 		go func() {
// 			defer close(takeStream)
// 			for i := 0; i < num; i++ {
// 				select {
// 				case <-done:
// 					return
// 				case takeStream <- <-valueStream:
// 				}
// 			}
// 		}()
// 		return takeStream
// 	}
// 	repeat := func(
// 		done <-chan interface{}, values ...interface{},
// 	) <-chan interface{} {
// 		valueStream := make(chan interface{})
// 		go func() {
// 			defer fmt.Println("exit repeat")
// 			defer close(valueStream)
// 			for {
// 				for _, v := range values {
// 					select {
// 					case <-done:
// 						return
// 					case valueStream <- v:
// 					}
// 				}
// 			}
// 		}()
// 		return valueStream
// 	}
// 	sleep := func(
// 		done <-chan interface{}, valueStream <-chan interface{}, num time.Duration,
// 	) <-chan interface{} {
// 		sleepStream := make(chan interface{})
// 		go func() {
// 			defer close(sleepStream)
// 			for {
// 				select {
// 				case <-done:
// 					return
// 				case <-time.After(num * time.Second):
// 					sleepStream <- valueStream
// 				}
// 			}

// 		}()
// 		return sleepStream
// 	}

// 	buffer := func(
// 		done <-chan interface{}, valueStream <-chan interface{}, num int,
// 	) <-chan interface{} {
// 		bufferedStream := make(chan interface{}, num)
// 		go func() {
// 			defer close(bufferedStream)
// 			for {
// 				select {
// 				case <-done:
// 					return
// 				case bufferedStream <- valueStream:
// 				}
// 			}

// 		}()
// 		return bufferedStream
// 	}

// 	done := make(chan interface{})
// 	defer close(done)
// 	zeros := take(done, 3, repeat(done, 0))
// 	short := sleep(done, 1*time.Second, zeros)
// 	buffer := buffer(done, 2, short) // Buffers sends from short by 2 long := sleep(done, 4*time.Second, short)
// 	pipeline := long

// }
