package main

import (
	"time"

	"github.com/limanmys/zabbix-cleaner-go/pkg/models"
	"github.com/limanmys/zabbix-cleaner-go/pkg/service"
	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
)

func main() {
	utils.InitLogger()
	for {
		var auditog []models.Audit
		data, err := service.Delete("auditlog", auditog)
		if err != nil {
			utils.Logger.Fatalf("error when deleting auditlog data err:%s", err.Error())
		}
		utils.WriteFile("auditlog", data)

		var events []models.Events
		data, err = service.Delete("events", events)
		if err != nil {
			utils.Logger.Fatalf("error when deleting events data err:%s", err.Error())
		}
		utils.WriteFile("events", data)

		var acknowledges []models.Acknowledges
		data, err = service.Delete("acknowledges", acknowledges)
		if err != nil {
			utils.Logger.Fatalf("error when deleting acknowledges data err:%s", err.Error())
		}
		utils.WriteFile("acknowledges", data)

		var history []models.History
		data, err = service.Delete("history", history)
		if err != nil {
			utils.Logger.Fatalf("error when deleting history data err:%s", err.Error())
		}
		utils.WriteFile("history", data)

		var history_str []models.History
		data, err = service.Delete("history_str", history_str)
		if err != nil {
			utils.Logger.Fatalf("error when deleting history_str data err:%s", err.Error())
		}
		utils.WriteFile("history_str", data)

		var history_text []models.History
		data, err = service.Delete("history_text", history_text)
		if err != nil {
			utils.Logger.Fatalf("error when deleting history_text data err:%s", err.Error())
		}
		utils.WriteFile("history_str", data)

		var history_uint []models.History
		data, err = service.Delete("history_uint", history_uint)
		if err != nil {
			utils.Logger.Fatalf("error when deleting history_uint data err:%s", err.Error())
		}
		utils.WriteFile("history_uint", data)

		var trends []models.Trends
		data, err = service.Delete("trends", trends)
		if err != nil {
			utils.Logger.Fatalf("error when deleting trends data err:%s", err.Error())
		}
		utils.WriteFile("trends", data)

		var trends_uint []models.TrendsUint
		data, err = service.Delete("trends_uint", trends_uint)
		if err != nil {
			utils.Logger.Fatalf("error when deleting trends_uint data err:%s", err.Error())
		}
		utils.WriteFile("trends_uint", data)

		time.Sleep(24 * time.Hour)
	}

}
