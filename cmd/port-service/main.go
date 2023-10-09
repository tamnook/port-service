package main

import (
	"fmt"
	"log"
	"port-service/internal/pkg/app"
)

func main() {
	fmt.Println("Hello")

	App := app.New()
	err := App.Run()
	if err != nil {
		log.Fatal(err)
	}
}
