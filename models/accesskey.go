package models

import "gorm.io/gorm"

type Accesskey struct {
	gorm.Model
	Key    string `json:"key"`
	Status string `json:"status"`
}
