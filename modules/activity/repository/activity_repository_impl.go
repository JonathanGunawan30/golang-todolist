package repository

import (
	"errors"
	"todolist-v1/modules/activity/entities"

	"gorm.io/gorm"
)

type activityRepositoryImpl struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepositoryImpl{DB: db}
}

func (repository *activityRepositoryImpl) FindAll() ([]entities.Activity, error) {
	var activities []entities.Activity
	if err := repository.DB.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (repository *activityRepositoryImpl) Save(activity entities.Activity) (entities.Activity, error) {
	if err := repository.DB.Create(&activity).Error; err != nil {
		return entities.Activity{}, err
	}
	return activity, nil
}

func (repository *activityRepositoryImpl) Update(id int, activity entities.Activity) (entities.Activity, error) {
	result := repository.DB.Model(&entities.Activity{}).Where("id = ?", id).Updates(map[string]any{
		"title":         activity.Title,
		"category":      activity.Category,
		"description":   activity.Description,
		"activity_date": activity.ActivityDate,
		"status":        activity.Status,
	})
	if result.Error != nil {
		return entities.Activity{}, result.Error
	}
	if result.RowsAffected == 0 {
		return entities.Activity{}, ErrActivityNotFound
	}

	var updated entities.Activity
	if err := repository.DB.First(&updated, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Activity{}, ErrActivityNotFound
		}
		return entities.Activity{}, err
	}

	return activity, nil
}

func (repository *activityRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&entities.Activity{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrActivityNotFound
	}
	return nil
}
