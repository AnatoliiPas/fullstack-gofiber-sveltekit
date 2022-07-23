package models

import "github.com/lib/pq"

type Recipe struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Image        string         `json:"image"`
	Descriptions string         `json:"descriptions"`
	Products     pq.StringArray `json:"products" gorm:"type:text[]"`
	Manuals      pq.StringArray `json:"manual" gorm:"type:text[]"`
}
