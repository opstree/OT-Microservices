package main

import (
	"ot-go-webapp/config"
	"ot-go-webapp/elastic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"strings"
	"net/http"
	"time"
	"os"
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
	City          string  `json:"city"`
	EmailID       string  `json:"email_id"`
	AnnualPackage float64 `json:"annual_package"`
	PhoneNumber   string  `json:"phone_number"`
}

// SalaryInfo struct will be the data structure for employee's information
type SalaryInfo struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Month         string  `json:"month"`
	Salary        float64  `json:"salary"`
}

func main() {
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for salary: %v", err)
	}
	logrus.Infof("Running employee-salary in webserver mode")
	logrus.Infof("employee-salary is listening on port: %v", conf.Salary.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v/create", conf.Salary.APIPort)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.GET("/salary/search", fetchEmployeeSalary)
	router.Run(":" + conf.Salary.APIPort)
}

func fetchEmployeeSalary(c *gin.Context) {
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
	
	salaryData := SalaryInfo{
		ID: response.ID,
		Name: response.Name,
		Salary: (response.AnnualPackage / 12),
		Month: time.Now().UTC().Format("Jan"),
	}
	c.JSON(http.StatusOK, salaryData)
}
