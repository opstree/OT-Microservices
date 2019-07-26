package webapp

import (
    "net/http"
    "time"
    "fmt"
    "github.com/magiconair/properties"
)

func Run() {
    uptimeTicker := time.NewTicker(5 * time.Second)
    createDatabaseTable()
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    for {
        select {
        case <-uptimeTicker.C:
            healthCheck()
        }
    }
    http.ListenAndServe(":8080", nil)
}

func healthCheck() {
    healthCheck := "/etc/conf.d/ot-go-webapp/healthcheck.properties"
    healthVaules := properties.MustLoadFiles([]string{healthCheck}, properties.UTF8, true)

    healthy := healthVaules.GetString("healthy", "healthy")
    livecheck := healthVaules.GetString("livecheck", "livecheck")

    mux := http.NewServeMux()

    if healthy == "true" {
        mux.Handler("/healthy", returnCode200)
    } else {
        mux.Handler("/healthy", returnCode404)
    }

    if livecheck == "true" {
        mux.Handler("/livecheck", returnCode200)
    } else {
        mux.Handler("/livecheck", returnCode404)
    }
    http.ListenAndServe(":9000", nil)
}

func returnCode200(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("☄ HTTP status code returned!"))
}

func returnCode404(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("☄ HTTP status code returned!"))
}
