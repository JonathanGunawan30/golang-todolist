package repository

import (
	"database/sql"
	"errors"
	"todolist-v1/modules/activity/entities"
)

type activityRepositoryImpl struct {
	DB *sql.DB
}

func NewActivityRepository(db *sql.DB) ActivityRepository {
	return &activityRepositoryImpl{DB: db}
}

func (repository *activityRepositoryImpl) FindAll() ([]entities.Activity, error) {
	query := "SELECT id, title, category, description, activity_date, status FROM activities"
	rows, err := repository.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []entities.Activity
	for rows.Next() {
		var act entities.Activity
		if err := rows.Scan(&act.Id, &act.Title, &act.Category, &act.Description, &act.ActivityDate, &act.Status); err != nil {
			return nil, err
		}
		activities = append(activities, act)
	}
	return activities, nil
}

func (repository *activityRepositoryImpl) Save(activity entities.Activity) (entities.Activity, error) {
	query := `INSERT INTO activities(title,category,description,activity_date,status) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := repository.DB.QueryRow(query, activity.Title, activity.Category, activity.Description, activity.ActivityDate, activity.Status).Scan(&activity.Id)
	return activity, err
}

func (repository *activityRepositoryImpl) Update(id int, activity entities.Activity) (entities.Activity, error) {
	query := `UPDATE activities 
	          SET title = $1, category = $2, description = $3, activity_date = $4, status = $5
	          WHERE id = $6 
	          RETURNING id, title, category, description, activity_date, status`

	var updatedActivity entities.Activity

	err := repository.DB.QueryRow(query,
		activity.Title,
		activity.Category,
		activity.Description,
		activity.ActivityDate,
		activity.Status,
		id,
	).Scan(
		&updatedActivity.Id,
		&updatedActivity.Title,
		&updatedActivity.Category,
		&updatedActivity.Description,
		&updatedActivity.ActivityDate,
		&updatedActivity.Status,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Activity{}, ErrActivityNotFound
		}
		return entities.Activity{}, err
	}

	return updatedActivity, nil
}

func (repository *activityRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM activities WHERE id = $1"

	result, err := repository.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrActivityNotFound
	}

	return nil
}
