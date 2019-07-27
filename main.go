package main

import (
	"gitlab.com/opstree/ot-go-webapp/webapp"
	"time"
	"fmt"
)

func main() {
	webapp.Run()
	for t := range time.NewTicker(2 * time.Second).C {
		fmt.Println(t)
		webapp.HealthCheck()
	}
}
