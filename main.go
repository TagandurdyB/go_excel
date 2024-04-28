package main

import (
	"fmt"
	"net/http"

	"auto_excel/config"
	models "auto_excel/models"
)

func main() {
	models.Person{}.Migrate()
	fmt.Println("Programma işledi!")
	fmt.Println("Brauzeri açyň we 127.0.0.1:8080 ýazyň!")
	http.ListenAndServe(":8080", config.Routes())
}
