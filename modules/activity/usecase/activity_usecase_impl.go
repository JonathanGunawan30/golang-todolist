package usecase

import (
	"todolist-v1/modules/activity/entities"
	"todolist-v1/modules/activity/repository"
)

type activityUsecaseImpl struct {
	activityRepository repository.ActivityRepository
}

func NewActivityUsecase(activityRepository repository.ActivityRepository) ActivityUsecase {
	return &activityUsecaseImpl{activityRepository}
}

func (usecase *activityUsecaseImpl) GetAll() ([]entities.Activity, error) {
	return usecase.activityRepository.FindAll()
}

func (usecase *activityUsecaseImpl) Create(activity entities.Activity) (entities.Activity, error) {
	activity.Status = "NEW"
	return usecase.activityRepository.Save(activity)
}

func (usecase *activityUsecaseImpl) Update(id int, activity entities.Activity) (entities.Activity, error) {
	return usecase.activityRepository.Update(id, activity)
}

func (usecase *activityUsecaseImpl) Delete(id int) error {
	return usecase.activityRepository.Delete(id)
}
