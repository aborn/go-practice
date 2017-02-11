package main

import (
	"./basic"
	"log"
)

func main() {
	for i := 0; i < 2005; i++ {
		basic.Put(i, i+1)
		if i == 2002 {
			log.Println("map[0]=", basic.Get(0), ", map[1]=", basic.Get(1),
				"map[2]=", basic.Get(2), ", map[3]=", basic.Get(3))
		}
	}
	log.Println("Hello Go !")
}
