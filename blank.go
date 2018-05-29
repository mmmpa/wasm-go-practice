package main

func main() {
	ch := make(chan string)

	go func(ch chan string) {
		ch <- "Hello, goroutine."
	}(ch)

	go func(ch chan string) {
		ch <- "Hello, goroutine."
	}(ch)

	println(<-ch)
	println(<-ch)
}
