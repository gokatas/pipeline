package main

import "fmt"

func counter() <-chan int {
	out := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			out <- x
		}
		close(out)
	}()
	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := counter()
	squares := squarer(naturals)
	printer(squares)
}
