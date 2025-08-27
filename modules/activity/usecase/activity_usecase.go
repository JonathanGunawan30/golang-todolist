package usecase

import "todolist-v1/modules/activity/entities"

type ActivityUsecase interface {
	GetAll() ([]entities.Activity, error)
	Create(activity entities.Activity) (entities.Activity, error)
	Update(id int, activity entities.Activity) (entities.Activity, error)
	Delete(id int) error
}
