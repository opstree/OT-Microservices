package main

import (
	"employee/config"
	"employee/elastic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"net/http"
	"os"
	"strings"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
)

// EmployeeInfo struct will be the data structure for employee's information
type EmployeeInfo struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	JobRole       string  `json:"job_role"`
	JoiningDate   string  `json:"joining_date"`
	Addresss      string  `json:"address"`
	Location      string  `json:"location"`
	Status        string  `json:"status"`
	EmailID       string  `json:"email"`
	AnnualPackage float64 `json:"annual_package"`
	PhoneNumber   string  `json:"phone_number"`
}

func main() {
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	logrus.Infof("Running employee-management in webserver mode")
	logrus.Infof("employee-management is listening on port: %v", conf.Employee.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v", conf.Employee.APIPort)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/employee/create", pushEmployeeData)
	router.GET("/employee/search", fetchEmployeeData)
	router.GET("/employee/search/all", fetchALLEmployeeData)
	router.GET("/employee/search/roles", fetchEmployeeRoles)
	router.GET("/employee/search/location", fetchEmployeeLocation)
	router.GET("/employee/search/status", fetchEmployeeStatus)
	router.GET("/employee/healthz", healthCheck)
	router.Run(":" + conf.Employee.APIPort)
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
		Location:      request.Location,
		Status:        request.Status,
		EmailID:       request.EmailID,
		AnnualPackage: request.AnnualPackage,
		PhoneNumber:   request.PhoneNumber,
	}
	conf, err := config.ParseFile(configFile)
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
	conf, err := config.ParseFile(configFile)
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
	conf, err := config.ParseFile(configFile)
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
	conf, err := config.ParseFile(configFile)
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
		_, exist := duplicate_frequency[role.JobRole]

		if exist {
			duplicate_frequency[role.JobRole] += 1
		} else {
			duplicate_frequency[role.JobRole] = 1 // else start counting from 1
		}
	}
	logrus.Infof("Successfully fetched all employee's roles")
	c.JSON(http.StatusOK, duplicate_frequency)
}

func fetchEmployeeLocation(c *gin.Context) {
	conf, err := config.ParseFile(configFile)
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
		_, exist := duplicate_frequency[role.Location]

		if exist {
			duplicate_frequency[role.Location] += 1
		} else {
			duplicate_frequency[role.Location] = 1
		}
	}
	logrus.Infof("Successfully fetched all employee's Location information")
	c.JSON(http.StatusOK, duplicate_frequency)
}

func fetchEmployeeStatus(c *gin.Context) {
	conf, err := config.ParseFile(configFile)
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
		_, exist := duplicate_frequency[role.Status]

		if exist {
			duplicate_frequency[role.Status] += 1
		} else {
			duplicate_frequency[role.Status] = 1
		}
	}
	logrus.Infof("Successfully fetched all employee's status information")
	c.JSON(http.StatusOK, duplicate_frequency)
}

func healthCheck(c *gin.Context) {
	conf, err := config.ParseFile(configFile)
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}

	status, err := elastic.CheckElasticHealth(conf)

	if err != nil {
		logrus.Errorf("Error while getting elasticsearch health: %v", err)
		errorResponse(c, http.StatusBadRequest, "Elasticsearch is not running")
		return
	}

	if status != false {
		c.JSON(http.StatusOK, gin.H{
			"status":   "up",
			"database": "elasticsearch",
			"message":  "Elasticsearch is running",
		})
		return
	}

	errorResponse(c, http.StatusBadRequest, "Elasticsearch is not running")
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
