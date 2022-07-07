package data

import (
	"math"
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
	Value                int       `json:"value"`
	Unit                 string    `json:"unit"`
	MeasurementTimestamp time.Time `json:"measurement_timestamp"`
	MeasurementType      string    `json:"measurement_type"`
}

func NewLineCounting(externalId string, value int, measuremntTimestamp time.Time, unit string) *LineCounting {
	var measurementType string
	if value == 0 {
		return nil
	}
	if value < 0 {
		measurementType = "flow_out"
	} else {
		measurementType = "flow_in"
	}
	l := LineCounting{externalId, int(math.Abs(float64(value))), unit, measuremntTimestamp, measurementType}

	return &l
}
