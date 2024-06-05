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
	"github.com/gorilla/mux"

	"github.com/nnn-omiya/campus-smart-api/routes"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	c := mysql.Config{
		DBName: os.Getenv("MYSQL_DATABASE"),
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("MYSQL_ADDR"), os.Getenv("MYSQL_PORT")),
		// Addr:      fmt.Sprintf("%s:%s", os.Getenv("MYSQL_ADDR"), "3316"),
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}

	MaxDBRetryCount := 10
	SleepTime := 5 * time.Second

	db, err := sql.Open("mysql", c.FormatDSN())
	for r := 1; r <= MaxDBRetryCount; r++ {
		fmt.Println("NewDB Connection Attempt #" + strconv.Itoa(r))
		db, err = sql.Open("mysql", c.FormatDSN())
		if err != nil {
			fmt.Println("NewDB Connection Error:" + err.Error())
			time.Sleep(SleepTime)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Println("NewDB Connection Error:" + err.Error())
			time.Sleep(SleepTime)
			continue
		} else {
			break
		}
	}

	if err != nil {
		fmt.Println("NewDB Connection Failed")
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	routes.Router(db, router)

	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	// t := time.NewTicker(5 * time.Minute)
	// defer func() {
	// 		fmt.Println("Stopping ticker...")
	// 		t.Stop()
	// }()

	// for {
	// 		select {
	// 		case now := <-t.C:
	// 				fmt.Println(now.Format(time.RFC3339))
	// 		}
	// }

}
