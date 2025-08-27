package repository

import (
	"errors"
	"todolist-v1/modules/activity/entities"
)

var ErrActivityNotFound = errors.New("activity not found")

type ActivityRepository interface {
	FindAll() ([]entities.Activity, error)
	Save(activity entities.Activity) (entities.Activity, error)
	Update(id int, activity entities.Activity) (entities.Activity, error)
	Delete(id int) error
}
