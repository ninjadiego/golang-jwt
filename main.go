package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, 世界")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
