package operations

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/limanmys/zabbix-cleaner-go/pkg/utils"
	"github.com/limanmys/zabbix-cleaner-go/platform/database"
	"gorm.io/gorm/clause"
)

func Delete(tableName string, model any) error {
	deleteTime, err := utils.GetTime()
	if err != nil {
		return err
	}

	database.Connection().Clauses(clause.Returning{}).Table(tableName).Where(fmt.Sprintf("clock < %d", deleteTime)).Delete(&model)

	b, err := json.Marshal(model)
	if err != nil {
		return err
	}

	if os.Getenv("SAVE_DELETED_ROWS") != "false" {
		utils.WriteFile(tableName, string(b))

	}
	utils.Logger.Printf("%s table successfully cleaned from %d.", tableName, deleteTime)
	return nil
}
