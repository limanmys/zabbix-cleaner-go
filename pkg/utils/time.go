package utils

import (
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func GetTime() (int, error) {
	interval, err := getInterval()
	if err != nil {
		return -1, err
	}

	time := time.Now().Add(time.Duration(interval) * 24 * time.Hour).Unix()

	return int(time), nil
}

func getInterval() (int, error) {
	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		return -1, err
	}
	interval = interval * -1

	return interval, nil
}
