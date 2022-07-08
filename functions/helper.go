package functions

import (
	"math/rand"
	"os"
	"time"
)

func GetTimeToHours() float64 {

	time := time.Now()

	return (float64(time.Hour()) + (float64(time.Minute()))/60.0 + float64(time.Second())/3600.0 + float64(time.Nanosecond())/36e11)

}

//fallback for different enviroment variables
func Getenv(name string, fallback string) string {
	variable, found := os.LookupEnv(name)

	if !found {
		variable = fallback
	}
	return variable
}

//generate random seed for a higher randomness
func randIntn(min float64, max float64) int {

	if int(max-min) == 0.0 {
		return 0
	}
	return int(min) + rand.Intn(int(max-min))
}
