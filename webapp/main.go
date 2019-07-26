package webapp

import (
    "net/http"
    "time"
    "github.com/magiconair/properties"
)

func Run() {
    createDatabaseTable()
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    doEvery(100*time.Millisecond, healthCheck)
    http.ListenAndServe(":8080", nil)
}

func healthCheck() {
    healthCheck := "/etc/conf.d/ot-go-webapp/healthcheck.properties"
    healthVaules := properties.MustLoadFiles([]string{healthCheck}, properties.UTF8, true)

    healthy := healthCheck.GetString("healthy", "healthy")
    livecheck := healthCheck.GetString("livecheck", "livecheck")

    if healthy == "true" {
        http.HandleFunc("/healthy", returnCode200)
    } else {
        http.HandleFunc("/healthy", returnCode404)
    }

    if livecheck == "true" {
        http.HandleFunc("/livecheck", returnCode200)
    } else {
        http.HandleFunc("/livecheck", returnCode404)
    }
}

func returnCode200(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("☄ HTTP status code returned!"))
}

func returnCode404(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("☄ HTTP status code returned!"))
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
