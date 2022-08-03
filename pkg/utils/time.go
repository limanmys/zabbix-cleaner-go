package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func GetTime() (int, error) {
	interval, _ := strconv.Atoi(os.Getenv("INTERVAL"))
	interval = interval * -1
	deleteTime := time.Now().Add(time.Duration(interval) * 24 * time.Hour).Unix()

	//var table models.Alerts
	//database.Connection().Table("alerts").Where(fmt.Sprintf("clock < %d", deleteTime)).Delete(&table)

	return 1, errors.New("")
}
