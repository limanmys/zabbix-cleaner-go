package models

import "time"

type Alerts struct {
	Alertid  int           `gorm:"type:text" json:"alert_id"`
	Actionid int           `gorm:"type:text" json:"action_id"`
	Eventid  int           `gorm:"type:text" json:"event_id"`
	Userid   int           `gorm:"type:text" json:"user_id"`
	Clock    time.Duration `gorm:"type:text" json:"time"`
}
