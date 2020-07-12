package main

import (
	"notification/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"net/http"
	"net/smtp"
	"os"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
	auth       smtp.Auth
)

// EmployeeInfo struct will be the data structure for employee's information
type EmployeeInfo struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	JobRole       string  `json:"job_role"`
	JoiningDate   string  `json:"joining_date"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	EmailID       string  `json:"email_id"`
	AnnualPackage float64 `json:"annual_package"`
	PhoneNumber   string  `json:"phone_number"`
}

func main() {
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for notification: %v", err)
	}
	logrus.Infof("Running notification in webserver mode")
	logrus.Infof("notification is listening on port: %v", conf.Notification.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v/create", conf.Notification.APIPort)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/notification/send", sendNotification)
	router.GET("/notification/healthz", healthCheck)
	router.Run(":" + conf.Notification.APIPort)
}

func sendNotification(c *gin.Context) {
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for notification: %v", err)
	}
	var request EmployeeInfo
	if err := c.BindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		return
	}
	auth = smtp.PlainAuth("", conf.SMTP.Username, conf.SMTP.Password, conf.SMTP.SMTPServer)
	to := request.EmailID

	msg := "From: " + conf.SMTP.From + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Your user is added in Opstree Employee Service"
	err = smtp.SendMail(conf.SMTP.SMTPServer+":"+conf.SMTP.SMTPPort, auth, conf.SMTP.From, []string{to}, []byte(msg))
	if err != nil {
		logrus.Errorf("Unable to send mail %v", err)
		errorResponse(c, http.StatusBadRequest, "Unable to send mail")
	}
	c.JSON(http.StatusOK, "Successfully sent notification")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "Notification service is running successfully",
		"status_code": 200,
	})
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
