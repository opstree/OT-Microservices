package main

import (
	"gitlab.com/opstree/ot-go-webapp/webapp"
	"time"
	"fmt"
	"net/http"
)

func main() {
	for t := range time.NewTicker(2 * time.Second).C {
		fmt.Println(t)
		mux := webapp.HealthCheck()
		http.ListenAndServe(":9000", mux)
		webapp.Run()
	}
}
