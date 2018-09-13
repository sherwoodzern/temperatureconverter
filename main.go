package main

import (
	"log"
	"net/http"

	"github.com/sherwoodzern/temperatureconverter/handlers"
)

func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
}

func main() {

	conversionHandler := handlers.ConversionHandler{}

	http.HandleFunc("/health", health)
	http.HandleFunc("/ftoc", conversionHandler.ConvertFtoC)
	http.HandleFunc("/ctof", conversionHandler.ConvertCtoF)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
