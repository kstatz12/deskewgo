package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/home/karl/Documents/t.png")
	if err != nil {
		log.Fatal(err)
	}
	angle := detectAngle(f)
	fmt.Printf("%d", angle)
	deskew(f, angle)
}
