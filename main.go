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
	pinStatus := false
	for i := 1; i < 41; i++ {
		pin := rpio.Pin(i).Read()
		// fmt.Printf(">>>>>>>>>>>>>>>>>>> %v - %v\n", i, pin)
		if pin == 1 {
			pinStatus = true
		} else {
			pinStatus = false
		}
		switch {
		case pinStatus == pinMap[i]:
			log.Printf("Pin %v (%v) status not changed\n", i, pinStatus)
		case pinStatus != pinMap[i]:
			pinMap[i] = pinStatus
			notify(i, pinStatus)
			log.Printf("Pin %v (%v) status changed\n", i, pinStatus)
		}
	}
}

func notify(pin int, status bool) {
	fmt.Printf("%v - %v\n", pin, status)
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
		log.Println("Sleeping 5s")
		time.Sleep(5000 * time.Millisecond)
	}
}
