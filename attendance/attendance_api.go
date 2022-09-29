package main

import (
	"attendance/config"
	"database/sql"
	// 	_ "github.com/go-sql-driver/mysql"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/mysql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"net/http"
	"os"
)

const (
	dbDriver = "mysql"
)

var (
	configFile = os.Getenv("CONFIG_FILE")
)

// AttendanceInfo struct will be the data structure for employee's attendance
type AttendanceInfo struct {
	ID     int    `json:"id"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

func main() {
	conf, err := config.ParseFile(configFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for attendance: %v", err)
	}
	logrus.Infof("Running employee-attendance in webserver mode")
	logrus.Infof("employee-attendance api is listening on port: %v", conf.Attendance.APIPort)
	logrus.Infof("Endpoint is available now - http://0.0.0.0:%v/create", conf.Attendance.APIPort)
	router := gin.Default()
	router.Use(apmgin.Middleware(router))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.POST("/attendance/create", pushAttendanceData)
	router.GET("/attendance/search", fetchAttendanceData)
	router.GET("attendance/healthz", healthCheckMySQL)
	router.Run(":" + conf.Attendance.APIPort)
}

func initDBConnection() (*sql.DB, error) {
	conf, err := config.ParseFile(configFile)
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for attendance: %v", err)
	}
	db, err := apmsql.Open(dbDriver, conf.MySQL.Username+":"+conf.MySQL.Password+"@tcp("+conf.MySQL.Host+")/"+conf.MySQL.DBName)
	if err != nil {
		return db, err
	}
	return db, nil
}

func pushAttendanceData(c *gin.Context) {
	conf, err := config.ParseFile(configFile)
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for attendance: %v", err)
	}
	db, err := initDBConnection()
	tx, err := db.BeginTx(c.Request.Context(), nil)
	if err != nil {
		logrus.Errorf("Error while creating sql connection for pushing attendance data: %v", err)
	}
	_, err = tx.ExecContext(c.Request.Context(), "USE "+conf.MySQL.DBName)
	if err != nil {
		logrus.Errorf("Not able to use database: %v", err)
	}
	_, err = tx.ExecContext(c.Request.Context(), "CREATE TABLE IF NOT EXISTS Employee ( id int(6) NOT NULL, status varchar(50) NOT NULL, date varchar(50), PRIMARY KEY (id) )")
	if err != nil {
		logrus.Errorf("Error while creating Table: %v", err)
	}
	defer db.Close()

	var request AttendanceInfo
	if err := c.BindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		return
	}
	insForm, err := tx.Prepare("INSERT INTO Employee(id, status, date) VALUES(?,?,?)")

	if err != nil {
		logrus.Errorf("Cannot create db insertion command: %v", err)
	}

	insForm.ExecContext(c.Request.Context(), request.ID, request.Status, request.Date)
	if err := tx.Commit(); err != nil {
		logrus.Errorf("Cannot create db insertion command: %v", err)
	}
	defer db.Close()
	logrus.Infof("Successfully pushed employee's attendance information")
}

func fetchAttendanceData(c *gin.Context) {
	db, err := initDBConnection()
	tx, err := db.BeginTx(c.Request.Context(), nil)
	if err != nil {
		logrus.Errorf("Error while creating sql connection for fetching attendance data: %v", err)
	}
	selDB, err := tx.QueryContext(c.Request.Context(), "SELECT * FROM Employee ORDER BY id DESC")

	var attendanceInfo []AttendanceInfo

	for selDB.Next() {
		var id int
		var status, date string
		err = selDB.Scan(&id, &status, &date)
		if err != nil {
			logrus.Errorf("Error while scanning data: %v", err)
		}

		attendanceInfo = append(attendanceInfo, AttendanceInfo{
			ID:     id,
			Status: status,
			Date:   date,
		})
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("Cannot create db fetch command: %v", err)
	}
	c.JSON(http.StatusOK, attendanceInfo)
}

func healthCheckMySQL(c *gin.Context) {
	db, err := initDBConnection()
	if err != nil {
		logrus.Errorf("Error while creating sql connection for fetching attendance data: %v", err)
	}

	err = db.PingContext(c.Request.Context())
	if err != nil {
		logrus.Errorf("Unable to communicate with MySQL database: %v", err)
		errorResponse(c, http.StatusBadRequest, "MySQL connection is not up")
		return
	}
	defer db.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":   "up",
		"database": "MySQL",
		"message":  "MySQL is running",
	})
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
