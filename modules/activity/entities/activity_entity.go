package entities

import "time"

type Activity struct {
	Id           int
	Title        string
	Category     string
	Description  string
	ActivityDate time.Time
	Status       string
}
