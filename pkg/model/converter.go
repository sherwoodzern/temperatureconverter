package model

import "fmt"

type Converter struct {
	Value string `json: "value"`
}

func (i Converter) GetValue() string {
	return i.Value
}

func (i Converter) String() string {
	return fmt.Sprintf("%g", i)
}
