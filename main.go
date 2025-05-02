package main

import (
	// "fmt"
	// "example.com/myapp/helloworld"
	"os"
	"time"

	"example.com/myapp/countdown"
)

func main() {
	// fmt.Println(helloworld.Hello("World", ""))
	sleeper := countdown.New(1*time.Second, time.Sleep)
	countdown.Countdown(os.Stdout, &sleeper)
}
