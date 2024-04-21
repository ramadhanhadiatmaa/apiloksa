package models

import "gorm.io/gorm"

type Loketsatu struct {
	gorm.Model
	Seq int `json:"ant"`
	Stat string `json:"status"`
}