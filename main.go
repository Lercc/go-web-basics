package main

import (
	"crud/router"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err == nil {
		fmt.Println("ENV variables loaded success")
	} else {
		fmt.Println("ENV variables loaded failed")
	}
}

func main() {
	routes := router.MuxRouter()

	fmt.Println("Server on port:3000")
	http.ListenAndServe(":3000", routes)
}