package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

type ApiError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *ApiError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *ApiError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *ApiError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func CustomeError(err error, status int, detail string) error {
	return &ApiError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here and Log the error.
		log.Printf("An error accured: %v", err)

		//lets check error type
		customError, ok := err.(ClientError)

		if !ok {
			//error is not http error type ,
			w.WriteHeader(500) // retrun 500 internal server error
			return
		}

		body, err := customError.ResponseBody()
		if err != nil {
			log.Printf("An error accured: %v", err)
			w.WriteHeader(500)
			return
		}
		status, headers := customError.ResponseHeaders() // Get http status code and headers.
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
	}
}
