package webapp

import (
    "net/http"
    "github.com/giantswarm/retry-go"
)

func Run() {
    op := createDatabaseTable()
    retry.Do(op,
        retry.RetryChecker(IsNetOpErr)
        retry.Timeout(15 * time.Second))
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
