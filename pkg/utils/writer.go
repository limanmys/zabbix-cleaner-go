package utils

import (
	"os"
	"time"
)

func WriteFile(table, data string) error {
	f, err := os.OpenFile("/var/log/zabbix-db-cleaner/deleted_rows/"+table+"/deleted_"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
