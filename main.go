package main

import (
	"time"

	"github.com/limanmys/zabbix-cleaner-go/pkg/operations"
	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
)

func main() {
	utils.InitLogger()

	tables := []string{"events", "history", "history_uint", "history_str", "history_text", "auditlog", "acknowledges", "trends", "trends_uint"}
	for {
		for _, table := range tables {
			err := operations.Delete(table)
			if err != nil {
				utils.Logger.Printf("error when deleting %s data err:%s", table, err.Error())
			}

		}
		time.Sleep(24 * time.Hour)
	}
}
