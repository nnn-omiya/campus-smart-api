package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/nnn-omiya/campus-smart-api/routes"
)

type Config struct {
	DBName string
	User   string
	Passwd string
	Addr   string
	Port   string
}

func NewConfig() (*Config, error) {
	return &Config{
		DBName: os.Getenv("MYSQL_DATABASE"),
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr:   os.Getenv("MYSQL_ADDR"),
		Port:   os.Getenv("MYSQL_PORT"),
	}, nil
}

func connectDB(c *mysql.Config) (*sql.DB, error) {
	MaxDBRetryCount := 10
	SleepTime := 5 * time.Second

	var db *sql.DB
	var err error

	for r := 1; r <= MaxDBRetryCount; r++ {
		log.Println("NewDB Connection Attempt #" + strconv.Itoa(r))
		db, err = sql.Open("mysql", c.FormatDSN())
		if err != nil {
			log.Println("NewDB Connection Error:" + err.Error())
			time.Sleep(SleepTime)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Println("NewDB Connection Error:" + err.Error())
			time.Sleep(SleepTime)
			continue
		} else {
			break
		}
	}

	if err != nil {
		log.Println("NewDB Connection Failed")
		return nil, err
	}

	return db, nil
}

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := mysql.Config{
		DBName:    config.DBName,
		User:      config.User,
		Passwd:    config.Passwd,
		Addr:      fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}

	db, err := connectDB(&c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", routes.Router(db))

	port := "8080"
	log.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
