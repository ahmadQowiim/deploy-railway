package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	//return nil // TODO: replace this
	/*err := t.db.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
	if err != nil {
		return err
	}
	return nil*/
	err := t.db.Model(model.Task{}).Where("id=?", id).UpdateColumns(map[string]interface{}{
		"title":       task.Title,
		"deadline":    task.Deadline,
		"priority":    task.Priority,
		"category_id": task.CategoryID,
		"status":      task.Status,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Delete(id int) error {
	//return nil // TODO: replace this
	var task model.Task
	err := t.db.Delete(&task, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	//return nil, nil // TODO: replace this
	result := make([]model.Task, 0)
	rows, err := t.db.Table("tasks").Select("*").Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := t.db.ScanRows(rows, &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	//return nil, nil // TODO: replace this
	var taskCategories = make([]model.TaskCategory, 0)
	err := t.db.Model(&taskCategories).
		Table("tasks").
		Select("tasks.id as id, tasks.title as title, categories.name as category").
		Joins("left join categories on tasks.category_id=categories.id").
		Where("tasks.id=?", id).
		Scan(&taskCategories).Error
	if err != nil {
		return taskCategories, err
	}

	return taskCategories, nil
}
