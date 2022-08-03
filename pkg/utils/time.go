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

	//var table models.Alerts
	//database.Connection().Table("alerts").Where(fmt.Sprintf("clock < %d", deleteTime)).Delete(&table)

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
