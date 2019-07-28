package webapp

import (
    "net/http"
)

func Run() {
    db := dbConn()
    createDatabaseTable()
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    mysql := dbcheck.NewMySQLChecker(db)
    handler := health.NewHandler()
    handler.AddChecker("MySQL", mysql)
    http.Handle("/health", handler)
    http.ListenAndServe(":8080", nil)
}
