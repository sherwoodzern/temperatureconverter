package model

import "fmt"

type Temperature struct {
	Measure string `json:"measurement"`
	Degrees string `json:"degrees"`
}

func (t Temperature) GetMeasure() string {
	return t.Measure
}

func (t Temperature) GetDegrees() string {
	return t.Degrees
}

func (t Temperature) String() string {
	return fmt.Sprintf("%s, %s", t.Measure, t.Degrees)
}
