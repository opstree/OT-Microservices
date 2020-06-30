package main

import (
	"ot-go-webapp/config"
	"ot-go-webapp/elastic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"net/http"
	"strings"
)

// EmployeeInfo struct will be the data structure for employee's information
type EmployeeInfo struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	JobRole       string  `json:"job_role"`
	JoiningDate   string  `json:"joining_date"`
	Addresss      string  `json:"address"`
	City          string  `json:"city"`
	EmailID       string  `json:"email_id"`
	AnnualPackage float64 `json:"annual_package"`
	PhoneNumber   string  `json:"phone_number"`
}

func main() {
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	logrus.Infof("Running employee-management in webserver mode")
	logrus.Infof("employee-management is listening on port: %v", conf.Management.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v/create", conf.Management.APIPort)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/create", pushEmployeeData)
	router.GET("/search", fetchEmployeeData)
	router.GET("/search/all", fetchALLEmployeeData)
	router.GET("/search/roles", fetchEmployeeRoles)
	router.GET("/search/city", fetchEmployeeCity)
	router.Run(":" + conf.Management.APIPort)
}

func pushEmployeeData(c *gin.Context) {
	var request EmployeeInfo
	if err := c.BindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		return
	}

	info := EmployeeInfo{
		ID:            request.ID,
		Name:          request.Name,
		JobRole:       request.JobRole,
		JoiningDate:   request.JoiningDate,
		Addresss:      request.Addresss,
		City:          request.City,
		EmailID:       request.EmailID,
		AnnualPackage: request.AnnualPackage,
		PhoneNumber:   request.PhoneNumber,
	}
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	elastic.PostDataInSearch(conf, request.ID, info)
	logrus.Infof("Successfully pushed employee's data to elasticsearch")
}

func fetchEmployeeData(c *gin.Context) {
	searchQuery := c.Request.URL.Query()
	var searchValue string
	response := &EmployeeInfo{}

	for _, value := range searchQuery {
		searchValue = strings.Join(value, "")
	}
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchDataInElastic(conf, searchValue)

	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
	}
	logrus.Infof("Successfully fetched employee information")
	c.JSON(http.StatusOK, response)
}

func fetchALLEmployeeData(c *gin.Context) {
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchALLDataInElastic(conf)

	var employeeInfo []EmployeeInfo
	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		response := &EmployeeInfo{}
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
		employeeInfo = append(employeeInfo, *response)
	}
	logrus.Infof("Successfully fetched all employee's information")
	c.JSON(http.StatusOK, employeeInfo)
}

func fetchEmployeeRoles(c *gin.Context) {
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchALLDataInElastic(conf)

	var employeeInfo []EmployeeInfo
	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		response := &EmployeeInfo{}
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
		employeeInfo = append(employeeInfo, *response)
	}
	duplicate_frequency := make(map[string]int)
	// var values int
	var finalData []map[string]interface{}
	for _, role := range employeeInfo {
		_, exist := duplicate_frequency[role.JobRole]

		if exist {
			duplicate_frequency[role.JobRole] += 1
		} else {
			duplicate_frequency[role.JobRole] = 1 // else start counting from 1
		}
	}
	finalData = append(finalData, map[string]interface{}{
		"type":  "DevOps",
		"value": duplicate_frequency["DevOps"],
	})

	finalData = append(finalData, map[string]interface{}{
		"type":  "Developer",
		"value": duplicate_frequency["Developer"],
	})
	logrus.Infof("Successfully fetched all employee's roles")
	c.JSON(http.StatusOK, finalData)
}

func fetchEmployeeCity(c *gin.Context) {
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchALLDataInElastic(conf)

	var employeeInfo []EmployeeInfo
	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		response := &EmployeeInfo{}
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
		employeeInfo = append(employeeInfo, *response)
	}
	duplicate_frequency := make(map[string]int)
	for _, role := range employeeInfo {
		_, exist := duplicate_frequency[role.City]

		if exist {
			duplicate_frequency[role.City] += 1
		} else {
			duplicate_frequency[role.City] = 1 // else start counting from 1
		}
	}
	logrus.Infof("Successfully fetched all employee's city information")
	c.JSON(http.StatusOK, duplicate_frequency)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
