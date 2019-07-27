package main

import (
	"gitlab.com/opstree/ot-go-webapp/webapp"
	"time"
)

func main() {
	webapp.Run()
	for t := range time.NewTicker(2 * time.Second).C {
		webapp.healthCheck()
	}
}
