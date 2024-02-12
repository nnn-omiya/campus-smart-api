package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AccessLogHandler[T interface{}](r *http.Request) (*T, error) {
	fmt.Println("Method:", r.Method)
	fmt.Println("URI:", r.RequestURI)
	switch r.Method {
	case "GET":
		r.ParseForm()
		fmt.Println("Parameters:")
		for key, values := range r.Form {
			for _, value := range values {
				fmt.Printf("\t%s: %s\n", key, value)
			}
		}
	case "POST":
		var body T
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return nil, err
		} else {
			fmt.Println("Body:", body)
		}
		return &body, nil
	}
	return nil, nil
}

func ErrorHandler(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
	fmt.Println(err.Error(), status)
}
