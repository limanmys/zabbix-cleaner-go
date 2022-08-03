package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	Logger *log.Logger
)

func InitLogger() {
	err := os.MkdirAll(filepath.FromSlash("/var/log/zabbix-db-cleaner/"), 0666)
	if err != nil {
		log.Fatal("cannot create log directories")
	}

	err = os.MkdirAll(filepath.FromSlash("/var/log/zabbix-db-cleaner/deleted_rows/"), 0666)
	if err != nil {
		log.Fatal("cannot create log directories")
	}

	directories := []string{"alerts", "events", "history", "history_uint", "history_str", "history_text", "auditlog", "acknowledges", "trends", "trends_uint"}
	for _, element := range directories {
		err = os.MkdirAll(filepath.FromSlash("/var/log/zabbix-db-cleaner/deleted_rows/"+element+"/"), 0666)
		if err != nil {
			log.Fatal("cannot create log directories")
		}
	}

	info, err := os.OpenFile(filepath.FromSlash("/var/log/zabbix-db-cleaner/cleaner.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("cannot create log file")
	}

	// Standart logger
	Logger = log.New(info, log.Prefix(), log.Ldate|log.Ltime|log.Lshortfile)
	InfoWriter := io.MultiWriter(os.Stdout, info)
	Logger.SetOutput(InfoWriter)
}
