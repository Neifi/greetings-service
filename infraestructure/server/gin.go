package server

import (
	"github.com/gin-gonic/gin"
	"greetings-service/infraestructure/rest/greetings"
	"log"
	"net/http"
)

func Run() {
	r := gin.Default()

	v1 := r.Group("api/v1/")

	v1.GET("/hello", greetings.HelloHandler())

	err := r.Run()
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":8080", r))
}
