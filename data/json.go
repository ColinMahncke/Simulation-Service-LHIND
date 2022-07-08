package data

import (
	"time"
)

//struct for json

type AreaCounting struct {
	ExternalId string `json:"externalId"`
	Value      int    `json:"value"`

	MeasurementTimestamp time.Time `json:"measurement_timestamp"`
	Unit                 string    `json:"unit"`
}

func NewAreaCounting(externalId string, value int, measurementTimestamp time.Time, unit string) *AreaCounting {

	a := AreaCounting{externalId, value, measurementTimestamp, unit}

	return &a
}

type LineCounting struct {
	ExternalId           string    `json:"externalId"`
	ValueIn              *int      `json:"valueIn"`
	ValueOut             *int      `json:"valueOut"`
	Unit                 string    `json:"unit"`
	MeasurementTimestamp time.Time `json:"measurement_timestamp"`
}

func NewLineCounting(externalId string, value int, measuremntTimestamp time.Time, unit string) *LineCounting {
	var valueIn *int
	var valueOut *int
	if value == 0 {
		return nil
	}
	if value < 0 {
		valueOut = &value
	} else {
		valueIn = &value
	}
	l := LineCounting{externalId, valueIn, valueOut, unit, measuremntTimestamp}

	return &l
}
