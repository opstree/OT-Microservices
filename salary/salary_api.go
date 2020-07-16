package main

import (
	"salary/config"
	"salary/elastic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fileName = "delay.txt"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
	delayTime  = os.Getenv("DELAY_TIME")
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
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Month  string  `json:"month"`
	Salary float64 `json:"salary"`
}

func main() {
	var waitTime int
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for salary: %v", err)
	}
	logrus.Infof("Running employee-salary in webserver mode")
	logrus.Infof("employee-salary is listening on port: %v", conf.Salary.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v/create", conf.Salary.APIPort)

	if delayTime == "" {
		waitTime = 1
	} else {
		waitTime, _ = strconv.Atoi(delayTime)
	}
	time.Sleep(time.Duration(waitTime) * time.Second)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/salary/configure/liveness", configureLiveness)
	router.GET("/salary/search", fetchEmployeeSalary)
	router.GET("/salary/healthz", healthCheck)
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
		ID:     response.ID,
		Name:   response.Name,
		Salary: (response.AnnualPackage / 12),
		Month:  time.Now().UTC().Format("Jan"),
	}
	c.JSON(http.StatusOK, salaryData)
}

func healthCheck(c *gin.Context) {
	var waitTime int
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

	if Exists(fileName) {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			logrus.Errorf("Delay file doesn't exists: %v", err)
		}
		waitTime, _ = strconv.Atoi(string(content))
	} else {
		waitTime = 1
	}

	logrus.Infof("Response is slow by: %v seconds", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Second)

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

func configureLiveness(c *gin.Context) {
	searchQuery := c.Request.URL.Query()
	var searchValue string

	for _, value := range searchQuery {
		searchValue = strings.Join(value, "")
	}

	file, err := os.Create(fileName)

	if err != nil {
		logrus.Errorf("Unable to set delay period: %v", err)
		errorResponse(c, http.StatusBadRequest, "Unable to set delay period")
		return
	}

	defer file.Close()

	file.WriteString(searchValue)
}

// Exists function checks if file exists or not
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
