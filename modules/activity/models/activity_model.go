package models

import "time"

type ActivityCreateRequest struct {
	Title        string    `json:"title" validate:"required,max=250,min=3"`
	Category     string    `json:"category" validate:"required,oneof=TASK EVENT"`
	Description  string    `json:"description" validate:"required"`
	ActivityDate time.Time `json:"activity_date" validate:"required"`
}

type ActivityUpdateRequest struct {
	Title        string    `json:"title" validate:"required,max=250"`
	Category     string    `json:"category" validate:"required,oneof=TASK EVENT"`
	Description  string    `json:"description" validate:"required"`
	ActivityDate time.Time `json:"activity_date" validate:"required"`
	Status       string    `json:"status" validate:"required,oneof=NEW 'ON PROGRESS' EXPIRED"`
}

type ActivityResponse struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	ActivityDate time.Time `json:"activity_date"`
	Status       string    `json:"status"`
}
