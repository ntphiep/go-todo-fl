package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main_te() {
	var next1 http.Handler
	var next2 http.Handler

	r1 := gin.Default()
	r2 := mux.NewRouter()

	next1 = r1
	next2 = r2

	fmt.Println(next1 == next2)

	// fmt.Println("Hello, World!")
}
