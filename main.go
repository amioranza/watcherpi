package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var pinMap map[int]bool

func checkPINS() {
	for i := 0; i < 40; i++ {
		log.Printf("Checking PIN %v: %v\n", i, rpio.Pin(i).Read())
		// filling the pinMap with the pin status
		if rpio.Pin(i).Read() == 1 {
			pinMap[i] = true
		} else {
			pinMap[i] = false
		}
	}
}

func init() {
	// initialize the map variable to avoid PANIC!
	pinMap = make(map[int]bool)
}

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	for {
		log.Println("Checking PINS")
		checkPINS()
		for key, value := range pinMap {
			fmt.Println("| PIN ", key, "| STATUS ", value, "|")
		}
		log.Println("Sleeping 5s")
		time.Sleep(5000 * time.Millisecond)
	}
}
