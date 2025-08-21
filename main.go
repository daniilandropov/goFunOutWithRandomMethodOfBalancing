package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func FanOut[T any](chIn <-chan T, n int) []<-chan T {
	res := make([]chan T, n)
	for i := range res {
		res[i] = make(chan T, cap(chIn))
	}

	go func() {
		idx := 0
		for v := range chIn {
			switch rand.Intn(3) {
			case 0:
				idx = (idx + 1) % n
				select {
				case res[idx] <- v:
				default:
					idx = (idx + 1) % n
					res[idx] <- v
				}

			case 1:
			loop:
				for {
					select {
					case res[idx] <- v:
						break loop
					default:
						idx = (idx + 1) % n
					}
				}
			case 2:
				minId := 0
				minLen := len(res[minId])
				for i, ch := range res {
					l := len(ch)
					if l < minLen {
						minId = i
						minLen = l
					}
				}

				res[minId] <- v
			}
		}

		for _, r := range res {
			close(r)
		}
	}()

	result := make([]<-chan T, n)

	for i, r := range res {
		result[i] = r
	}

	return result
}

func main() {
	ch := make(chan int, 5000)

	for i := 0; i < 5000; i++ {
		ch <- i
	}
	close(ch)

	wg := sync.WaitGroup{}
	wg.Add(50)

	for _, c := range FanOut(ch, 50) {
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println(v)
			}
		}()
	}

	wg.Wait()
}
