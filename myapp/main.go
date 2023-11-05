package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("%v\n", time.Now())
		time.Sleep(time.Second)
	}
}
