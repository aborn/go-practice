package main

import (
	"./basic"
	"log"
)

func main() {
	for i := 0; i < 2000; i++ {
		basic.Put(i, i+1)
	}
	log.Println("Hello Go !")
}
