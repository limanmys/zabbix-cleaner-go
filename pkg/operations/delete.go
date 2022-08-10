package operations

import (
	"fmt"

	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
	"github.com/limanmys/zabbix-cleaner-go/platform/database"
)

func Delete(tableName string) error {
	deleteTime, err := utils.GetTime()
	if err != nil {
		return err
	}

	rawQuery := fmt.Sprintf("DELETE FROM %s WHERE clock < %d", tableName, deleteTime)
	if err != nil {
		return err
	}

	result, err := database.Connection().Exec(rawQuery)
	if err != nil {
		return err
	}
	utils.Logger.Printf("SQL Result : %v\n", result)
	utils.Logger.Printf("%s table successfully cleaned from %d.\n", tableName, deleteTime)
	return nil
}
