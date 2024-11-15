package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// http.HandleFunc("/", handler)
	// http.HandleFunc("/root", root)
	//
	// log.Println("server running on localhost:8000")
	// http.ListenAndServe(":8000", nil)

	router := gin.Default()

	router.GET("/", home)
	log.Println("server running on localhost:8000")
}

//http packages: these are used to establish a connection between a client (web browser) and a server (your base)
//DIFFERENT HTTP METHODS: GET, POST, PUT, PATCH, DELETE.

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "hello welcome to our web page")
		return
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error while parsing", http.StatusBadRequest)
		}
		name := r.FormValue("name")
		age := r.FormValue("age")

		fmt.Fprintf(w, "%s is %s years old", name, age)
		return
	}

	fmt.Fprintf(w, "we are only processing GET and POST requests")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the root url: %s", r.URL.Path[1:])
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, "welcome to my home page")
	// testing new branch
}
