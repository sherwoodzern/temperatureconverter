package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sherwoodzern/temperatureconverter/pkg/convert"
	"github.com/sherwoodzern/temperatureconverter/pkg/model"
)

// ConversionHandler Interface for conversion handlers
type ConversionHandler struct{}

// ConvertFtoC Handler to convert Fahrenheit to Celsius
func (c ConversionHandler) ConvertFtoC(writer http.ResponseWriter, request *http.Request) {
	result, err := decodeRequest(request)

	if err != nil {
		handleError(writer, err)
		return
	}

	celsius := convert.FahrenheitToCelsius(convert.Fahrenheit(result))
	formatAndSendResponse(writer, "Celsius", celsius.String())
}

// ConvertCtoF Handler to convert Celsius to Fahrenheit
func (c ConversionHandler) ConvertCtoF(writer http.ResponseWriter, request *http.Request) {
	result, err := decodeRequest(request)

	if err != nil {
		handleError(writer, err)
		return
	}

	fahrenheit := convert.CelsiusToFahrenheit(convert.Celsius(result))
	formatAndSendResponse(writer, "Fahrenheit", fahrenheit.String())
}

func handleError(writer http.ResponseWriter, err error) {

	e := &model.ConversionError{
		ErrorCode:    500,
		ErrorMessage: err.Error(),
	}

	//b, err := json.NewDecoder(e)
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(500)
	//writer.Write(b)
	encoder := json.NewEncoder(writer)
	encoder.Encode(e)

}

func decodeRequest(request *http.Request) (float64, error) {
	var input model.Converter
	decoder := json.NewDecoder(request.Body)

	//fmt.Printf("The ContentLength: %v\n" + strconv.FormatInt(request.ContentLength, 10))
	//fmt.Printf("The HOST: %s\n" + request.Host)

	err := decoder.Decode(&input)

	//fmt.Printf("The input value: %s\n" + input.String())

	if err != nil {
		fmt.Println("Decoder Failure")
		return 0, err
	}

	value, parseErr := strconv.ParseFloat(input.GetValue(), 64)
	if parseErr != nil {
		fmt.Println("ParseFloat Failure")
		return 0, parseErr
	}
	return value, nil
}

func formatAndSendResponse(writer http.ResponseWriter, measurement string, result string) {

	response := &model.Temperature{
		Measure: measurement,
		Degrees: result,
	}
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
