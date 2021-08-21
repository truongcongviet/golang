package main

import (
	"fmt"
	"friend-management-v1/cmd/handler/router"
	"net/http"
)

func main() {
	r := router.SetUpRouter()
	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}
