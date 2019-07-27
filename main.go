package main

import (
	"gitlab.com/opstree/ot-go-webapp/webapp"
	"time"
	"fmt"
)

func main() {
	for t := range time.NewTicker(2 * time.Second).C {
		fmt.Println(t)
		webapp.HealthCheck()
		webapp.Run()
	}
}
