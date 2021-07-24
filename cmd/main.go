package main

import (
	"log"

	greeting "github.com/german9304/demo/greeting"
)

func main() {
	result := greeting.Hello("world")
	log.Println(result)
}
