package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	time_n, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current time:", time.Now())
	fmt.Println("exact time:", time_n)

}
