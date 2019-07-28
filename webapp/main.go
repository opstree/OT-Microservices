package webapp

import (
    "net/http"
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
    mysql = healthCheckShow()
    handler := health.NewHandler()
    handler.AddChecker("MySQL", mysql)
    http.Handle("/health", healthCheckShow)
    http.ListenAndServe(":8080", nil)
}
