package main

import (
	"time"

	"github.com/limanmys/zabbix-cleaner-go/pkg/models"
	"github.com/limanmys/zabbix-cleaner-go/pkg/operations"
	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
)

func main() {
	utils.InitLogger()

	for {
		var auditog []models.Audit
		err := operations.Delete("auditlog", auditog)
		if err != nil {
			utils.Logger.Printf("error when deleting auditlog data err:%s", err.Error())
		}

		var events []models.Events
		err = operations.Delete("events", events)
		if err != nil {
			utils.Logger.Printf("error when deleting events data err:%s", err.Error())
		}

		var acknowledges []models.Acknowledges
		err = operations.Delete("acknowledges", acknowledges)
		if err != nil {
			utils.Logger.Printf("error when deleting acknowledges data err:%s", err.Error())
		}

		var history []models.History
		err = operations.Delete("history", history)
		if err != nil {
			utils.Logger.Printf("error when deleting history data err:%s", err.Error())
		}

		var history_str []models.History
		err = operations.Delete("history_str", history_str)
		if err != nil {
			utils.Logger.Printf("error when deleting history_str data err:%s", err.Error())
		}

		var history_text []models.History
		err = operations.Delete("history_text", history_text)
		if err != nil {
			utils.Logger.Printf("error when deleting history_text data err:%s", err.Error())
		}

		var history_uint []models.History
		err = operations.Delete("history_uint", history_uint)
		if err != nil {
			utils.Logger.Printf("error when deleting history_uint data err:%s", err.Error())
		}

		var trends []models.Trends
		err = operations.Delete("trends", trends)
		if err != nil {
			utils.Logger.Printf("error when deleting trends data err:%s", err.Error())
		}

		var trends_uint []models.TrendsUint
		err = operations.Delete("trends_uint", trends_uint)
		if err != nil {
			utils.Logger.Printf("error when deleting trends_uint data err:%s", err.Error())
		}

		time.Sleep(24 * time.Hour)
	}
}
