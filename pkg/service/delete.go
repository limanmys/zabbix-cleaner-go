package service

import (
	"encoding/json"
	"fmt"

	"github.com/limanmys/zabbix-cleaner-go/pkg/database"
	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
	"gorm.io/gorm/clause"
)

func Delete(tableName string, model any) (string, error) {
	deleteTime, err := utils.GetTime()
	if err != nil {
		return "", err
	}

	database.Connection().Clauses(clause.Returning{}).Table(tableName).Where(fmt.Sprintf("clock < %d", deleteTime)).Delete(&model)

	b, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	utils.Logger.Printf("%s table successfully cleaned from %d.", tableName, deleteTime)
	return string(b), nil
}
