package main

import (
	"crud/router"
	"fmt"
	"net/http"
)

func main() {
	routes := router.MuxRouter()
	fmt.Println("Server on port:3000")
	http.ListenAndServe(":3000", routes)

}