package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_main(t *testing.T) {
	mockResponse := `{"database":"elasticsearch","message":"Elasticsearch is running","status":"up"}`
	r := SetUpRouter()
	r.GET("/employee/healthz", healthCheck)
	req, _ := http.NewRequest("GET", "/employee/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_pushEmployeeData(t *testing.T) {
	r := SetUpRouter()
	r.POST("/employee/create", pushEmployeeData)
	employeeId := xid.New().String()

	employeeInfo := EmployeeInfo{
		ID:            employeeId,
		Name:          "Abhishek Dubey",
		JobRole:       "DevOps",
		JoiningDate:   "25-09-2017",
		Addresss:      "Nangloi",
		Location:      "New Delhi",
		Status:        "Current Employee",
		EmailID:       "abhishek@example.com",
		AnnualPackage: 10000,
		PhoneNumber:   "9999999999",
	}

	jsonValue, _ := json.Marshal(employeeInfo)
	req, _ := http.NewRequest("POST", "/employee/create", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}


func Test_fetchALLEmployeeData(t *testing.T) {
	r := SetUpRouter()
	r.GET("/employee/search/all", fetchALLEmployeeData)
	req, _ := http.NewRequest("GET", "/employee/search/all", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
