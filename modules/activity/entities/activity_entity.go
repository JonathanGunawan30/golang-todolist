package entities

import "time"

type Activity struct {
	Id           int       `json:"id"            gorm:"column:id;primaryKey;autoIncrement"`
	Title        string    `json:"title"         gorm:"column:title;size:250;not null"`
	Category     string    `json:"category"      gorm:"column:category;not null"`
	Description  string    `json:"description"   gorm:"column:description;type:text;not null"`
	ActivityDate time.Time `json:"activity_date" gorm:"column:activity_date;not null"`
	Status       string    `json:"status"        gorm:"column:status;not null;default:NEW"`
}

func (Activity) TableName() string { return "activities" }
