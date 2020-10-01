package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rubrikinc/rubrik-polaris-sdk-for-go/rubrikpolaris"
)

func main() {

	polaris := rubrikpolaris.Connect("rubrik-se", "drew.russell@rubrik.com", "3vy7Ezw8c69AG2j3i")

	go func() {
		for {
			radar24, err := polaris.RadarEvents24Hours()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(radar24)

			time.Sleep(time.Duration(3) * time.Second)

		}
	}()

	// keep application open until terminated
	for {
		time.Sleep(time.Duration(1) * time.Hour)
	}

}
