package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	type Message struct {
		Text string
		CurrentTime string
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Handle("GET", "/", rootHandler)
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// pase json body.
	dec := json.NewDecoder(strings.NewReader(w.Body.String()))
	var m Message
	err2 := dec.Decode(&m)
	if err2 != nil {
		log.Fatal(err2)
	}

	if m.Text != "hello world" {
		t.Errorf("JSON Text value didn't return hello world")
	}

	if w.Code != http.StatusOK {
		fmt.Printf("%d - %s", w.Code, w.Body.String())
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
